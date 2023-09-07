package iok

import (
	"bytes"
	"context"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
	"phish.report/urlscanio-go"
	"sort"
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

	input := Input{}
	u, err := url.Parse(result.Page.Url)
	if err != nil {
		return Input{}, fmt.Errorf("failed to parse result URL: %w", err)
	}
	input.Hostname = u.Hostname()

	domReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://urlscan.io/dom/"+result.Task.Uuid, nil)
	domResp, err := client.Do(domReq)
	if err != nil {
		return Input{}, fmt.Errorf("failed to get result html: %w", err)
	}
	defer domResp.Body.Close()

	resultHTML, _ := io.ReadAll(domResp.Body)
	input.DOM = string(resultHTML)

	// parse any JS/CSS from the dom
	node, err := html.Parse(bytes.NewReader(resultHTML))
	if err == nil {
		extractEmbedded(node, &input)
		extractTitle(node, &input)
	}

	for _, cookie := range result.Data.Cookies {
		input.Cookies = append(input.Cookies, cookie.Name+"="+cookie.Value)
	}
	foundHTML := false
	for _, request := range result.Data.Requests {
		input.Requests = append(input.Requests, request.Request.Request.Url)

		// TODO: how does this check behave in the case of redirects?
		if request.Request.PrimaryRequest {
			// this is the "primary" page load, so we need to extract the response headers
			for headerKey, headerValue := range request.Response.Response.Headers {
				input.Headers = append(input.Headers, http.CanonicalHeaderKey(headerKey)+": "+headerValue)
			}
			sort.Slice(input.Headers, func(i, j int) bool {
				return input.Headers[i] < input.Headers[j]
			})
		}

		if request.Response.Hash == "" {
			// this isn't a response we can fetch
			continue
		}

		switch request.Request.Type {
		case "Stylesheet", "Script", "Document":
			resourceReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://urlscan.io/responses/"+request.Response.Hash, nil)
			resp, err := client.Do(resourceReq)
			if err != nil {
				return Input{}, fmt.Errorf("failed to fetch resource %s %s: %w", request.Request.RequestId, request.Response.Hash, err)
			}
			resource, _ := io.ReadAll(resp.Body) // always read the body to completion to ensure proper connection re-use + caching
			resp.Body.Close()
			if resp.StatusCode/100 != 2 {
				// not all resources are saved by urlscan.io e.g. stylesheets are frequently missing
				continue
			}
			switch request.Request.Type {
			case "Stylesheet":
				input.CSS = append(input.CSS, string(resource))
			case "Script":
				input.JS = append(input.JS, string(resource))
			case "Document":
				if request.Request.PrimaryRequest {
					foundHTML = true
					if input.HTML != "" {
						fmt.Println("oops already have response html")
					}
					// this is the initial page load
					input.HTML = string(resource)

					// parse any JS/CSS from the html
					// This does result in duplicate values (for sites that don't have any dynamically inserted JS/CSS),
					// but that doesn't affect correctness
					node, err := html.Parse(bytes.NewReader(resource))
					if err == nil {
						extractEmbedded(node, &input)
						extractTitle(node, &input)
					}
				}
			}
		}
	}

	if !foundHTML {
		return input, fmt.Errorf("failed to get response html")
	}

	return input, nil
}

func extractEmbedded(node *html.Node, input *Input) {
	if node.Type == html.ElementNode && node.Data == "style" && node.FirstChild != nil {
		input.CSS = append(input.CSS, node.FirstChild.Data)
	}
	if node.Type == html.ElementNode && node.Data == "script" && node.FirstChild != nil {
		input.JS = append(input.JS, node.FirstChild.Data)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extractEmbedded(c, input)
	}
}

func extractTitle(node *html.Node, input *Input) {
	if node.Type == html.ElementNode && node.Data == "title" && node.FirstChild != nil {
		input.Title = append(input.Title, node.FirstChild.Data)
		return
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extractTitle(c, input)
	}
}
