package iok

import (
	"net"
	"net/url"
	"slices"
)

//go:generate go run ./internal/genlogsource
//go:generate go fmt input.go
//go:generate go run ./internal/gendocs
type Input struct {
	URL             *url.URL
	urlHostname     string
	urlPath         string
	urlQuery        []string
	IP              []net.IP
	CNAME           []string
	asn             int
	Registrar       int
	Status          int
	Header          []string
	Title           []string
	HTML            []string
	JS              []string
	CSS             []string
	Cookie          []string
	Referrer        string
	Redirect        string
	RedirectKind    string
	FaviconHash     []string
	Request         []*url.URL
	requestHostname []string
	RequestIP       []net.IP
	requestASN      []int
	RequestHeader   []string
	ResponseHash    []string
	ResponseHeader  []string
	Supported       []string
	unsupported     []string
}

func (i *Input) SetURL(v *url.URL) {
	i.URL = v
	i.Supported = append(i.Supported, "url")
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
func (i *Input) AddRequestIP(v ...net.IP) {
	i.RequestIP = append(i.RequestIP, v...)
	i.Supported = append(i.Supported, "request.ip")
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
