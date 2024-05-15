package iok

import (
	"bytes"
	"context"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
)

// InputFromHTTPResponse takes a *http.Response returns an Input suitable for calling GetMatches with.
// This is significantly weaker than e.g. InputFromURLScan because it cannot:
// * Identify all requests made by the page (e.g. ones triggered by JavaScript)
// * Fetch the contents of JavaScript/CSS files referenced by the page
func InputFromHTTPResponse(resp *http.Response) (Input, error) {
	input := Input{}
	input.SetURL(resp.Request.URL)

	for header, values := range resp.Header {
		for _, value := range values {
			input.AddHeader(http.CanonicalHeaderKey(header) + ": " + value)
		}
	}

	for _, cookie := range resp.Cookies() {
		input.AddCookie(cookie.Name + "=" + cookie.Value)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return input, err
	}
	input.AddHTML(string(body))

	// parse any JS/CSS from the html
	node, err := html.Parse(bytes.NewReader(body))
	if err == nil {
		extractHTML(node, &input,
			extractEmbeddedAssets,
			extractTitle,
			extractRequests(resp.Request.URL), // TODO: what is the correct url to use here in the case of redirects?
		)
	}

	return input, nil
}

// InputFromURL fetches the contents of a URL using the supplied HTTP client
// and constructs an Input suitable for calling GetMatches with.
func InputFromURL(ctx context.Context, u *url.URL, client *http.Client) (Input, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	resp, err := client.Do(req)
	if err != nil {
		return Input{}, err
	}
	return InputFromHTTPResponse(resp)
}
