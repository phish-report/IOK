package iok

import (
	"context"
	"net"
	"net/url"
	"slices"
)

//go:generate go run ./internal/genlogsource
//go:generate go fmt input.go
//go:generate go run ./internal/gendocs
type Input struct {
	ipToASN         func(context.Context, net.IP) int
	URL             *url.URL   // The full URL of the page
	URLHostname     string     // The hostname of the page (derived from url)
	URLPath         string     // Path component of the full URL of the page (derived from url)
	URLQuery        []string   // Query parameters extracted from the full URL of the page (derived from url)
	IP              []net.IP   // IP addresses from which the page was loaded (resolved from url.hostname)
	CNAME           []string   // Any CNAME records for the primary page URL, allowing you to detect things like pages hosted on GitHub
	ASN             []int      // ASN from which the page was loaded (derived from ip)
	Registrar       int        // IANA ID for the registrar of this domain
	Status          int        // The HTTP status code of the primary page load
	Header          []string   // HTTP Headers returned as part of page load. Each is in the canonical form Header-Name: value
	Title           []string   // The title of the page as shown in a browser. If multiple titles are set (e.g. by JavaScript), this contains all of them
	HTML            []string   // The contents of the page HTML. If JavaScript is used to modify the contents of the page, this contains multiple values
	JS              []string   // Contents of JavaScript from the page (includes inline scripts as well as scripts loaded externally)
	CSS             []string   // Contents of CSS from the page (includes inline stylesheets as well as externally loaded stylesheets)
	Cookie          []string   // Cookies from the page. Each is in the form cookieName=value
	Referrer        string     // If this page is part of a redirect chain, the URL it redirected from
	Redirect        string     // If this page redirected the browser to another page, the URL it redirected to
	RedirectKind    string     // If this page redirected the browser to another page, how was it redirected. One of: http{status}, js, meta
	FaviconHash     []string   // SHA256 hash of favicons specified by the page. All favicons found in the page will be included.
	Request         []*url.URL // URLs of requests made by the page (and assets loaded by the page)
	RequestHostname []string   // Hostnames of requests made by the page (and assets loaded by the page) (derived from requests)
	RequestIP       []net.IP   // IPs contacted as part of requests made by the page
	RequestASN      []int      // ASNs contacted as part of requests made by the page (derived from requests.ip)
	RequestHeader   []string   // Headers in requests made during the page load. Each is in the canonical form Header-Name: value
	ResponseHash    []string   // SHA256 Hashes of response bodies
	ResponseHeader  []string   // Headers returned as part of responses. Each is in the canonical form Header-Name: value
	Supported       []string   // Fields supported by this IOK implementation. This allows you to distinguish between fields which are genuinely empty (e.g. the page had no <title> element) from cases where the data was not loaded.
	Unsupported     []string   // Fields *not* supported by this IOK implementation. For example, the favicon may not have been fetched and so the favicon.hash field will be unset. (derived from supported)
}

func (i *Input) SetURL(v *url.URL) {
	i.URL = v
	i.Supported = append(i.Supported, "url")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
	i.setURLHostname(v.Hostname())
	i.setURLPath(v.Path)
}
func (i *Input) setURLHostname(v string) {
	i.URLHostname = v
	i.Supported = append(i.Supported, "url.hostname")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) setURLPath(v string) {
	i.URLPath = v
	i.Supported = append(i.Supported, "url.path")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) addURLQuery(v ...string) {
	i.URLQuery = append(i.URLQuery, v...)
	i.Supported = append(i.Supported, "url.query")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) AddIP(v ...net.IP) {
	i.IP = append(i.IP, v...)
	i.Supported = append(i.Supported, "ip")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) AddCNAME(v ...string) {
	i.CNAME = append(i.CNAME, v...)
	i.Supported = append(i.Supported, "cname")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) addASN(v ...int) {
	i.ASN = append(i.ASN, v...)
	i.Supported = append(i.Supported, "asn")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) SetRegistrar(v int) {
	i.Registrar = v
	i.Supported = append(i.Supported, "registrar")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) SetStatus(v int) {
	i.Status = v
	i.Supported = append(i.Supported, "status")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) AddHeader(v ...string) {
	i.Header = append(i.Header, v...)
	i.Supported = append(i.Supported, "header")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) AddTitle(v ...string) {
	i.Title = append(i.Title, v...)
	i.Supported = append(i.Supported, "title")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) AddHTML(v ...string) {
	i.HTML = append(i.HTML, v...)
	i.Supported = append(i.Supported, "html")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) AddJS(v ...string) {
	i.JS = append(i.JS, v...)
	i.Supported = append(i.Supported, "js")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) AddCSS(v ...string) {
	i.CSS = append(i.CSS, v...)
	i.Supported = append(i.Supported, "css")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) AddCookie(v ...string) {
	i.Cookie = append(i.Cookie, v...)
	i.Supported = append(i.Supported, "cookie")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) SetReferrer(v string) {
	i.Referrer = v
	i.Supported = append(i.Supported, "referrer")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) SetRedirect(v string) {
	i.Redirect = v
	i.Supported = append(i.Supported, "redirect")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) SetRedirectKind(v string) {
	i.RedirectKind = v
	i.Supported = append(i.Supported, "redirect.kind")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) AddFaviconHash(v ...string) {
	i.FaviconHash = append(i.FaviconHash, v...)
	i.Supported = append(i.Supported, "favicon.hash")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) AddRequest(v ...*url.URL) {
	i.Request = append(i.Request, v...)
	i.Supported = append(i.Supported, "request")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) addRequestHostname(v ...string) {
	i.RequestHostname = append(i.RequestHostname, v...)
	i.Supported = append(i.Supported, "request.hostname")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) AddRequestIP(v ...net.IP) {
	i.RequestIP = append(i.RequestIP, v...)
	i.Supported = append(i.Supported, "request.ip")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) addRequestASN(v ...int) {
	i.RequestASN = append(i.RequestASN, v...)
	i.Supported = append(i.Supported, "request.asn")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) AddRequestHeader(v ...string) {
	i.RequestHeader = append(i.RequestHeader, v...)
	i.Supported = append(i.Supported, "request.header")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) AddResponseHash(v ...string) {
	i.ResponseHash = append(i.ResponseHash, v...)
	i.Supported = append(i.Supported, "response.hash")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) AddResponseHeader(v ...string) {
	i.ResponseHeader = append(i.ResponseHeader, v...)
	i.Supported = append(i.Supported, "response.header")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) AddSupported(v ...string) {
	i.Supported = append(i.Supported, v...)
	i.Supported = append(i.Supported, "supported")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
func (i *Input) addUnsupported(v ...string) {
	i.Unsupported = append(i.Unsupported, v...)
	i.Supported = append(i.Supported, "unsupported")
	slices.Sort(i.Supported)
	i.Supported = slices.Compact(i.Supported)
}
