package spec

import . "github.com/gregoryv/web"

func theme() *CSS {
	css := NewCSS()
	css.Import("https://fonts.googleapis.com/css?family=Inconsolata|Source+Sans+Pro")

	css.Style("html, body",
		"margin: 0 0",
		"body: 0 0",
		"font-family: 'Source Sans Pro', sans-serif",
	)
	css.Style("body",
		"padding: 1em 1.618em 1800px 1.618em",
		"max-width: 21cm",
		"line-height: 1.3em",
	)
	css.Style("h1:first-child",
		"margin-top: 0",
	)
	css.Style("nav ul",
		"list-style-type: none",
		"padding-left: 0",
		"margin-bottom: 3.236em",
	)
	css.Style("section",
		"margin-bottom: 3.236em",
	)
	css.Style("li.h3",
		"margin-left: 1.618em",
	)
	css.Style("li.h4",
		"margin-left: 3.236em",
	)
	css.Style("a:link",
		"text-decoration: none",
	)
	css.Style("a:link:hover",
		"text-decoration: underline",
	)
	css.Style("h1 a, h2 a, h3 a, h4 a",
		"color: black",
	)
	css.Style("h4.question",
		"color: darkred",
	)
	css.Style(".timestamp",
		"width: 100%",
		"text-align: right",
	)
	css.Style(".warn",
		"color: red",
	)

	css.Style("ul.issues",
		"list-style-type: square",
	)
	return css
}
