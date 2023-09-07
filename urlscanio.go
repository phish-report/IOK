package iok

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"time"

	"golang.org/x/net/html"
)

type urlscanResult struct {
	DOM  []byte
	Data struct {
		Requests []struct {
			Request struct {
				RequestId   string `json:"requestId"`
				LoaderId    string `json:"loaderId"`
				DocumentURL string `json:"documentURL"`
				Request     struct {
					Url     string `json:"url"`
					Method  string `json:"method"`
					Headers struct {
						UpgradeInsecureRequests string `json:"Upgrade-Insecure-Requests,omitempty"`
						UserAgent               string `json:"User-Agent"`
						AcceptLanguage          string `json:"accept-language"`
						Referer                 string `json:"Referer,omitempty"`
						Origin                  string `json:"Origin,omitempty"`
					} `json:"headers"`
					MixedContentType string `json:"mixedContentType"`
					InitialPriority  string `json:"initialPriority"`
					ReferrerPolicy   string `json:"referrerPolicy"`
					IsSameSite       bool   `json:"isSameSite"`
					IsLinkPreload    bool   `json:"isLinkPreload,omitempty"`
				} `json:"request"`
				Timestamp float64 `json:"timestamp"`
				WallTime  float64 `json:"wallTime"`
				Initiator struct {
					Type         string `json:"type"`
					Url          string `json:"url,omitempty"`
					LineNumber   int    `json:"lineNumber,omitempty"`
					ColumnNumber int    `json:"columnNumber,omitempty"`
					Stack        struct {
						CallFrames []struct {
							FunctionName string `json:"functionName"`
							ScriptId     string `json:"scriptId"`
							Url          string `json:"url"`
							LineNumber   int    `json:"lineNumber"`
							ColumnNumber int    `json:"columnNumber"`
						} `json:"callFrames"`
					} `json:"stack,omitempty"`
				} `json:"initiator"`
				RedirectHasExtraInfo bool   `json:"redirectHasExtraInfo"`
				Type                 string `json:"type"`
				FrameId              string `json:"frameId"`
				HasUserGesture       bool   `json:"hasUserGesture"`
				PrimaryRequest       bool   `json:"primaryRequest,omitempty"`
			} `json:"request"`
			Response struct {
				EncodedDataLength int    `json:"encodedDataLength"`
				DataLength        int    `json:"dataLength"`
				RequestId         string `json:"requestId"`
				Type              string `json:"type"`
				Response          struct {
					Url               string            `json:"url"`
					Status            int               `json:"status"`
					StatusText        string            `json:"statusText"`
					Headers           map[string]string `json:"headers"`
					MimeType          string            `json:"mimeType"`
					RemoteIPAddress   string            `json:"remoteIPAddress"`
					RemotePort        int               `json:"remotePort"`
					EncodedDataLength int               `json:"encodedDataLength"`
					Timing            struct {
						RequestTime              float64 `json:"requestTime"`
						ProxyStart               int     `json:"proxyStart"`
						ProxyEnd                 int     `json:"proxyEnd"`
						DnsStart                 float64 `json:"dnsStart"`
						DnsEnd                   float64 `json:"dnsEnd"`
						ConnectStart             float64 `json:"connectStart"`
						ConnectEnd               float64 `json:"connectEnd"`
						SslStart                 float64 `json:"sslStart"`
						SslEnd                   float64 `json:"sslEnd"`
						WorkerStart              int     `json:"workerStart"`
						WorkerReady              int     `json:"workerReady"`
						WorkerFetchStart         int     `json:"workerFetchStart"`
						WorkerRespondWithSettled int     `json:"workerRespondWithSettled"`
						SendStart                float64 `json:"sendStart"`
						SendEnd                  float64 `json:"sendEnd"`
						PushStart                int     `json:"pushStart"`
						PushEnd                  int     `json:"pushEnd"`
						ReceiveHeadersEnd        float64 `json:"receiveHeadersEnd"`
					} `json:"timing"`
					ResponseTime    float64 `json:"responseTime"`
					Protocol        string  `json:"protocol"`
					SecurityState   string  `json:"securityState"`
					SecurityDetails struct {
						Protocol                          string        `json:"protocol"`
						KeyExchange                       string        `json:"keyExchange"`
						KeyExchangeGroup                  string        `json:"keyExchangeGroup"`
						Cipher                            string        `json:"cipher"`
						CertificateId                     int           `json:"certificateId"`
						SubjectName                       string        `json:"subjectName"`
						SanList                           []string      `json:"sanList"`
						Issuer                            string        `json:"issuer"`
						ValidFrom                         int           `json:"validFrom"`
						ValidTo                           int           `json:"validTo"`
						SignedCertificateTimestampList    []interface{} `json:"signedCertificateTimestampList"`
						CertificateTransparencyCompliance string        `json:"certificateTransparencyCompliance"`
					} `json:"securityDetails"`
					SecurityHeaders []struct {
						Name  string `json:"name"`
						Value string `json:"value"`
					} `json:"securityHeaders,omitempty"`
				} `json:"response"`
				HasExtraInfo bool   `json:"hasExtraInfo"`
				Hash         string `json:"hash"`
				Size         int    `json:"size"`
				Asn          struct {
					Ip          string `json:"ip"`
					Asn         string `json:"asn"`
					Country     string `json:"country"`
					Registrar   string `json:"registrar"`
					Date        string `json:"date"`
					Description string `json:"description"`
					Route       string `json:"route"`
					Name        string `json:"name"`
				} `json:"asn"`
				Geoip struct {
					Country     string    `json:"country"`
					Region      string    `json:"region"`
					Timezone    string    `json:"timezone"`
					City        string    `json:"city"`
					Ll          []float64 `json:"ll"`
					CountryName string    `json:"country_name"`
					Metro       int       `json:"metro"`
				} `json:"geoip"`
				Rdns struct {
					Ip  string `json:"ip"`
					Ptr string `json:"ptr"`
				} `json:"rdns,omitempty"`
			} `json:"response"`
			InitiatorInfo struct {
				Url  string `json:"url"`
				Host string `json:"host"`
				Type string `json:"type"`
			} `json:"initiatorInfo,omitempty"`
		} `json:"requests"`
		Cookies []struct {
			Name         string  `json:"name"`
			Value        string  `json:"value"`
			Domain       string  `json:"domain"`
			Path         string  `json:"path"`
			Expires      float64 `json:"expires"`
			Size         int     `json:"size"`
			HttpOnly     bool    `json:"httpOnly"`
			Secure       bool    `json:"secure"`
			Session      bool    `json:"session"`
			Priority     string  `json:"priority"`
			SameParty    bool    `json:"sameParty"`
			SourceScheme string  `json:"sourceScheme"`
			SourcePort   int     `json:"sourcePort"`
		} `json:"cookies"`
		Console []struct {
			Message struct {
				Source    string      `json:"source"`
				Level     string      `json:"level"`
				Text      string      `json:"text"`
				Timestamp float64     `json:"timestamp"`
				Url       string      `json:"url"`
				Line      interface{} `json:"line"`
			} `json:"message"`
		} `json:"console"`
		Links []struct {
			Href string `json:"href"`
			Text string `json:"text"`
		} `json:"links"`
		Timing struct {
			BeginNavigation      time.Time `json:"beginNavigation"`
			FrameStartedLoading  time.Time `json:"frameStartedLoading"`
			FrameNavigated       time.Time `json:"frameNavigated"`
			DomContentEventFired time.Time `json:"domContentEventFired"`
			FrameStoppedLoading  time.Time `json:"frameStoppedLoading"`
		} `json:"timing"`
		Globals []struct {
			Prop string `json:"prop"`
			Type string `json:"type"`
		} `json:"globals"`
	} `json:"data"`
	Stats struct {
		ResourceStats []struct {
			Count       int      `json:"count"`
			Size        int      `json:"size"`
			EncodedSize int      `json:"encodedSize"`
			Latency     int      `json:"latency"`
			Countries   []string `json:"countries"`
			Ips         []string `json:"ips"`
			Type        string   `json:"type"`
			Compression string   `json:"compression"`
			Percentage  int      `json:"percentage"`
		} `json:"resourceStats"`
		ProtocolStats []struct {
			Count         int      `json:"count"`
			Size          int      `json:"size"`
			EncodedSize   int      `json:"encodedSize"`
			Ips           []string `json:"ips"`
			Countries     []string `json:"countries"`
			SecurityState struct {
			} `json:"securityState"`
			Protocol string `json:"protocol"`
		} `json:"protocolStats"`
		TlsStats []struct {
			Count       int      `json:"count"`
			Size        int      `json:"size"`
			EncodedSize int      `json:"encodedSize"`
			Ips         []string `json:"ips"`
			Countries   []string `json:"countries"`
			Protocols   struct {
				TLS13AES128GCM int `json:"TLS 1.3 /  / AES_128_GCM"`
			} `json:"protocols"`
			SecurityState string `json:"securityState"`
		} `json:"tlsStats"`
		ServerStats []struct {
			Count       int      `json:"count"`
			Size        int      `json:"size"`
			EncodedSize int      `json:"encodedSize"`
			Ips         []string `json:"ips"`
			Countries   []string `json:"countries"`
			Server      string   `json:"server"`
		} `json:"serverStats"`
		DomainStats []struct {
			Count       int      `json:"count"`
			Ips         []string `json:"ips"`
			Domain      string   `json:"domain"`
			Size        int      `json:"size"`
			EncodedSize int      `json:"encodedSize"`
			Countries   []string `json:"countries"`
			Index       int      `json:"index"`
			Initiators  []string `json:"initiators"`
			Redirects   int      `json:"redirects"`
		} `json:"domainStats"`
		RegDomainStats []struct {
			Count       int           `json:"count"`
			Ips         []string      `json:"ips"`
			RegDomain   string        `json:"regDomain"`
			Size        int           `json:"size"`
			EncodedSize int           `json:"encodedSize"`
			Countries   []interface{} `json:"countries"`
			Index       int           `json:"index"`
			SubDomains  []struct {
				Domain  string `json:"domain"`
				Country string `json:"country"`
			} `json:"subDomains"`
			Redirects int `json:"redirects"`
		} `json:"regDomainStats"`
		SecureRequests   int `json:"secureRequests"`
		SecurePercentage int `json:"securePercentage"`
		IPv6Percentage   int `json:"IPv6Percentage"`
		UniqCountries    int `json:"uniqCountries"`
		TotalLinks       int `json:"totalLinks"`
		Malicious        int `json:"malicious"`
		AdBlocked        int `json:"adBlocked"`
		IpStats          []struct {
			Requests int      `json:"requests"`
			Domains  []string `json:"domains"`
			Ip       string   `json:"ip"`
			Asn      struct {
				Ip          string `json:"ip"`
				Asn         string `json:"asn"`
				Country     string `json:"country"`
				Registrar   string `json:"registrar"`
				Date        string `json:"date"`
				Description string `json:"description"`
				Route       string `json:"route"`
				Name        string `json:"name"`
			} `json:"asn"`
			Dns struct {
			} `json:"dns"`
			Geoip struct {
				Country     string    `json:"country"`
				Region      string    `json:"region"`
				Timezone    string    `json:"timezone"`
				City        string    `json:"city"`
				Ll          []float64 `json:"ll"`
				CountryName string    `json:"country_name"`
				Metro       int       `json:"metro"`
			} `json:"geoip"`
			Size        int         `json:"size"`
			EncodedSize int         `json:"encodedSize"`
			Countries   []string    `json:"countries"`
			Index       int         `json:"index"`
			Ipv6        bool        `json:"ipv6"`
			Redirects   int         `json:"redirects"`
			Count       interface{} `json:"count"`
			Rdns        struct {
				Ip  string `json:"ip"`
				Ptr string `json:"ptr"`
			} `json:"rdns,omitempty"`
		} `json:"ipStats"`
	} `json:"stats"`
	Meta struct {
		Processors struct {
			Umbrella struct {
				Data []struct {
					Hostname string `json:"hostname"`
					Rank     int    `json:"rank"`
				} `json:"data"`
			} `json:"umbrella"`
			Geoip struct {
				Data []struct {
					Ip    string `json:"ip"`
					Geoip struct {
						Country     string    `json:"country"`
						Region      string    `json:"region"`
						Timezone    string    `json:"timezone"`
						City        string    `json:"city"`
						Ll          []float64 `json:"ll"`
						CountryName string    `json:"country_name"`
						Metro       int       `json:"metro"`
					} `json:"geoip"`
				} `json:"data"`
			} `json:"geoip"`
			Rdns struct {
				Data []struct {
					Ip  string `json:"ip"`
					Ptr string `json:"ptr"`
				} `json:"data"`
			} `json:"rdns"`
			Asn struct {
				Data []struct {
					Ip          string `json:"ip"`
					Asn         string `json:"asn"`
					Country     string `json:"country"`
					Registrar   string `json:"registrar"`
					Date        string `json:"date"`
					Description string `json:"description"`
					Route       string `json:"route"`
					Name        string `json:"name"`
				} `json:"data"`
			} `json:"asn"`
			Wappa struct {
				Data []struct {
					Confidence []struct {
						Confidence int    `json:"confidence"`
						Pattern    string `json:"pattern"`
					} `json:"confidence"`
					ConfidenceTotal int    `json:"confidenceTotal"`
					App             string `json:"app"`
					Icon            string `json:"icon"`
					Website         string `json:"website"`
					Categories      []struct {
						Name     string `json:"name"`
						Priority int    `json:"priority"`
					} `json:"categories"`
				} `json:"data"`
			} `json:"wappa"`
		} `json:"processors"`
	} `json:"meta"`
	Task struct {
		Uuid          string        `json:"uuid"`
		Time          time.Time     `json:"time"`
		Url           string        `json:"url"`
		Visibility    string        `json:"visibility"`
		Method        string        `json:"method"`
		Source        string        `json:"source"`
		Tags          []interface{} `json:"tags"`
		ReportURL     string        `json:"reportURL"`
		ScreenshotURL string        `json:"screenshotURL"`
		DomURL        string        `json:"domURL"`
	} `json:"task"`
	Page struct {
		Url     string `json:"url"`
		Domain  string `json:"domain"`
		Country string `json:"country"`
		City    string `json:"city"`
		Server  string `json:"server"`
		Ip      string `json:"ip"`
		Asn     string `json:"asn"`
		Asnname string `json:"asnname"`
	} `json:"page"`
	Lists struct {
		Ips          []string `json:"ips"`
		Countries    []string `json:"countries"`
		Asns         []string `json:"asns"`
		Domains      []string `json:"domains"`
		Servers      []string `json:"servers"`
		Urls         []string `json:"urls"`
		LinkDomains  []string `json:"linkDomains"`
		Certificates []struct {
			SubjectName string `json:"subjectName"`
			Issuer      string `json:"issuer"`
			ValidFrom   int    `json:"validFrom"`
			ValidTo     int    `json:"validTo"`
		} `json:"certificates"`
		Hashes []string `json:"hashes"`
	} `json:"lists"`
	Verdicts struct {
		Overall struct {
			Score       int           `json:"score"`
			Categories  []interface{} `json:"categories"`
			Brands      []interface{} `json:"brands"`
			Tags        []interface{} `json:"tags"`
			Malicious   bool          `json:"malicious"`
			HasVerdicts bool          `json:"hasVerdicts"`
		} `json:"overall"`
		Urlscan struct {
			Score       int           `json:"score"`
			Categories  []interface{} `json:"categories"`
			Brands      []interface{} `json:"brands"`
			Tags        []interface{} `json:"tags"`
			Malicious   bool          `json:"malicious"`
			HasVerdicts bool          `json:"hasVerdicts"`
		} `json:"urlscan"`
		Engines struct {
			Score             int           `json:"score"`
			Categories        []interface{} `json:"categories"`
			EnginesTotal      int           `json:"enginesTotal"`
			MaliciousTotal    int           `json:"maliciousTotal"`
			BenignTotal       int           `json:"benignTotal"`
			MaliciousVerdicts []interface{} `json:"maliciousVerdicts"`
			BenignVerdicts    []interface{} `json:"benignVerdicts"`
			Malicious         bool          `json:"malicious"`
			HasVerdicts       bool          `json:"hasVerdicts"`
		} `json:"engines"`
		Community struct {
			Score          int           `json:"score"`
			Categories     []interface{} `json:"categories"`
			Brands         []interface{} `json:"brands"`
			VotesTotal     int           `json:"votesTotal"`
			VotesMalicious int           `json:"votesMalicious"`
			VotesBenign    int           `json:"votesBenign"`
			Malicious      bool          `json:"malicious"`
			HasVerdicts    bool          `json:"hasVerdicts"`
		} `json:"community"`
	} `json:"verdicts"`
	Submitter struct {
		Country string `json:"country"`
	} `json:"submitter"`
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// InputFromURLScan takes a urlscan.io result ID and returns an Input suitable for calling GetMatches with.
// The provided http.Client should inject your API key if you have one.
func InputFromURLScan(ctx context.Context, urlscanUUID string, client httpClient) (Input, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://urlscan.io/api/v1/result/"+urlscanUUID, nil)
	resp, err := client.Do(req)
	if err != nil {
		return Input{}, fmt.Errorf("failed to get search result: %w", err)
	}
	defer resp.Body.Close()

	result := urlscanResult{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Input{}, fmt.Errorf("failed to decode search result: %w", err)
	}

	input := Input{}
	u, err := url.Parse(result.Page.Url)
	if err != nil {
		return Input{}, fmt.Errorf("failed to parse result URL: %w", err)
	}
	input.Hostname = u.Hostname()

	domReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://urlscan.io/dom/"+result.Task.Uuid, nil)
	domResp, err := client.Do(domReq)
	if err != nil {
		return Input{}, fmt.Errorf("failed to get result html: %w", err)
	}
	defer domResp.Body.Close()

	resultHTML, _ := io.ReadAll(domResp.Body)
	input.DOM = string(resultHTML)

	// parse any JS/CSS from the dom
	node, err := html.Parse(bytes.NewReader(resultHTML))
	if err == nil {
		extractEmbedded(node, &input)
		extractTitle(node, &input)
	}

	for _, cookie := range result.Data.Cookies {
		input.Cookies = append(input.Cookies, cookie.Name+"="+cookie.Value)
	}
	foundHTML := false
	for _, request := range result.Data.Requests {
		input.Requests = append(input.Requests, request.Request.Request.Url)

		// TODO: how does this check behave in the case of redirects?
		if request.Request.PrimaryRequest {
			// this is the "primary" page load, so we need to extract the response headers
			for headerKey, headerValue := range request.Response.Response.Headers {
				input.Headers = append(input.Headers, http.CanonicalHeaderKey(headerKey)+": "+headerValue)
			}
			sort.Slice(input.Headers, func(i, j int) bool {
				return input.Headers[i] < input.Headers[j]
			})
		}

		if request.Response.Hash == "" {
			// this isn't a response we can fetch
			continue
		}

		switch request.Request.Type {
		case "Stylesheet", "Script", "Document":
			resourceReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://urlscan.io/responses/"+request.Response.Hash, nil)
			resp, err := client.Do(resourceReq)
			if err != nil {
				return Input{}, fmt.Errorf("failed to fetch resource %s %s: %w", request.Request.RequestId, request.Response.Hash, err)
			}
			resource, _ := io.ReadAll(resp.Body) // always read the body to completion to ensure proper connection re-use + caching
			resp.Body.Close()
			if resp.StatusCode/100 != 2 {
				// not all resources are saved by urlscan.io e.g. stylesheets are frequently missing
				continue
			}
			switch request.Request.Type {
			case "Stylesheet":
				input.CSS = append(input.CSS, string(resource))
			case "Script":
				input.JS = append(input.JS, string(resource))
			case "Document":
				if request.Request.PrimaryRequest {
					foundHTML = true
					if input.HTML != "" {
						fmt.Println("oops already have response html")
					}
					// this is the initial page load
					input.HTML = string(resource)

					// parse any JS/CSS from the html
					node, err := html.Parse(bytes.NewReader(resource))
					if err == nil {
						extractEmbedded(node, &input)
						extractTitle(node, &input)
					}
				}
			}
		}
	}

	if !foundHTML {
		return input, fmt.Errorf("failed to get response html")
	}

	return input, nil
}

func extractEmbedded(node *html.Node, input *Input) {
	if node.Type == html.ElementNode && node.Data == "style" && node.FirstChild != nil {
		input.CSS = append(input.CSS, node.FirstChild.Data)
	}
	if node.Type == html.ElementNode && node.Data == "script" && node.FirstChild != nil {
		input.JS = append(input.JS, node.FirstChild.Data)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extractEmbedded(c, input)
	}
}

func extractTitle(node *html.Node, input *Input) {
	if node.Type == html.ElementNode && node.Data == "title" && node.FirstChild != nil {
		input.Title = append(input.Title, node.FirstChild.Data)
		return
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extractTitle(c, input)
	}
}
