package iok

import (
	"bufio"
	"github.com/bradleyjkemp/cupaloy/v2"
	"net/http"
	"strings"
	"testing"
)

func TestInputFromHTTPResponse(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "https://html5test.com", nil)
	resp, err := http.ReadResponse(bufio.NewReader(strings.NewReader(testHTTPResponse)), req)
	if err != nil {
		t.Fatal(err)
	}

	input, err := InputFromHTTPResponse(resp)
	if err != nil {
		t.Fatal(err)
	}

	// This is parsed by http.ParseResponse into a map so is randomly ordered
	sortField(input.Headers)

	cupaloy.New(cupaloy.ShouldUpdate(func() bool {
		return true
	})).SnapshotT(t, input)
}

var testHTTPResponse = `HTTP/1.0 200 OK
Accept-Ranges: bytes
Connection: Keep-Alive
Content-Encoding: gzip
Content-Length: 6081
Content-Type: text/html; charset=UTF-8
Date: Wed, 20 Dec 2023 16:09:03 GMT
Etag: 4a30-561faa1383116-gzip
Keep-Alive: timeout=5, max=100
Last-Modified: Thu, 04 Jan 2018 22:12:38 GMT
Server: Apache
Vary: Accept-Encoding
Set-Cookie: cookieName=cookieValue; HttpOnly

<!DOCTYPE html>
<html lang="en">
	<head>
		<title>HTML5test - How well does your browser support HTML5?</title>

		<meta charset="UTF-8">

		<meta http-equiv="X-UA-Compatible" content="IE=EDGE">

		<meta name="apple-mobile-web-app-capable" content="yes">
		<meta name="apple-mobile-web-app-status-bar-style" content="black">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">

		<link rel="stylesheet" href="/css/main.css" type="text/css">

		<script src='/scripts/base.js' type='text/javascript'></script>
		<script src='/scripts/8/engine.js' type='text/javascript'></script>
		<script src='/scripts/8/data.js' type='text/javascript'></script>

		<script>
			(function(){var p=[],w=window,d=document,e=f=0;p.push('ua='+encodeURIComponent(navigator.userAgent));e|=w.ActiveXObject?1:0;e|=w.opera?2:0;e|=w.chrome?4:0;
			e|='getBoxObjectFor' in d || 'mozInnerScreenX' in w?8:0;e|=('WebKitCSSMatrix' in w||'WebKitPoint' in w||'webkitStorageInfo' in w||'webkitURL' in w)?16:0;
			e|=(e&16&&({}.toString).toString().indexOf("\n")===-1)?32:0;p.push('e='+e);f|='sandbox' in d.createElement('iframe')?1:0;f|='WebSocket' in w?2:0;
			f|=w.Worker?4:0;f|=w.applicationCache?8:0;f|=w.history && history.pushState?16:0;f|=d.documentElement.webkitRequestFullScreen?32:0;f|='FileReader' in w?64:0;
			p.push('f='+f);p.push('r='+Math.random().toString(36).substring(7));p.push('w='+screen.width);p.push('h='+screen.height);var s=d.createElement('script');
			s.src='//api.whichbrowser.net/rel/detect.js?' + p.join('&');d.getElementsByTagName('head')[0].appendChild(s);})();
		</script>

		<meta name="application-name" content="HTML5test"/>

		<link rel="apple-touch-icon" sizes="57x57" href="/images/icons/apple-touch-icon-57x57.png" />
		<link rel="apple-touch-icon" sizes="114x114" href="/images/icons/apple-touch-icon-114x114.png" />
		<link rel="apple-touch-icon" sizes="72x72" href="/images/icons/apple-touch-icon-72x72.png" />
		<link rel="apple-touch-icon" sizes="144x144" href="/images/icons/apple-touch-icon-144x144.png" />
		<link rel="apple-touch-icon" sizes="60x60" href="/images/icons/apple-touch-icon-60x60.png" />
		<link rel="apple-touch-icon" sizes="120x120" href="/images/icons/apple-touch-icon-120x120.png" />
		<link rel="apple-touch-icon" sizes="76x76" href="/images/icons/apple-touch-icon-76x76.png" />
		<link rel="apple-touch-icon" sizes="152x152" href="/images/icons/apple-touch-icon-152x152.png" />
		<link rel="icon" type="image/png" href="/images/icons/favicon-16x16.png" sizes="16x16" />
		<link rel="icon" type="image/png" href="/images/icons/favicon-32x32.png" sizes="32x32" />
		<link rel="icon" type="image/png" href="/images/icons/favicon-96x96.png" sizes="96x96" />
		<link rel="icon" type="image/png" href="/images/icons/favicon-160x160.png" sizes="160x160" />
		<link rel="icon" type="image/png" href="/images/icons/favicon-196x196.png" sizes="196x196" />
		<meta name="msapplication-TileColor" content="#0092bf" />
		<meta name="msapplication-TileImage" content="/images/icons/mstile-144x144.png" />

		<meta property="og:title" content="The HTML5 test - How well does your browser support HTML5?" />
		<meta property="og:description" content="The HTML5 test score is an indication of how well your browser supports the upcoming HTML5 standard and related specifications. How well does your browser support HTML5?" />
		<meta property="og:type" content="website" />
		<meta property="og:url" content="https://html5test.com" />
		<meta property="og:image" content="https://html5test.com/images/thumbnail.jpg" />
		<meta property="og:site_name" content="HTML5test.com" />
		<meta property="fb:admins" content="1087142132" />
		
		<link href="https://twitter.com/html5test" rel="me">
	</head>


<!--
	Copyright (c) 2010-2016 Niels Leenheer

	Permission is hereby granted, free of charge, to any person obtaining
	a copy of this software and associated documentation files (the
	"Software"), to deal in the Software without restriction, including
	without limitation the rights to use, copy, modify, merge, publish,
	distribute, sublicense, and/or sell copies of the Software, and to
	permit persons to whom the Software is furnished to do so, subject to
	the following conditions:

	The above copyright notice and this permission notice shall be
	included in all copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
	EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
	MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
	NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
	LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
	OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
	WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
-->


	<body>
		<div id="fb-root"></div>
		<div id='contentwrapper'>
			<div class='header'>
				<h1><span>HTML<strong>5</strong>test</span> <em>how well does your browser support HTML5?</em></h1>

				<div class='navigation'>
					<ul class='left'>
						<li class='selected'><a href='/index.html'>Your browser</a></li>
						<li><a href='/results/desktop.html'>Other browsers</a></li>
						<li><a href='/compare/browser/index.html'>Compare</a></li>
					</ul>
					<ul class='right'>
						<li><a href='http://blog.html5test.com/'>News</a></li>
						<li><a href='/devicelab'>Device Lab</a></li>
						<li><a href='/about.html'>About the test</a></li>
					</ul>
				</div>
			</div>

			<div class='page'>
				<noscript>
					<h2>
						To view the results of your browser you need to enable Javascript!
					</h2>
				</noscript>

				<div id='loading'><div></div></div>

				<div id='warning'></div>

				<div id='contents' class='column' style='visibility: hidden;'>
					<div id='score'></div>
					<div id='results'></div>

					<div id='headerad'>
						<!-- BuySellAds.com Zone Code -->
						<div id="bsap_1262954" class="bsarocks bsap_6c18c40d896c427f049e74e4c708ab51"></div>
						<!-- End BuySellAds.com Zone Code -->
					</div>

					<div class='paper'>
						<div>
							<h2>About HTML5test</h2>

							<div class='text'>
								<span id='html5'></span>

								<p>
									The HTML5 test score is an indication of how well your browser supports the
									HTML5 standard and related specifications. Find out which parts of HTML5 are
									supported by your browser today and compare the results with other browsers.
								</p>

								<p>
									The HTML5 test does not try to test all of the new features offered by HTML5, nor does it try to test the functionality of each feature it does detect.
									Despite these shortcomings we hope that by quantifying the level of support users and web developers will get an idea of how hard the
									browser manufacturers work on improving their browsers and the web as a development platform.
								</p>

								<p>
									The score is calculated by testing for the many new features of HTML5. Each feature is worth one or more points. Apart from the main HTML5
									specification and other specifications created the W3C HTML Working Group or WHATWG, this test also awards points for supporting related drafts and
									specifications.
								</p>
								<p>
									Please be aware that although the HTML5 specification is now an official recommendation, other specifications that are being tested are still in development
									and could change before receiving an official status.
									In the future new tests will be added for new specifications and existing tests will be updated when the specifications change.
								</p>

								<h3>Help us improve HTML5 test by donating</h3>
								<p>
									This will allow us to spend more time on HTML5test.com and acquire more devices for our testing lab.

								</p>
							</div>
						</div>

						<div class='buttons'>
							<a class='button donate' href="https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=9DNBJPQFEHYSC"><span>Donate using PayPal</span></a>
							<a class='button followme' href="https://www.twitter.com/html5test"><span>Follow me on Twitter</span></a>
							<a class='button contact' href="mailto:info@html5test.com"><span>Contact me</span></a>
							<a class='button developed' href="https://github.com/NielsLeenheer/html5test"><span>Developed at GitHub</span></a>
						</div>
					</div>

					<div id='footerad'>
						<!-- BuySellAds.com Zone Code -->
						<div id="bsap_1262955" class="bsarocks bsap_6c18c40d896c427f049e74e4c708ab51"></div>
						<!-- End BuySellAds.com Zone Code -->
					</div>


					<div class='footer'>
						<div>
							<div class='copyright'>
								<p>
									June 2016 - version 8.0. Created by Niels Leenheer.<br>
									Please note that the HTML5 test is not affiliated with the W3C or the HTML5 working group.<br>
									HTML5 Logo by <a href="http://www.w3.org/"><abbr title="World Wide Web Consortium">W3C</abbr></a>.
									Browser detection by <a href='http://whichbrowser.net'>WhichBrowser</a>.
								</p>

								<div id='cloudvps'>
									<a href="http://www.cloudvps.nl" target="_blank" ><b>CloudVPS</b><br> High Availability<br> Cloud Servers</a>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
		<div id='index'></div>

		<script>
		<!--


			function waitForWhichBrowser(cb) {
				var callback = cb;

				function wait() {
					if (typeof WhichBrowser == 'undefined')
						window.setTimeout(wait, 100)
					else
						callback();
				}

				wait();
			}


			waitForWhichBrowser(function() {

				Browsers = new WhichBrowser({
					useFeatures:		true,
					detectCamouflage:	true
				});

				Blacklist = { 'OneBrowser' : { blocked: true, identifier: Browsers.browser.toString() } };

				if (typeof Blacklist[Browsers.browser.name] != 'undefined')
				{
					if (Blacklist[Browsers.browser.name].blocked)
					{
						var contents = document.getElementById('contents');
						contents.style.visibility = 'visible';

						var warning = document.getElementById('warning');
						warning.innerHTML =
							"<div class='blocked'>" +
								"<h2>" + Blacklist[Browsers.browser.name].identifier + " has been blocked!</h2>" +
								"<p>It appears that your browser is too buggy to test or is deliberately attempting to get higher scores " +
								"by claiming it supports features that it does not. In any case consider switching to a different browser, " +
								"because this will also cause problems for websites that attempt to use these features.</p>" +
								"<button onclick='this.style.display=\"none\"; start();'>Show results anyway</button>" +
 							"</div>"

						var loading = document.getElementById('loading');
						loading.style.display = 'none';

						this.load();
						return;
					}
				}

				start();
			});


			function start() {
				var params = {};

				if (location.search) {
					var parts = location.search.substring(1).split('&');

					for (var i = 0; i < parts.length; i++) {
						var nv = parts[i].split('=');
						if (!nv[0]) continue;
						params[nv[0]] = decodeURIComponent(nv[1]) || true;
					}
				}

				var identifier = typeof params.identifier != 'undefined' ? params.identifier : '';
				var source = typeof params.source != 'undefined' ? params.source : '';
				var task = typeof params.task != 'undefined' ? params.task : '';

				new Test(function(r) {
					var m = new Metadata(tests);

					var c = new Calculate(r, m.data);

					/* Submit results */
					try {
						var payload = '{' +
									  '"release": "' + r.release + '",' +
									  '"source": "' + source + '",' +
									  '"identifier": "' + escapeSlashes(identifier) + '",' +
									  '"task": "' + task + '",' +
									  '"uniqueid": "' + r.uniqueid + '",' +
									  '"score": ' + c.score + ',' +
									  '"maximum": ' + c.maximum + ',' +
									  '"camouflage": "' + (Browsers.camouflage ? '1' : '0') + '",' +
									  '"features": "' + (Browsers.features.join(',')) + '",' +
									  '"browserName": "' + (Browsers.browser.name ? Browsers.browser.name : '') + '",' +
									  '"browserChannel": "' + (Browsers.browser.channel ? Browsers.browser.channel : '') + '",' +
									  '"browserVersion": "' + (Browsers.browser.version ? Browsers.browser.version.toString() : '') + '",' +
									  '"browserVersionType": "' + (Browsers.browser.version ? Browsers.browser.version.type : '') + '",' +
									  '"browserVersionMajor": "' + (Browsers.browser.version ? Browsers.browser.version.major : '') + '",' +
									  '"browserVersionMinor": "' + (Browsers.browser.version ? Browsers.browser.version.minor : '') + '",' +
									  '"browserVersionOriginal": "' + (Browsers.browser.version ? Browsers.browser.version.original : '') + '",' +
									  '"browserMode": "' + (Browsers.browser.mode ? Browsers.browser.mode : '') + '",' +
									  '"engineName": "' + (Browsers.engine.name ? Browsers.engine.name : '') + '",' +
									  '"engineVersion": "' + (Browsers.engine.version ? Browsers.engine.version.toString() : '') + '",' +
									  '"osName": "' + (Browsers.os.name ? Browsers.os.name : '') + '",' +
									  '"osFamily": "' + (Browsers.os.family ? Browsers.os.family : '') + '",' +
									  '"osVersion": "' + (Browsers.os.version ? Browsers.os.version.toString() : '') + '",' +
									  '"deviceManufacturer": "' + (Browsers.device.manufacturer ? Browsers.device.manufacturer : '') + '",' +
									  '"deviceModel": "' + (Browsers.device.model ? Browsers.device.model : '') + '",' +
									  '"deviceSeries": "' + (Browsers.device.series ? Browsers.device.series : '') + '",' +
									  '"deviceType": "' + (Browsers.device.type ? Browsers.device.type : '') + '",' +
									  '"deviceIdentified": "' + (Browsers.device.identified ? '1' : '0' ) + '",' +
									  '"deviceWidth": "' + (screen.width) + '",' +
									  '"deviceHeight": "' + (screen.height) + '",' +
									  '"useragent": "' + navigator.userAgent + '",' +
									  '"humanReadable": "' + Browsers.toString() + '",' +
									  '"points": "' + c.points + '",' +
									  '"results": "' + r.results + '"' +
									  '}';

						submit('submit', payload);
					} catch(e) {
						alert('Could not submit results: ' + e.message);
					}


					/* Update total score */
					var container = document.getElementById('score');

					container.innerHTML = tim(
						"<div class='pointsPanel'>" +
						"<h2><span>Your browser scores</span> <strong>{{score}}</strong> <span>out of {{maximum}} points</span></h2>" +
						"</div>",
					c);


					/* Show box for confirming useragent detection */
					new Confirm(container, {
						id:			r.uniqueid,
						onConfirm:	function() { submit('confirm', '{"uniqueid": "' + r.uniqueid + '"}'); },
						onReport:	function() { submit('report', '{"uniqueid": "' + r.uniqueid + '"}'); },
						onFeedback:	function(value) { submit('feedback', JSON.stringify({ uniqueid: r.uniqueid, value: value })); }
					});

					/* Show action buttons */
					var wrapper = document.createElement('div');
					wrapper.className = 'wrapper';
					container.appendChild(wrapper);

					var buttons = document.createElement('div');
					buttons.className = 'buttons';
					wrapper.appendChild(buttons);

					var button = document.createElement('span');
					button.className = 'button save';
					button.innerHTML = '<span>Save results</span>';
					buttons.appendChild(button);

					new Save(button, {
						id:			r.uniqueid,
						onSave:		function() { submit('save', '{"uniqueid": "' + r.uniqueid + '"}'); }
					});

					var button = document.createElement('a');
					button.className = 'button compare';
					button.href = '/compare/browser/mybrowser.html';
					button.innerHTML = '<span>Compare to...</span>';
					buttons.appendChild(button);

					var button = document.createElement('span');
					button.className = 'button share';
					button.innerHTML = '<span>Share</span>';
					buttons.appendChild(button);

					new Share(button, {
						score: 		c.score,
						browser:	Browsers.toString()
					});

					var button = document.createElement('a');
					button.className = 'button donate';
					button.href = 'https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=9DNBJPQFEHYSC';
					button.innerHTML = '<span>Donate</span>';
					buttons.appendChild(button);

					/* Show detailed report of scores */
					var container = document.getElementById('results');
					var div = document.createElement('div');
					div.className = 'resultsTable detailsTable';
					container.appendChild(div);

					var table = new ResultsTable({
						parent:			div,
						tests:			m.data,
						header:			false,
						links:			true,
						explainations: 	true,
						grading:		true,
						bonus:			true,
						distribute:		true,
						columns:		1
					});

					table.updateColumn(0, { points: c.points, maximum: c.maximum, score: c.score, results: r.results });

					new Index({
						tests:			tests,
						index:			document.getElementById('index'),
						wrapper:		document.getElementById('contentwrapper')
					});


					function submit(method, payload) {
						var httpRequest;
						if (window.XMLHttpRequest) {
							httpRequest = new XMLHttpRequest();
						} else if (window.ActiveXObject) {
							httpRequest = new ActiveXObject("Microsoft.XMLHTTP");
						}

					   	httpRequest.open('POST','/api/' + method, true);
					   	httpRequest.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
	   					httpRequest.send('payload=' + encodeURIComponent(payload));
					}

					window.setTimeout(function() {
						var contents = document.getElementById('contents');
						contents.style.visibility = 'visible';

						var loading = document.getElementById('loading');
						loading.style.display = 'none';
					}, 100);
				},

				function(e) {
					if (typeof console != 'undefined') console.log(e);

					alert('Test has failed: ' + e.message);

					var contents = document.getElementById('contents');
					contents.style.visibility = 'visible';

					var loading = document.getElementById('loading');
					loading.style.display = 'none';
				});

				this.load();
			}

			function load() {
				(function(){
				  	var bsa = document.createElement('script');
					 bsa.type = 'text/javascript';
					 bsa.async = true;
					 bsa.src = '//s3.buysellads.com/ac/bsa.js';
					 (document.getElementsByTagName('head')[0]||document.getElementsByTagName('body')[0]).appendChild(bsa);
				})();


				var _gaq = _gaq || [];
				_gaq.push(['_setAccount', 'UA-68192-4']);
				_gaq.push(['_trackPageview']);

				(function() {
					var ga = document.createElement('script');
					ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
					ga.setAttribute('async', 'true');
					document.documentElement.firstChild.appendChild(ga);
				})();
			}

		//-->
		</script>
	</body>
</html>
`
