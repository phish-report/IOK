package schema

import (
	"github.com/dave/jennifer/jen"
	"slices"
)

type Field struct {
	SigmaName   string
	Name        string
	Alias       string
	GoType      string
	Type        FieldType
	Description string
	Example     string
	Modifiers   []Modifier
	Derived     string // If this field is automatically derived from another by the IOK library
	PostSetter  func(g *jen.Group)
	Deprecated  bool
}

type Modifier struct {
	Name        string
	Description string
}

type FieldType string

var (
	String     FieldType = "string"
	StringList FieldType = "string list"
	Number     FieldType = "number"
	NumberList FieldType = "number list"
)

var Fields = []Field{
	// Combo fields
	//{
	//	SigmaName:   "url",
	//	Name:        "URL",
	//	Type:        String,
	//	Description: "URL of the requested page",
	//	Example:     "https://phish.domain/?foo=bar",
	//	Modifiers:   append(StandardModifiers, ...),
	//},
	//{
	//	SigmaName:   "ips",
	//	Name:        "IPs",
	//	Type:        StringList,
	//	Alias:       []string{"asn", "request.asn"},
	//	Description: "All IPs contacted during page load. Contains values from both ip and request.ip",
	//	Example:     "1.1.1.1",
	//	Modifiers:   append(StandardModifiers, IPModifiers...),
	//},
	//{
	//	SigmaName:   "asns",
	//	Name:        "ASNs",
	//	Type:        StringList,
	//	Alias:       []string{"asn", "request.asn"},
	//	Description: "All ASNs contacted during page load. Contains values from both page.asn and request.asn",
	//	Example:     "13335",
	//	Modifiers:   StandardModifiers,
	//},

	// Fields relating to the primary page load (url, content, etc.)
	{
		SigmaName:   "url",
		Name:        "Page URL",
		Type:        String,
		GoType:      "*url.URL",
		Description: "The full URL of the page",
		Example:     "https://phish.domain/?foo=bar",
		Modifiers:   StandardModifiers,
		PostSetter: func(g *jen.Group) {
			g.Id("i").Dot("setURLHostname").Call(jen.Id("v").Dot("Hostname").Call())
			g.Id("i").Dot("setURLPath").Call(jen.Id("v").Dot("Path"))
			// TODO: URLQuery
		},
	},
	{
		SigmaName:   "url.hostname",
		Name:        "Page URL Hostname",
		Type:        String,
		Description: "The hostname of the page",
		Example:     "phish.domain",
		Modifiers:   slices.Concat(StandardModifiers, HostnameModifiers),
		Derived:     "url",
	},
	{
		SigmaName:  "hostname",
		Name:       "Page URL Hostname",
		Alias:      "url.hostname",
		Deprecated: true,
	},
	{
		SigmaName:   "url.path",
		Name:        "Page URL Path",
		Type:        String,
		Description: "Path component of the full URL of the page",
		Example:     "/foo/bar",
		Modifiers:   StandardModifiers,
		Derived:     "url",
	},
	{
		SigmaName:   "url.query",
		Name:        "Page URL Query",
		Type:        StringList,
		Description: "Query parameters extracted from the full URL of the page",
		Example:     "key=value",
		Modifiers:   StandardModifiers,
		Derived:     "url",
	},
	{
		SigmaName:   "ip",
		Name:        "Page IP",
		Type:        StringList,
		GoType:      "[]net.IP",
		Description: "IP addresses from which the page was loaded (resolved from url.hostname)",
		Example:     "1.1.1.1",
		Modifiers:   slices.Concat(StandardModifiers, IPModifiers),
	},
	{
		SigmaName:   "cname",
		Name:        "Page CNAMEs",
		Type:        StringList,
		Description: "Any CNAME records for the primary page URL, allowing you to detect things like pages hosted on GitHub",
		Example:     "github.io",
		Modifiers:   slices.Concat(StandardModifiers),
	},
	{
		SigmaName:   "asn",
		Name:        "Page ASN",
		Type:        NumberList,
		Description: "ASN from which the page was loaded",
		Example:     "13335",
		Modifiers:   slices.Concat(StandardModifiers, NumberModifiers),
		Derived:     "ip",
	},
	{
		SigmaName:   "registrar",
		Name:        "Registrar",
		Type:        Number,
		Description: "IANA ID for the registrar of this domain",
		Example:     "1910",
		Modifiers:   slices.Concat(StandardModifiers, NumberModifiers),
	},
	{
		SigmaName:   "status",
		Name:        "Page HTTP Status",
		Type:        Number,
		Description: "The HTTP status code of the primary page load",
		Example:     "200",
		Modifiers:   slices.Concat(StandardModifiers, NumberModifiers),
	},
	{
		SigmaName:   "header",
		Name:        "Page HTTP headers",
		Type:        StringList,
		Description: "HTTP Headers returned as part of page load. Each is in the canonical form Header-Name: value",
		Example:     "X-Powered-By: PHP/7.4.33",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:  "headers",
		Name:       "Page HTTP headers",
		Alias:      "header",
		Deprecated: true,
	},

	// Fields relating to the HTML content returned by the server
	{
		SigmaName:   "title",
		Name:        "Title",
		Type:        StringList,
		Description: "The title of the page as shown in a browser. If multiple titles are set (e.g. by JavaScript), this contains all of them",
		Example:     "Login | My Bank",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:  "titles",
		Name:       "Title",
		Alias:      "titles",
		Deprecated: true,
	},
	{
		SigmaName:   "html",
		Name:        "HTML",
		Type:        StringList,
		Description: "The contents of the page HTML. If JavaScript is used to modify the contents of the page, this contains multiple values",
		Example:     "<html><head>....</html>",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:  "dom",
		Name:       "DOM",
		Alias:      "html",
		Deprecated: true,
	},
	{
		SigmaName:   "js",
		Name:        "JS",
		Type:        StringList,
		Description: "Contents of JavaScript from the page (includes inline scripts as well as scripts loaded externally)",
		Example:     "!function(e,t){\"object\"==typeof module&&",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "css",
		Name:        "CSS",
		Type:        StringList,
		Description: "Contents of CSS from the page (includes inline stylesheets as well as externally loaded stylesheets)",
		Example:     ".redTitle {color: red;}",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "cookie",
		Name:        "Cookie",
		Type:        StringList,
		Description: "Cookies from the page. Each is in the form cookieName=value",
		Example:     "PHPSESSID=el4ukv0kqbvoirg7nkp4dncpk3",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:  "cookies",
		Name:       "Cookies",
		Alias:      "cookie",
		Deprecated: true,
	},
	{
		SigmaName:   "referrer",
		Name:        "Referrer",
		Type:        String,
		Description: "If this page is part of a redirect chain, the URL it redirected from",
		Example:     "https://www.phish.domain/initial",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "redirect",
		Name:        "Redirect URL",
		Type:        String,
		Description: "If this page redirected the browser to another page, the URL it redirected to",
		Example:     "https://www.phish.domain/see-other",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "redirect.kind",
		Name:        "Redirect Kind",
		Type:        String,
		Description: "If this page redirected the browser to another page, how was it redirected. One of: http{status}, js, meta",
		Example:     "http",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "favicon.hash",
		Name:        "Favicon Hash",
		Type:        StringList,
		Description: "SHA256 hash of favicons specified by the page. All favicons found in the page will be included.",
		Example:     "f39889eb42412593927c5136480e66f4ff0b813e071f2e5ddab70c14154692e4",
		Modifiers:   StandardModifiers,
	},

	// Fields relating to requests made outside the primary page load
	{
		SigmaName:   "request",
		Name:        "Requested URLs",
		Type:        StringList,
		GoType:      "[]*url.URL",
		Description: "URLs of requests made by the page (and assets loaded by the page)",
		Example:     "https://www.phish.domain/css/style.css",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:  "requests",
		Name:       "Requested URLs",
		Alias:      "request",
		Deprecated: true,
	},
	{
		SigmaName:   "request.hostname",
		Name:        "Requested Hostnames",
		Type:        StringList,
		Description: "Hostnames of requests made by the page (and assets loaded by the page)",
		Example:     "www.phish.domain",
		Modifiers:   slices.Concat(StandardModifiers, HostnameModifiers),
		Derived:     "requests",
	},
	{
		SigmaName:   "request.ip",
		Name:        "Requested IPs",
		Type:        StringList,
		GoType:      "[]net.IP",
		Description: "IPs contacted as part of requests made by the page",
		Example:     "1.1.1.1",
		Modifiers:   slices.Concat(StandardModifiers, IPModifiers),
	},
	{
		SigmaName:   "request.asn",
		Name:        "Requested ASNs",
		Type:        NumberList,
		Description: "ASNs contacted as part of requests made by the page",
		Example:     "13335",
		Modifiers:   slices.Concat(StandardModifiers, NumberModifiers),
		Derived:     "requests.ip",
	},
	{
		SigmaName:   "request.header",
		Name:        "Request headers",
		Type:        StringList,
		Description: "Headers in requests made during the page load. Each is in the canonical form Header-Name: value",
		Example:     "Authentication: Bearer el4ukv0kqbvoirg7nkp4dncpk3",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "response.hash",
		Name:        "Response hash",
		Type:        StringList,
		Description: "SHA256 Hashes of response bodies",
		Example:     "f39889eb42412593927c5136480e66f4ff0b813e071f2e5ddab70c14154692e4",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "response.header",
		Name:        "Response headers",
		Type:        StringList,
		Description: "Headers returned as part of responses. Each is in the canonical form Header-Name: value",
		Example:     "X-Powered-By: PHP/7.4.33",
		Modifiers:   StandardModifiers,
	},

	// Meta fields
	{
		SigmaName:   "supported",
		Name:        "Supported fields",
		Type:        StringList,
		Description: "Fields supported by this IOK implementation. This allows you to distinguish between fields which are genuinely empty (e.g. the page had no <title> element) from cases where the data was not loaded.",
		Example:     "url.query",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "unsupported",
		Name:        "Unsupported fields",
		Type:        StringList,
		Description: "Fields *not* supported by this IOK implementation. For example, the favicon may not have been fetched and so the favicon.hash field will be unset.",
		Example:     "favicon.hash",
		Modifiers:   StandardModifiers,
		Derived:     "supported",
	},
}

var StandardModifiers = []Modifier{
	{
		Name:        "contains",
		Description: "True if the field contains the string",
	},
	{
		Name:        "startswith",
		Description: "True if the field starts with the string",
	},
	{
		Name:        "endswith",
		Description: "True if the field ends with the string",
	},
	{
		Name:        "all",
		Description: "Requires this matches all provided values, not just one",
	},
	{
		Name:        "re",
		Description: "True if the field matches the given regular expression",
	},
}

var NumberModifiers = []Modifier{
	{
		Name:        "lt",
		Description: "True if the number is less than the threshold",
	},
	{
		Name:        "lte",
		Description: "True if the number is less than or equal the threshold",
	},
	{
		Name:        "gt",
		Description: "True if the number is greater than the threshold",
	},
	{
		Name:        "gte",
		Description: "True if the number is greater than or equal the threshold",
	},
}

var HostnameModifiers = []Modifier{
	{
		Name:        "subdomain",
		Description: "True if the hostname is equal to or a subdomain of the given hostname",
	},
}

var IPModifiers = []Modifier{
	{
		Name:        "cidr",
		Description: "True if the IP is within the given CIDR",
	},
}

var Modifiers = slices.Concat(StandardModifiers, HostnameModifiers, IPModifiers)
