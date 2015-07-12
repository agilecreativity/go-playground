package main

import (
	"html"
	"log"
)

func main() {
	raw := []string{
		"hello",
		"<i>Hello</i>",
		"alert('hello');",
		"foo & bar",
		`"how are you?" He asked.`,
	}

	log.Println("html.EscapeString")
	for _, s := range raw {
		log.Printf("\t%s -> %s", s, html.EscapeString(s))
	}

	log.Println("html.UnescapeString(html.EscapeString)")
	for _, s := range raw {
		flipped := html.UnescapeString(html.EscapeString(s))
		log.Printf("\t%s -> %s", s, flipped)
	}

	escaped := []string{
		"&#225",
		"&raquo;",
		"&middot;",
		"&lt;i&gt;htllo&lt;/i&gt;",
	}

	log.Println("html.UnescapeString")
	for _, s := range escaped {
		log.Printf("\t%s -> %s", s, html.UnescapeString(s))
	}
}

/** Output:
2015/07/12 18:06:21 html.EscapeString
2015/07/12 18:06:21 	hello -> hello
2015/07/12 18:06:21 	<i>Hello</i> -> &lt;i&gt;Hello&lt;/i&gt;
2015/07/12 18:06:21 	alert('hello'); -> alert(&#39;hello&#39;);
2015/07/12 18:06:21 	foo & bar -> foo &amp; bar
2015/07/12 18:06:21 	"how are you?" He asked. -> &#34;how are you?&#34; He asked.
2015/07/12 18:06:21 html.UnescapeString(html.EscapeString)
2015/07/12 18:06:21 	hello -> hello
2015/07/12 18:06:21 	<i>Hello</i> -> <i>Hello</i>
2015/07/12 18:06:21 	alert('hello'); -> alert('hello');
2015/07/12 18:06:21 	foo & bar -> foo & bar
2015/07/12 18:06:21 	"how are you?" He asked. -> "how are you?" He asked.
2015/07/12 18:06:21 html.UnescapeString
2015/07/12 18:06:21 	&#225 -> á
2015/07/12 18:06:21 	&raquo; -> »
2015/07/12 18:06:21 	&middot; -> ·
2015/07/12 18:06:21 	&lt;i&gt;htllo&lt;/i&gt; -> <i>htllo</i>
*/
