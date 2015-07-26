package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ", os.Args[0], "http://proxy-host:port http://host:port/page")
	}

	proxyString := os.Args[1]
	proxyUrl, err := url.Parse(proxyString)
	checkError(err)

	rawURL := os.Args[2]
	url, err := url.Parse(rawURL)
	checkError(err)

	transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	client := &http.Client{Transport: transport}

	request, err := http.NewRequest("GET", url.String(), nil)
	dump, _ := httputil.DumpRequest(request, false)
	fmt.Println(string(dump))

	response, err := client.Do(request)
	checkError(err)
	fmt.Println("Read OK")

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}
	fmt.Println("Response OK")

	var buf [512]byte
	reader := response.Body

	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(1)
		}
		fmt.Print(string(buf[0:n]))
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

/* Sample output:
$./08-06-01-simple-proxy http://XYZ.com:8080/ http://www.google.com

GET / HTTP/1.1
Host: www.google.com

Read OK
Response OK
<!DOCTYPE html>
<!--[if lt IE 7 ]><html class="ie6" lang="en"><![endif]-->
<!--[if IE 7 ]><html class="ie7" lang="en"><![endif]-->
<!--[if IE 8 ]><html class="ie8" lang="en"><![endif]-->
<!--[if IE]><html id="ie" lang="en"><![endif]-->
<!--[if (gte IE 9)|!(IE)]><!--><html lang="en"><!--<![endif]-->
<head>
<meta charset="UTF-8" />
<meta name="viewport" content="width=device-width,initial-scale=1, user-scalable=no">
<link rel="shortcut icon" type="image/x-icon" href="/wp-content/themes/xyz/favicon.ico?v=1">
<title>XYZ - Domain Names for Generation XYZ&reg;</title>
<link rel="stylesheet" type="text/css"  href="style.css?v=1"/>
<link rel="alternate" type="application/rss+xml" title=".xyz Registry - Generic Top Level Domain (gTLD) &raquo; Feed" href="http://xyz.com/feed" />
<link rel="alternate" type="application/rss+xml" title=".xyz Registry - Generic Top Level Domain (gTLD) &raquo; Comments Feed" href="http://xyz.com/comments/feed" />
<link rel="alternate" type="application/rss+xml" title=".xyz Registry - Generic Top Level Domain (gTLD) &raquo; Homepage Comments Feed" href="http://xyz.com/homepage/feed" />
<link rel="EditURI" type="application/rsd+xml" title="RSD" href="http://xyz.com/xmlrpc.php?rsd" />
<link rel="wlwmanifest" type="application/wlwmanifest+xml" href="http://xyz.com/wp-includes/wlwmanifest.xml" />
<link rel='next' title='About .xyz' href='http://xyz.com/registry' />
<meta name="generator" content="WordPress 3.6.1" />

<meta property="og:title" content="XYZ - Domain Names for Generation XYZ&reg;" />
<meta property="og:type" content="website" />
<meta property="og:url" content="http://xyz.com" />
<meta property="og:image" content="http://xyz.com/images/og-logo.png" />
<meta property="og:image:width" content="350" />
<meta property="og:image:height" content="350" />

<meta name="description" content="XYZ is the domain name registry operator for .xyz, .college, and many more domain extensions." />

<meta name="keywords" content="xyz, xyz.com, .xyz, .college, gtlds, new domain names, domain extensions, ntlds, registry operator" />

<link rel="canonical" href="http://xyz.com/" />

</head>
<body class="home page page-id-4 page-template page-template-template-home-php">

	<h1 id="title"><span id="reg">XYZ<sup>&reg;</sup></span></h1>
	<span class="tagline">Thank you for visiting the past. Cick below to experience the future.</span>

	<table>
		<tr>
			<th>Community</th>
			<th>Registry</th>
		</tr>
		<tr>
			<td><a target="_blank" href="https://gen.xyz">.xyz</a></td>
			<td><a target="_blank" href="http://nic.xyz">nic.xyz</a></td>
		</tr>
		<tr>
			<td><a target="_blank" href="https://go.college">.college</a></td>
			<td><a target="_blank" href="http://nic.college">nic.college</a></td>
		</tr>
		<tr>
			<td></td>
			<td><a target="_blank" href="http://nic.rent">nic.rent</a></td>
		</tr>
		<tr>
			<td></td>
			<td>nic.theatre (coming soon)</td>
		</tr>
		<tr>
			<td></td>
			<td>nic.security (coming soon)</td>
		</tr>
		<tr>
			<td></td>
			<td>nic.protection (coming soon)</td>
		</tr>
		<tr>
			<td></td>
			<th class="th2">Cars Registry Limited</th>
		</tr>
		<tr>
			<td></td>
			<td>nic.car (coming soon)</td>
		</tr>
		<tr>
			<td></td>
			<td><a href="http://nic.cars" target="_blank">nic.cars</a></td>
		</tr>
		<tr>
			<td></td>
			<td><a href="http://nic.auto" target="_blank">nic.auto</a></td>
		</tr>
	</table>

	<p class="careers"><a target="_blank" href="https://gen.xyz/careers">Careers</a></p>

	<p class="copywrite">&copy; 2015 <a href="/">XYZ.COM</a> LLC</p>
	<p><a href="/privacypolicy.html">Privacy Policy</a> & <a href="/termsandconditions.html">Terms of Use</a></p>

<script type="text/javascript" src="/wp-content/themes/xyz/js/theme.js"></script>

<!-- Optional code for enabling touch -->
<script src="/wp-content/themes/xyz/js/jquery.touchSwipe.min.js"></script>

<!-- This is Liquid Slider code. The full version (not .min) is also included in the js directory -->
<script src="/wp-content/themes/xyz/js/jquery.liquid-slider.js"></script>
<script type="text/javascript">
$(function() {

	$('#slider-id').liquidSlider({
	autoSlide:false,
	hideSideArrows:false,
	hideArrowsWhenMobile:false,
	mobileNavigation:false,
	hoverArrowDuration:250,
	hideSlideArrowsDuration:250,
	responsive: true,
	swipe: true
	});
	$('#quotes').liquidSlider({
	autoSlide:true,
	autoSlideInterval:1800,
	hideSideArrows:true,
	hideArrowsWhenMobile:true,
	mobileNavigation:false,
	responsive: true,
	swipe: true
	});

})

  var _gaq = _gaq || [];
  _gaq.push(['_setAccount', 'UA-33306381-1']);
  _gaq.push(['_setDomainName', 'XYZ.com']);
  _gaq.push(['_trackPageview']);

  (function() {
    var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
  })();

</script>

</body>
</html>
*/
