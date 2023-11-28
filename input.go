package iok

//go:generate go run ./internal/genlogsource

type Input struct {
  Url []string // URLs of all requests made during page load. Contains values from both page.url and request.url
  Ip []string // All IPs contacted during page load. Contains values from both page.ip and request.ip
  Asn []string // All ASNs contacted during page load. Contains values from both page.asn and request.asn
  PageUrl string // The full URL of the page
  PageStatus string // The HTTP status code of the primary page load
  Title []string // The title of the page as shown in a browser. If multiple titles are set (e.g. by JavaScript), this contains all of them.
  Html []string // The contents of the page HTML. If JavaScript is used to modify the contents of the page, this contains multiple values.
  Js []string // Contents of JavaScript from the page (includes inline scripts as well as scripts loaded externally).
  Css []string // Contents of CSS from the page (includes inline stylesheets as well as externally loaded stylesheets)
  Cookies []string // Cookies from the page. Each is in the form cookieName=value
  PageIp string // IP address from which the page was loaded.
  PageAsn string // ASN from which the page was loaded.
  Referrer string // If this page is part of a redirect chain, the URL it redirected from.
  RedirectUrl string // If this page redirected the browser to another page, the URL it redirected to.
  RedirectKind string // If this page redirected the browser to another page, how was it redirected. One of: http, js, meta
  RequestUrl []string // URLs of requests made by the page (and assets loaded by the page).
  RequestIp []string // IPs contacted as part of requests made by the page.
  RequestAsn []string // ASNs contacted as part of requests made by the page.
  RequestHeader []string // Headers in requests made during the page load. Each is in the canonical form Header-Name: value
  ResponseHeader []string // Headers returned by the server. Each is in the canonical form Header-Name: value
  ResponseHash []string // SHA256 Hashes of response bodies.
}
