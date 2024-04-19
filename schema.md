| Field | Description | Example |
|-------|-------------|---------|
|`url`|The full URL of the page|`https://phish.domain/?foo=bar`|
|`url.hostname`|The hostname of the page|`phish.domain`|
|`url.path`|Path component of the full URL of the page|`/foo/bar`|
|`url.query`|Query parameters extracted from the full URL of the page|`key=value`|
|`ip`|IP addresses from which the page was loaded (resolved from url.hostname)|`1.1.1.1`|
|`asn`|ASN from which the page was loaded|`13335`|
|`registrar`|IANA ID for the registrar of this domain|`1910`|
|`status`|The HTTP status code of the primary page load|`200`|
|`header`|HTTP Headers returned as part of page load. Each is in the canonical form Header-Name: value|`X-Powered-By: PHP/7.4.33`|
|`title`|The title of the page as shown in a browser. If multiple titles are set (e.g. by JavaScript), this contains all of them|`Login | My Bank`|
|`html`|The contents of the page HTML. If JavaScript is used to modify the contents of the page, this contains multiple values|`<html><head>....</html>`|
|`js`|Contents of JavaScript from the page (includes inline scripts as well as scripts loaded externally)|`!function(e,t){"object"==typeof module&&`|
|`css`|Contents of CSS from the page (includes inline stylesheets as well as externally loaded stylesheets)|`.redTitle {color: red;}`|
|`cookie`|Cookies from the page. Each is in the form cookieName=value|`PHPSESSID=el4ukv0kqbvoirg7nkp4dncpk3`|
|`referrer`|If this page is part of a redirect chain, the URL it redirected from|`https://www.phish.domain/initial`|
|`redirect`|If this page redirected the browser to another page, the URL it redirected to|`https://www.phish.domain/see-other`|
|`redirect.kind`|If this page redirected the browser to another page, how was it redirected. One of: http{status}, js, meta|`http`|
|`favicon.hash`|SHA256 hash of favicons specified by the page. All favicons found in the page will be included.|`f39889eb42412593927c5136480e66f4ff0b813e071f2e5ddab70c14154692e4`|
|`requests`|URLs of requests made by the page (and assets loaded by the page)|`https://www.phish.domain/css/style.css`|
|`requests.hostname`|Hostnames of requests made by the page (and assets loaded by the page)|`www.phish.domain`|
|`requests.ip`|IPs contacted as part of requests made by the page|`1.1.1.1`|
|`requests.asn`|ASNs contacted as part of requests made by the page|`13335`|
|`request.header`|Headers in requests made during the page load. Each is in the canonical form Header-Name: value|`Authentication: Bearer el4ukv0kqbvoirg7nkp4dncpk3`|
|`response.hash`|SHA256 Hashes of response bodies|`f39889eb42412593927c5136480e66f4ff0b813e071f2e5ddab70c14154692e4`|
|`response.header`|Headers returned as part of responses. Each is in the canonical form Header-Name: value|`X-Powered-By: PHP/7.4.33`|
|`supported`|Fields supported by this IOK implementation|`url.query`|
|`unsupported`|Fields *not* supported by this IOK implementation|`url.query`|
