package iok

import "net/url"

//go:generate go run ./internal/genlogsource
//go:generate go fmt input.go

type Input struct {
	Url              *url.URL   // The full URL of the page
	urlHostname      string     // The hostname of the page
	urlPath          string     // Path component of the full URL of the page
	urlQuery         []string   // Query parameters extracted from the full URL of the page
	Ip               []string   // IP addresses from which the page was loaded (resolved from url.hostname)
	asn              int        // ASN from which the page was loaded
	Registrar        int        // IANA ID for the registrar of this domain
	Status           int        // The HTTP status code of the primary page load
	Header           []string   // HTTP Headers returned as part of page load. Each is in the canonical form Header-Name: value
	Title            []string   // The title of the page as shown in a browser. If multiple titles are set (e.g. by JavaScript), this contains all of them
	Html             []string   // The contents of the page HTML. If JavaScript is used to modify the contents of the page, this contains multiple values
	Js               []string   // Contents of JavaScript from the page (includes inline scripts as well as scripts loaded externally)
	Css              []string   // Contents of CSS from the page (includes inline stylesheets as well as externally loaded stylesheets)
	Cookie           []string   // Cookies from the page. Each is in the form cookieName=value
	Referrer         string     // If this page is part of a redirect chain, the URL it redirected from
	Redirect         string     // If this page redirected the browser to another page, the URL it redirected to
	RedirectKind     string     // If this page redirected the browser to another page, how was it redirected. One of: http{status}, js, meta
	FaviconHash      []string   // SHA256 hash of favicons specified by the page. All favicons found in the page will be included.
	Requests         []*url.URL // URLs of requests made by the page (and assets loaded by the page)
	requestsHostname []string   // Hostnames of requests made by the page (and assets loaded by the page)
	RequestsIp       []string   // IPs contacted as part of requests made by the page
	requestsAsn      []int      // ASNs contacted as part of requests made by the page
	RequestHeader    []string   // Headers in requests made during the page load. Each is in the canonical form Header-Name: value
	ResponseHash     []string   // SHA256 Hashes of response bodies
	ResponseHeader   []string   // Headers returned as part of responses. Each is in the canonical form Header-Name: value
	Supported        []string   // Fields supported by this IOK implementation. This allows you to distinguish between fields which are genuinely empty (e.g. the page had no <title> element) from cases where the data was not loaded.
	unsupported      []string   // Fields *not* supported by this IOK implementation. For example, the favicon may not have been fetched and so the favicon.hash field will be unset.
}
