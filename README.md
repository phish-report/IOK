<p align="center">
  <a href="https://phish.report/IOK">
  <img
    width="400"
    src="https://raw.githubusercontent.com/phish-report/IOK/main/.github/logo.svg"
    alt="IOK logo"
       /></a>
</p>
<p align="center">
  <a href="https://phish.report/IOK">View detections on phish.report 🐟</a>
</p>
<hr>
<img
src="https://raw.githubusercontent.com/phish-report/IOK/main/.github/fake-chrome-error.yml.svg"
alt="Screenshot of one of the IOK indicator rules"
width="50%"
align="right"
/>

**[Indicator of Kit](https://phish.report) is an open source detection language for phishing site techniques, kits, and threat actors 🕵️**

- **Simple:** based on [Sigma](https://github.com/sigmahq/Sigma), a simple detection rules language 🚀
- **Rich metadata:** rules have descriptions, tags, and links to blog posts or related rules.

**Use cases:**

- [Identify fingerprints of known threat actors](https://github.com/phish-report/IOK/blob/main/indicators/cazanova-cookie.yml)
- [Discover anti-analysis techniques](https://github.com/phish-report/IOK/blob/main/indicators/fake-chrome-error.yml)
- [Classify which specific phishing kit is in use on a page](https://github.com/phish-report/IOK/blob/main/indicators/123-reg-63c26.yml)
- [Identify deceptive websites dropping malicious software](https://github.com/phish-report/IOK/blob/main/indicators/bbystealer-family-dropper-website-7019ae4.yml)
- [Discover APT infrastructure](https://github.com/phish-report/IOK/blob/main/indicators/kimsuky-nginx-fake-error-9b43f670.yml)
- [Detect malware C&C panels](https://github.com/phish-report/IOK/blob/main/indicators/rhadamanthys-stealer-26461dbb.yml)


## 📝 Creating indicators

IOK indicators are written using [Sigma](https://github.com/SigmaHQ/sigma)

| Field name |   Type   | Description                                                                                                           |
|:----------:|:--------:|-----------------------------------------------------------------------------------------------------------------------|
|   title    | []string | The title of the site as shown in a browser. If multiple titles are set (e.g. by JavaScript), this contains each one. |
|  hostname  |  string  | The hostname of the site                                                                                              |
|    html    |  string  | The contents of the page HTML (as returned by the server)                                                             |
|    dom     |  string  | The contents of the page HTML after loading (e.g. after javascript has executed)                                      |
|     js     | []string | Contents of JavaScript from the page (includes inline scripts as well as scripts loaded externally)                   |
|    css     | []string | Contents of CSS from the page (includes inline stylesheets as well as externally loaded stylesheets)                  |
|  cookies   | []string | Cookies from the page. Each is in the form `cookieName=value`                                                         |
|  headers   | []string | Headers sent by the server. Each is in the form `Header-Name: value`                                                  |
|  requests  | []string | URLs of requests made by the page (and assets loaded by the page)                                                     |

We are always looking for contributions: there's far more phishing kits and techniques than a single team can analyse!

To contribute a new rule:

1. Try to make sure it doesn't already exist
2. Open a pull request, adding your new file in the `indicators/` folder
3. We'll review it and merge your PR
4. It'll go live on [phish.report/IOK](https://phish.report/IOK)!

## 💭 Comparison to similar projects

|                                       |          IOK          | [PhishingKit-Yara-Rules] |   [Wappalyzer]    |
|---------------------------------------|:---------------------:|:------------------------:|:-----------------:|
| Open Source                           |           ✅           |            ✅             |         ✅         |
| Ruleset size                          | &gt; 215 Rules 🦐 |     500 rules 🐠     | 1000s of rules 🐳 |
| Can scan                              |   Live websites 🕸    |   Phishing kit zips 📦   | Live websites 🕸  |
| Phishing focused                      |           ✅           |            ✅             |         ❌         |
| Supports complex conditions           |           ✅           |            ✅             |         ❌         |
| Sends out stickers to contributors 🎁 |           ✅           |            ❌             |         ❌         |

[PhishingKit-Yara-Rules]: https://github.com/t4d/PhishingKit-Yara-Rules

[Wappalyzer]: https://www.wappalyzer.com/

## 🤝 Contributing

There's a [reference on how to write IOK rules](https://phish.report/docs/iok-rule-reference) in the Phish Report documentation.

## 📝 License

This project is [ODbL](https://github.com/phish-report/IOK/blob/main/LICENSE) licensed.
You're free to use the rules in your own projects (including commercial ones!)
as long as you credit [phish.report/IOK](https://phish.report/IOK) as the source.

For more details, read [OpenStreetMap's guidance](https://wiki.openstreetmap.org/wiki/License/Use_Cases) (who also use
the ODbL license).
