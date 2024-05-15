package iok

import (
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

func extractHTML(node *html.Node, input *Input, funcs ...func(node *html.Node, input *Input)) {
	for _, f := range funcs {
		f(node, input)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extractHTML(c, input, funcs...)
	}
}

func extractEmbeddedAssets(node *html.Node, input *Input) {
	if node.Type == html.ElementNode && node.Data == "style" && node.FirstChild != nil {
		input.AddCSS(node.FirstChild.Data)
	}
	if node.Type == html.ElementNode && node.Data == "script" && node.FirstChild != nil {
		input.AddJS(node.FirstChild.Data)
	}
}

func extractTitle(node *html.Node, input *Input) {
	if node.Type == html.ElementNode &&
		node.Data == "title" && // this is a <title> tag
		node.Namespace == "" && // we're not within an <svg> element
		node.FirstChild != nil {
		input.AddTitle(strings.TrimSpace(node.FirstChild.Data))
	}
}

// extractRequests analyses the html to determine what requests *would*
// be made if rendered in a browser.
// This is a subset of the actual requests made because this won't include
// any requests triggered dynamically e.g. by JavaScript or by CSS styles
func extractRequests(base *url.URL) func(node *html.Node, input *Input) {
	return func(node *html.Node, input *Input) {
		// TODO: handle <base> tag changing the base url
		if node.Type != html.ElementNode {
			return
		}
		for _, attr := range node.Attr {
			var u string
			switch {
			case node.Data == "link" && attr.Key == "href": // todo: filter to just rel=stylesheet
				fallthrough
			case node.Data == "img" && attr.Key == "src":
				fallthrough
			case node.Data == "script" && attr.Key == "src":
				u = attr.Val
			}
			if u == "" {
				continue
			}
			href, err := url.Parse(u)
			if err != nil {
				continue
			}
			input.AddRequest(base.ResolveReference(href))
		}
	}
}
