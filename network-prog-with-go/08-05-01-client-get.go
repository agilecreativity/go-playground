package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "http://host:port/page")
		os.Exit(1)
	}

	url, err := url.Parse(os.Args[1])
	checkError(err)

	client := &http.Client{}
	request, err := http.NewRequest("GET", url.String(), nil)

	// only accept UTF-8
	request.Header.Add("Accept-Charset", "UTF-8;q=1, ISO-8859-1;q=0")
	checkError(err)

	response, err := client.Do(request)

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	chSet := getCharset(response)
	fmt.Printf("Got charset :%s\n", chSet)

	if chSet != "UTF-8" {
		fmt.Println("Cannot handle", chSet)
		os.Exit(4)
	}

	var buf [512]byte
	reader := response.Body
	fmt.Println("Got body")

	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Print(string(buf[0:n]))
	}
	os.Exit(0)
}

func getCharset(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")
	if contentType == "" {
		// guess
		return "UTF-8"
	}
	idx := strings.Index(contentType, "charset:")
	if idx == -1 {
		// guess
		return "UTF-8"
	}
	return strings.Trim(contentType[idx:], " ")
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

/* Sample output:
./08-05-01-client-get http://golang.org:80/doc/

Got charset :UTF-8
Got body
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">

  <title>Documentation - The Go Programming Language</title>

<link type="text/css" rel="stylesheet" href="/lib/godoc/style.css">

<link rel="search" type="application/opensearchdescription+xml" title="godoc" href="/opensearch.xml" />

<link rel="stylesheet" href="/lib/godoc/jquery.treeview.css">
<script type="text/javascript">window.initFuncs = [];</script>
<script type="text/javascript">
var _gaq = _gaq || [];
_gaq.push(["_setAccount", "UA-11222381-2"]);
_gaq.push(["b._setAccount", "UA-49880327-6"]);
window.trackPageview = function() {
  _gaq.push(["_trackPageview", location.pathname+location.hash]);
  _gaq.push(["b._trackPageview", location.pathname+location.hash]);
};
window.trackPageview();
window.trackEvent = function(category, action, opt_label, opt_value, opt_noninteraction) {
  _gaq.push(["_trackEvent", category, action, opt_label, opt_value, opt_noninteraction]);
  _gaq.push(["b._trackEvent", category, action, opt_label, opt_value, opt_noninteraction]);
};
</script>
</head>
<body>

<div id='lowframe' style="position: fixed; bottom: 0; left: 0; height: 0; width: 100%; border-top: thin solid grey; background-color: white; overflow: auto;">
...
</div><!-- #lowframe -->

<div id="topbar" class="wide"><div class="container">

<form method="GET" action="/search">
<div id="menu">
<a href="/doc/">Documents</a>
<a href="/pkg/">Packages</a>
<a href="/project/">The Project</a>
<a href="/help/">Help</a>
<a href="/blog/">Blog</a>

<a id="playgroundButton" href="http://play.golang.org/" title="Show Go Playground">Play</a>

<input type="text" id="search" name="q" class="inactive" value="Search" placeholder="Search">
</div>
<div id="heading"><a href="/">The Go Programming Language</a></div>
</form>

</div></div>


<div id="playground" class="play">
	<div class="input"><textarea class="code">package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}</textarea></div>
	<div class="output"></div>
	<div class="buttons">
		<a class="run" title="Run this code [shift-enter]">Run</a>
		<a class="fmt" title="Format this code">Format</a>
		<a class="share" title="Share this code">Share</a>
	</div>
</div>


<div id="page" class="wide">
<div class="container">


  <h1>Documentation</h1>




<div id="nav"></div>




<p>
The Go programming language is an open source project to make programmers more
productive.
</p>

<p>
Go is expressive, concise, clean, and efficient. Its concurrency
mechanisms make it easy to write programs that get the most out of multicore
and networked machines, while its novel type system enables flexible and
modular program construction. Go compiles quickly to machine code yet has the
convenience of garbage collection and the power of run-time reflection. It's a
fast, statically typed, compiled language that feels like a dynamically typed,
interpreted language.
</p>

<div id="manual-nav"></div>

<h2>Installing Go</h2>

<h3><a href="/doc/install">Getting Started</a></h3>
<p>
Instructions for downloading and installing the Go compilers, tools, and
libraries.
</p>


<h2 id="learning">Learning Go</h2>

<img class="gopher" src="/doc/gopher/doc.png"/>

<h3 id="go_tour"><a href="//tour.golang.org/">A Tour of Go</a></h3>
<p>
An interactive introduction to Go in three sections.
The first section covers basic syntax and data structures; the second discusses
methods and interfaces; and the third introduces Go's concurrency primitives.
Each section concludes with a few exercises so you can practice what you've
learned. You can <a href="//tour.golang.org/">take the tour online</a> or
<a href="//code.google.com/p/go-tour/">install it locally</a>.
</p>

<h3 id="code"><a href="code.html">How to write Go code</a></h3>
<p>
Also available as a
<a href="//www.youtube.com/watch?v=XCsL89YtqCs">screencast</a>, this doc
explains how to use the <a href="/cmd/go/">go command</a> to fetch, build, and
install packages, commands, and run tests.
</p>

<h3 id="effective_go"><a href="effective_go.html">Effective Go</a></h3>
<p>
A document that gives tips for writing clear, idiomatic Go code.
A must read for any new Go programmer. It augments the tour and
the language specification, both of which should be read first.
</p>

<h3 id="faq"><a href="/doc/faq">Frequently Asked Questions (FAQ)</a></h3>
<p>
Answers to common questions about Go.
</p>

<h3 id="wiki"><a href="/wiki">The Go Wiki</a></h3>
<p>A wiki maintained by the Go community.</p>

<h4 id="learn_more">More</h4>
<p>
See the <a href="/wiki/Learn">Learn</a> page at the <a href="/wiki">Wiki</a>
for more Go learning resources.
</p>


<h2 id="references">References</h2>

<h3 id="pkg"><a href="/pkg/">Package Documentation</a></h3>
<p>
The documentation for the Go standard library.
</p>

<h3 id="cmd"><a href="/doc/cmd">Command Documentation</a></h3>
<p>
The documentation for the Go tools.
</p>

<h3 id="spec"><a href="/ref/spec">Language Specification</a></h3>
<p>
The official Go Language specification.
</p>

<h3 id="go_mem"><a href="/ref/mem">The Go Memory Model</a></h3>
<p>
A document that specifies the conditions under which reads of a variable in
one goroutine can be guaranteed to observe values produced by writes to the
same variable in a different goroutine.
</p>

<h3 id="release"><a href="/doc/devel/release.html">Release History</a></h3>
<p>A summary of the changes between Go releases.</p>


<h2 id="articles">Articles</h2>

<h3 id="blog"><a href="//blog.golang.org/">The Go Blog</a></h3>
<p>The official blog of the Go project, featuring news and in-depth articles by
the Go team and guests.</p>

<h4>Codewalks</h4>
<p>
Guided tours of Go programs.
</p>
<ul>
<li><a href="/doc/codewalk/functions">First-Class Functions in Go</a></li>
<li><a href="/doc/codewalk/markov">Generating arbitrary text: a Markov chain algorithm</a></li>
<li><a href="/doc/codewalk/sharemem">Share Memory by Communicating</a></li>
<li><a href="/doc/articles/wiki/">Writing Web Applications</a> - building a simple web application.</li>
</ul>

<h4>Language</h4>
<ul>
<li><a href="/blog/json-rpc-tale-of-interfaces">JSON-RPC: a tale of interfaces</a></li>
<li><a href="/blog/gos-declaration-syntax">Go's Declaration Syntax</a></li>
<li><a href="/blog/defer-panic-and-recover">Defer, Panic, and Recover</a></li>
<li><a href="/blog/go-concurrency-patterns-timing-out-and">Go Concurrency Patterns: Timing out, moving on</a></li>
<li><a href="/blog/go-slices-usage-and-internals">Go Slices: usage and internals</a></li>
<li><a href="/blog/gif-decoder-exercise-in-go-interfaces">A GIF decoder: an exercise in Go interfaces</a></li>
<li><a href="/blog/error-handling-and-go">Error Handling and Go</a></li>
<li><a href="/blog/organizing-go-code">Organizing Go code</a></li>
</ul>

<h4>Packages</h4>
<ul>
<li><a href="/blog/json-and-go">JSON and Go</a> - using the <a href="/pkg/encoding/json/">json</a> package.</li>
<li><a href="/blog/gobs-of-data">Gobs of data</a> - the design and use of the <a href="/pkg/encoding/gob/">gob</a> package.</li>
<li><a href="/blog/laws-of-reflection">The Laws of Reflection</a> - the fundamentals of the <a href="/pkg/reflect/">reflect</a> package.</li>
<li><a href="/blog/go-image-package">The Go image package</a> - the fundamentals of the <a href="/pkg/image/">image</a> package.</li>
<li><a href="/blog/go-imagedraw-package">The Go image/draw package</a> - the fundamentals of the <a href="/pkg/image/draw/">image/draw</a> package.</li>
</ul>

<h4>Tools</h4>
<ul>
<li><a href="/doc/articles/go_command.html">About the Go command</a> - why we wrote it, what it is, what it's not, and how to use it.</li>
<li><a href="/blog/c-go-cgo">C? Go? Cgo!</a> - linking against C code with <a href="/cmd/cgo/">cgo</a>.</li>
<li><a href="/doc/gdb">Debugging Go Code with GDB</a></li>
<li><a href="/blog/godoc-documenting-go-code">Godoc: documenting Go code</a> - writing good documentation for <a href="/cmd/godoc/">godoc</a>.</li>
<li><a href="/blog/profiling-go-programs">Profiling Go Programs</a></li>
<li><a href="/doc/articles/race_detector.html">Data Race Detector</a> - a manual for the data race detector.</li>
<li><a href="/blog/race-detector">Introducing the Go Race Detector</a> - an introduction to the race detector.</li>
<li><a href="/doc/asm">A Quick Guide to Go's Assembler</a> - an introduction to the assembler used by Go.</li>
</ul>

<h4 id="articles_more">More</h4>
<p>
See the <a href="/wiki/Articles">Articles page</a> at the
<a href="/wiki">Wiki</a> for more Go articles.
</p>


<h2 id="talks">Talks</h2>

<img class="gopher" src="/doc/gopher/talks.png"/>

<h3 id="video_tour_of_go"><a href="http://research.swtch.com/gotour">A Video Tour of Go</a></h3>
<p>
Three things that make Go fast, fun, and productive:
interfaces, reflection, and concurrency. Builds a toy web crawler to
demonstrate these.
</p>

<h3 id="go_code_that_grows"><a href="//vimeo.com/53221560">Code that grows with grace</a></h3>
<p>
One of Go's key design goals is code adaptability; that it should be easy to take a simple design and build upon it in a clean and natural way. In this talk Andrew Gerrand describes a simple "chat roulette" server that matches pairs of incoming TCP connections, and then use Go's concurrency mechanisms, interfaces, and standard library to extend it with a web interface and other features. While the function of the program changes dramatically, Go's flexibility preserves the original design as it grows.
</p>

<h3 id="go_concurrency_patterns"><a href="//www.youtube.com/watch?v=f6kdp27TYZs">Go Concurrency Patterns</a></h3>
<p>
Concurrency is the key to designing high performance network services. Go's concurrency primitives (goroutines and channels) provide a simple and efficient means of expressing concurrent execution. In this talk we see how tricky concurrency problems can be solved gracefully with simple Go code.
</p>

<h3 id="advanced_go_concurrency_patterns"><a href="//www.youtube.com/watch?v=QDDwwePbDtw">Advanced Go Concurrency Patterns</a></h3>
<p>
This talk expands on the <i>Go Concurrency Patterns</i> talk to dive deeper into Go's concurrency primitives.
</p>

<h4 id="talks_more">More</h4>
<p>
See the <a href="/talks">Go Talks site</a> and <a href="/wiki/GoTalks">wiki page</a> for more Go talks.
</p>


<h2 id="nonenglish">Non-English Documentation</h2>

<p>
See the <a href="/wiki/NonEnglish">NonEnglish</a> page
at the <a href="/wiki">Wiki</a> for localized
documentation.
</p>


<div id="footer">
Build version go1.4.2.<br>
Except as <a href="https://developers.google.com/site-policies#restrictions">noted</a>,
the content of this page is licensed under the
Creative Commons Attribution 3.0 License,
and code is licensed under a <a href="/LICENSE">BSD license</a>.<br>
<a href="/doc/tos.html">Terms of Service</a> |
<a href="http://www.google.com/intl/en/policies/privacy/">Privacy Policy</a>
</div>

</div><!-- .container -->
</div><!-- #page -->

<!-- TODO(adonovan): load these from <head> using "defer" attribute? -->
<script type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js"></script>
<script type="text/javascript" src="/lib/godoc/jquery.treeview.js"></script>
<script type="text/javascript" src="/lib/godoc/jquery.treeview.edit.js"></script>


<script type="text/javascript" src="/lib/godoc/playground.js"></script>

<script type="text/javascript" src="/lib/godoc/godocs.js"></script>

<script type="text/javascript">
(function() {
  var ga = document.createElement("script"); ga.
*/
