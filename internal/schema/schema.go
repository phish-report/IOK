package schema

type FieldType string

var (
	StringField     FieldType = "string"
	StringListField FieldType = "string list"
)

var Fields = []Field{
	// Fields relating to the loaded url
	{
		SigmaName:   "url",
		Name:        "URL",
		Type:        StringField,
		Description: "The full URL of the site",
		Example:     "https://phish.domain/?foo=bar",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "hostname",
		Name:        "Hostname",
		Type:        StringField,
		Description: "The hostname of the site.",
		Example:     "www.phish.domain",
		Modifiers:   append(StandardModifiers, SubdomainModifier),
	},
	{
		SigmaName:   "query",
		Name:        "Query Parameters",
		Type:        StringListField,
		Description: "The GET/query parameters in the URL, in the form key=value.",
		Example:     "key=value",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "path",
		Name:        "Path",
		Type:        StringField,
		Description: "The path component of the URL, in the form key=value.",
		Example:     "/admin/login",
		Modifiers:   StandardModifiers,
	},

	// HTTP-level fields: response codes, headers, etc.
	{
		SigmaName:   "headers",
		Name:        "Headers",
		Type:        StringListField,
		Description: "Headers returned by the server. Each is in the form Header-Name: value",
		Example:     "X-Powered-By: PHP/7.4.33",
		Modifiers:   StandardModifiers,
	},

	{
		SigmaName:   "title",
		Name:        "Title",
		Type:        StringListField,
		Description: "The title of the site as shown in a browser. If multiple titles are set (e.g. by JavaScript), this contains all of them.",
		Example:     "Login | My Bank",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "html",
		Name:        "HTML",
		Type:        StringField,
		Description: "The contents of the page HTML (as returned by the server).",
		Example:     "<html><head>....</html>",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "dom",
		Name:        "DOM",
		Type:        StringField,
		Description: "The contents of the page HTML after loading (e.g. after javascript has executed and modified the page).",
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
		SigmaName:   "headers",
		Name:        "Headers",
		Type:        StringListField,
		Description: "Headers sent by the server. Each is in the form Header-Name: value",
		Example:     "X-Powered-By: PHP/7.4.33",
		Modifiers:   StandardModifiers,
	},
	{
		SigmaName:   "requests",
		Name:        "Requests",
		Type:        StringListField,
		Description: "URLs of requests made by the page (and assets loaded by the page).",
		Example:     "https://www.phish.domain/css/style.css",
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

var SubdomainModifier = Modifier{
	Name:        "subdomain",
	Description: "True if the hostname is equal to or a subdomain of the given hostname",
}

var Modifiers = append(StandardModifiers)
