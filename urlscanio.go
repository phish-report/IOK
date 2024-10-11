package iok

import (
	"bytes"
	"context"
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net/http"
	"net/url"
	urlscan "phish.report/urlscanio-go"
	"sync"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// InputFromURLScan takes a urlscan.io result ID and returns an Input suitable for calling GetMatches with.
// The provided http.Client should inject your API key if you have one.
// Deprecated: use InputsFromURLScan
func InputFromURLScan(ctx context.Context, urlscanUUID string, client httpClient) (Input, error) {
	inputs, err := InputsFromURLScan(ctx, urlscanUUID, client)
	if err != nil || len(inputs) == 0 {
		return Input{}, err
	}
	for _, input := range inputs {
		if input.primary {
			return input.Input, nil
		}
	}

	return inputs[len(inputs)-1].Input, nil
}

type URLScanInput struct {
	LoaderID    string
	DocumentURL string
	Input
	primary bool
}

func InputsFromURLScan(ctx context.Context, urlscanUUID string, client httpClient) ([]*URLScanInput, error) {
	urlscanClient := urlscan.NewClient(urlscan.HTTPClient(client))
	result, err := urlscanClient.RetrieveResult(ctx, urlscanUUID)
	if err != nil {
		return nil, err
	}

	inputs := []*URLScanInput{}
	inputsByLoader := map[string]*URLScanInput{}
	var primaryRequest *URLScanInput
	for _, request := range result.Data.Requests {
		loader := request.Request.LoaderId
		if _, ok := inputsByLoader[loader]; ok {
			continue
		}

		docURL, err := url.Parse(request.Request.DocumentURL)
		if err != nil {
			log.Println("iok: failed to parse document url", urlscanUUID, request.Request.DocumentURL)
		}
		hostname := ""
		if docURL != nil {
			hostname = docURL.Hostname()
		}
		input := &URLScanInput{
			LoaderID:    loader,
			DocumentURL: request.Request.DocumentURL,
			Input: Input{
				Hostname: hostname,
			},
		}
		inputs = append(inputs, input)
		inputsByLoader[loader] = input

		// Keep track of the "primary" request (i.e. the loader ID of the final page shown in the browser)
		// because this is the one that corresponds to urlscan's DOM response
		if request.Request.PrimaryRequest {
			primaryRequest = input
			primaryRequest.primary = true
		}
	}

	// Some sites have many resources (100+) so fetching each one sequentially takes too long.
	// This fetches up to 5 resources in parallel
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(5)
	mu := sync.Mutex{}

	g.Go(func() error {
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
		primaryRequest.DOM = string(resultHTML)

		// parse any JS/CSS from the dom
		node, err := html.Parse(bytes.NewReader(resultHTML))
		if err == nil {
			extractHTML(node, &primaryRequest.Input, extractEmbeddedAssets, extractTitle)
		}
		return nil
	})

	// we can't tell which request set the cookie, so we just set it on the primary input
	for _, cookie := range result.Data.Cookies {
		primaryRequest.Cookies = append(primaryRequest.Cookies, cookie.Name+"="+cookie.Value)
	}

	for _, request := range result.Data.Requests {
		request := request
		input := &inputsByLoader[request.Request.LoaderId].Input
		g.Go(func() error {
			mu.Lock()
			input.Requests = append(input.Requests, request.Request.Request.Url)
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
				input.CSS = append(input.CSS, string(resource))
			case "Script":
				input.JS = append(input.JS, string(resource))
			case "Document":
				if input.HTML != "" {
					// already have initial HTML
					break
				}

				// this is the initial page load (for this loaderID)
				input.HTML = string(resource)
				// extract the response headers
				for headerKey, headerValue := range request.Response.Response.Headers {
					input.Headers = append(input.Headers, http.CanonicalHeaderKey(headerKey)+": "+headerValue)
				}

				// parse any JS/CSS from the html
				// This does result in duplicate values (for sites that don't have any dynamically inserted JS/CSS),
				// but that doesn't affect correctness
				node, err := html.Parse(bytes.NewReader(resource))
				if err == nil {
					extractHTML(node, input, extractEmbeddedAssets, extractTitle)
				}
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return inputs, err
	}

	return inputs, nil
}
