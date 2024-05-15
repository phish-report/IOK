package iok

import (
	"bytes"
	"context"
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/sync/errgroup"
	"io"
	"net"
	"net/http"
	"net/url"
	"phish.report/urlscanio-go"
	"sync"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// InputFromURLScan takes a urlscan.io result ID and returns an Input suitable for calling GetMatches with.
// The provided http.Client should inject your API key if you have one.
func InputFromURLScan(ctx context.Context, urlscanUUID string, client httpClient) (Input, error) {
	urlscanClient := urlscanio.NewClient(urlscanio.HTTPClient(client))
	result, err := urlscanClient.RetrieveResult(ctx, urlscanUUID)
	if err != nil {
		return Input{}, err
	}

	input := Input{
		Supported: []string{ // even if these fields aren't set below, they're still supported
			"html",
			"js",
			"css",
		},
	}
	u, err := url.Parse(result.Page.Url)
	if err != nil {
		return Input{}, fmt.Errorf("failed to parse result URL: %w", err)
	}
	input.SetURL(u)

	// Some sites have many resources (100+) so fetching each one sequentially takes too long.
	// This fetches up to 5 resources in parallel
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(5)
	mu := &sync.Mutex{}

	g.Go(func() error {
		return getDom(ctx, client, result, mu, &input)
	})

	for _, cookie := range result.Data.Cookies {
		input.AddCookie(cookie.Name + "=" + cookie.Value)
	}

	for _, request := range result.Data.Requests {
		request := request
		g.Go(func() error {
			mu.Lock()
			reqURL, _ := url.Parse(request.Request.Request.Url)
			input.AddRequest(reqURL)

			// TODO: how does this check behave in the case of redirects?
			if request.Request.PrimaryRequest {
				// this is the "primary" page load, so we need to extract the response headers
				input.AddIP(net.ParseIP(request.Response.Response.RemoteIPAddress))
				for headerKey, headerValue := range request.Response.Response.Headers {
					input.AddHeader(http.CanonicalHeaderKey(headerKey) + ": " + headerValue)
				}
			} else {
				input.AddRequestIP(net.ParseIP(request.Response.Response.RemoteIPAddress))
				for headerKey, headerValue := range request.Request.Request.Headers {
					input.AddRequestHeader(http.CanonicalHeaderKey(headerKey) + ": " + headerValue)
				}
			}

			input.AddResponseHash(request.Response.Hash)
			for headerKey, headerValue := range request.Response.Response.Headers {
				input.AddResponseHeader(http.CanonicalHeaderKey(headerKey) + ": " + headerValue)
			}
			mu.Unlock()

			if request.Response.Hash == "" {
				// this isn't a response we can fetch
				return nil
			}

			switch request.Request.Type {
			default:
				return nil
			case "Stylesheet", "Script", "Document":
			}

			// Fetch the response in parallel with other threads, only lock the mutex once we're modifying the Input{}
			resourceReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://urlscan.io/responses/"+request.Response.Hash, nil)
			resp, err := client.Do(resourceReq)
			if err != nil || resp.StatusCode != 200 && resp.StatusCode != 404 {
				if err == nil {
					err = fmt.Errorf(resp.Status)
				}
				return fmt.Errorf("failed to fetch resource %s %s: %w", request.Request.RequestId, request.Response.Hash, err)
			}
			resource, _ := io.ReadAll(resp.Body) // always read the body to completion to ensure proper connection re-use + caching
			resp.Body.Close()
			if resp.StatusCode/100 != 2 {
				// not all resources are saved by urlscan.io e.g. stylesheets are frequently missing
				return nil
			}

			mu.Lock()
			defer mu.Unlock()
			switch request.Request.Type {
			case "Stylesheet":
				input.AddCSS(string(resource))
			case "Script":
				input.AddJS(string(resource))
			case "Document":
				if request.Request.PrimaryRequest {
					// this is the initial page load
					input.AddHTML(string(resource))

					// parse any JS/CSS from the html
					// This does result in duplicate values (for sites that don't have any dynamically inserted JS/CSS),
					// but that doesn't affect correctness
					node, err := html.Parse(bytes.NewReader(resource))
					if err == nil {
						extractHTML(node, &input, extractEmbeddedAssets, extractTitle)
					}
				}
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return input, err
	}

	return input, nil
}

func getDom(ctx context.Context, client httpClient, result urlscanio.ScanResult, mu *sync.Mutex, input *Input) error {
	domReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://urlscan.io/dom/"+result.Task.Uuid, nil)
	domResp, err := client.Do(domReq)
	if err != nil || domResp.StatusCode != 200 {
		if err == nil {
			err = fmt.Errorf(domResp.Status)
		}
		return fmt.Errorf("failed to get result dom: %w", err)
	}
	defer domResp.Body.Close()

	mu.Lock()
	defer mu.Unlock()
	resultHTML, _ := io.ReadAll(domResp.Body)
	input.HTML = append(input.HTML, string(resultHTML))
	input.Supported = append(input.Supported, "html")

	// parse any JS/CSS from the dom
	node, err := html.Parse(bytes.NewReader(resultHTML))
	if err == nil {
		extractEmbeddedAssets(node, input)
		extractTitle(node, input)
	}
	return nil
}
