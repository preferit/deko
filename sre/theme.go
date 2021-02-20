package sre

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
	css.Style(".doctitle",
		"font-size: 2em",
	)
	css.Style("nav ul",
		"list-style-type: none",
		"padding-left: 0",
		"margin-bottom: 3.236em",
	)
	css.Style("section",
		"margin-bottom: 3.236em",
	)
	css.Style("li.h2", "margin-left: 1.618em")
	css.Style("li.h3", "margin-left: 2.618em")
	css.Style("li.h4", "margin-left: 3.618em")

	css.Style("a:link",
		"text-decoration: none",
	)
	css.Style("a:link:hover",
		"text-decoration: underline",
	)
	css.Style("h1 a, h2 a, h3 a, h4 a",
		"color: black",
	)

	// questions related
	css.Style("li.h5",
		"text-align: right",
		"border-bottom: 1px dotted #727272",
	)
	css.Style("li.h5 a, td.question a",
		"color: darkred",
		"font-style: italic",
	)
	css.Style("table.question",
		"width: 100%",
		"border-bottom: 1px dotted #727272",
	)
	css.Style("td.question",
		"text-align: right",
		"width: 100%",
	)
	css.Style("td.question h5",
		"color: darkred",
		"padding: 0 0",
		"margin: 0 0",
		"font-size: 1em",
		"font-weight: normal",
	)
	css.Style("td.answer",
		"background-color: #e2e2e2",
		"border-radius: 5px",
		"padding: 2px 1em 2px 1em",
	)
	css.Style("td.answer a",
		"color: black",
	)
	// end questions

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

	css.Style("dd",
		"margin-bottom: 1em",
	)
	return css
}
