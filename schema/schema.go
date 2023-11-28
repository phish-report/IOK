package schema

type FieldType string

var (
	StringField     FieldType = "string"
	StringListField FieldType = "string list"
)

var Fields = []Field{
	// Combo fields
	{
		SigmaName:   "url",
		Name:        "URLs",
		Type:        StringListField,
		Description: "URLs of all requests made during page load. Contains values from both page.url and request.url",
		Example:     "https://phish.domain/?foo=bar",
		Modifiers:   append(StandardModifiers, URLModifiers...),
	},
	{
		SigmaName:   "ip",
		Name:        "IPs",
		Type:        StringListField,
		Description: "All IPs contacted during page load. Contains values from both page.ip and request.ip",
		Example:     "1.1.1.1",
		Modifiers:   append(StandardModifiers, IPModifiers...),
	},
	{
		SigmaName:   "ASN",
		Name:        "ASNs",
		Type:        StringListField,
		Description: "All ASNs contacted during page load. Contains values from both page.asn and request.asn",
		Example:     "13335",
		Modifiers:   StandardModifiers,
	},

	// Fields relating to the primary page load (url, content, etc.)
	{
		SigmaName:   "page.url",
		Name:        "Page URL",
		Type:        StringField,
		Description: "The full URL of the page",
		Example:     "https://phish.domain/?foo=bar",
		Modifiers:   append(StandardModifiers, URLModifiers...),
	},
	{
		SigmaName:   "page.status",
		Name:        "Page HTTP Status",
		Type:        StringField,
		Description: "The HTTP status code of the primary page load",
		Example:     "200",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "title",
		Name:        "Title",
		Type:        StringListField,
		Description: "The title of the page as shown in a browser. If multiple titles are set (e.g. by JavaScript), this contains all of them.",
		Example:     "Login | My Bank",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "html",
		Name:        "HTML",
		Type:        StringListField,
		Description: "The contents of the page HTML. If JavaScript is used to modify the contents of the page, this contains multiple values.",
		Example:     "<html><head>....</html>",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "js",
		Name:        "JS",
		Type:        StringListField,
		Description: "Contents of JavaScript from the page (includes inline scripts as well as scripts loaded externally).",
		Example:     "!function(e,t){\"object\"==typeof module&&",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "css",
		Name:        "CSS",
		Type:        StringListField,
		Description: "Contents of CSS from the page (includes inline stylesheets as well as externally loaded stylesheets)",
		Example:     ".redTitle {color: red;}",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "cookies",
		Name:        "Cookies",
		Type:        StringListField,
		Description: "Cookies from the page. Each is in the form cookieName=value",
		Example:     "PHPSESSID=el4ukv0kqbvoirg7nkp4dncpk3",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "page.ip",
		Name:        "Page IP",
		Type:        StringField,
		Description: "IP address from which the page was loaded.",
		Example:     "1.1.1.1",
		Modifiers:   append(StandardModifiers, IPModifiers...),
	},
	{
		SigmaName:   "page.asn",
		Name:        "Page ASN",
		Type:        StringField,
		Description: "ASN from which the page was loaded.",
		Example:     "13335",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "referrer",
		Name:        "Referrer",
		Type:        StringField,
		Description: "If this page is part of a redirect chain, the URL it redirected from.",
		Example:     "https://www.phish.domain/initial",
		Modifiers:   append(StandardModifiers, URLModifiers...),
	},
	{
		SigmaName:   "redirect.url",
		Name:        "Redirect URL",
		Type:        StringField,
		Description: "If this page redirected the browser to another page, the URL it redirected to.",
		Example:     "https://www.phish.domain/see-other",
		Modifiers:   append(StandardModifiers, URLModifiers...),
	},
	{
		SigmaName:   "redirect.kind",
		Name:        "Redirect Kind",
		Type:        StringField,
		Description: "If this page redirected the browser to another page, how was it redirected. One of: http, js, meta",
		Example:     "http",
		Modifiers:   StandardModifiers,
	},

	// Fields relating to requests made outside the primary page load
	{
		SigmaName:   "request.url",
		Name:        "Requested URLs",
		Type:        StringListField,
		Description: "URLs of requests made by the page (and assets loaded by the page).",
		Example:     "https://www.phish.domain/css/style.css",
		Modifiers:   append(StandardModifiers, URLModifiers...),
	},
	{
		SigmaName:   "request.ip",
		Name:        "Requested IPs",
		Type:        StringListField,
		Description: "IPs contacted as part of requests made by the page.",
		Example:     "1.1.1.1",
		Modifiers:   append(StandardModifiers, IPModifiers...),
	},
	{
		SigmaName:   "request.asn",
		Name:        "Requested ASNs",
		Type:        StringListField,
		Description: "ASNs contacted as part of requests made by the page.",
		Example:     "13335",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "request.header",
		Name:        "Request headers",
		Type:        StringListField,
		Description: "Headers in requests made during the page load. Each is in the canonical form Header-Name: value",
		Example:     "Authentication: Bearer el4ukv0kqbvoirg7nkp4dncpk3",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "response.header",
		Name:        "Response headers",
		Type:        StringListField,
		Description: "Headers returned by the server. Each is in the canonical form Header-Name: value",
		Example:     "X-Powered-By: PHP/7.4.33",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "response.hash",
		Name:        "Response hash",
		Type:        StringListField,
		Description: "SHA256 Hashes of response bodies.",
		Example:     "f39889eb42412593927c5136480e66f4ff0b813e071f2e5ddab70c14154692e4",
		Modifiers:   StandardModifiers,
	},
}

type Field struct {
	SigmaName   string
	Name        string
	Type        FieldType
	Description string
	Example     string
	Modifiers   []Modifier
}

type Modifier struct {
	Name        string
	Description string
}

var StandardModifiers = []Modifier{
	{
		Name:        "contains",
		Description: "True if the field contains the string.",
	},
	{
		Name:        "startswith",
		Description: "True if the field starts with the string.",
	},
	{
		Name:        "endswith",
		Description: "True if the field ends with the string.",
	},
	{
		Name:        "all",
		Description: "Requires this matches all provided values, not just one.",
	},
	{
		Name:        "re",
		Description: "True if the field matches the given regular expression",
	},
}

var URLModifiers = []Modifier{
	{
		Name:        "hostname",
		Description: "Extracts the hostname from the URL",
	},
	{
		Name:        "subdomain",
		Description: "True if the hostname is equal to or a subdomain of the given hostname",
	},
	{
		Name:        "query",
		Description: "Extracts just the query parameters from the URL in the form key=value",
	},
	{
		Name:        "path",
		Description: "Extracts just the path from the URL",
	},
}

var IPModifiers = []Modifier{
	{
		Name:        "cidr",
		Description: "True if the IP is within the given CIDR",
	},
}

var Modifiers = append(StandardModifiers)
