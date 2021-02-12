package deko

import . "github.com/gregoryv/web"

const LastUpdate = "2020-12-12 10:32" //time.Now().Format("2006-01-02 15:04"))

func NewChangelog(n *Hn) *Element {
	return Article(
		n.H1("Changelog"),

		P(`All notable changes to this project are documented
		here. Further details are found at `,

			A(
				Href("https://github.com/preferit/deko/commits/main"),
				"deko/commits.",
			),
		),

		newSection(n, "0.1.0-unreleased", "",
			"Elicit requirements of the current state",
			"Add initial background and goal of project",
		),
	)
}

func newSection(n *Hn, version, date string, changes ...interface{}) *Element {
	ul := Ul()
	for _, c := range changes {
		ul.With(Li(c))
	}
	return Section(
		n.H2(version, date), ul,
	)
}
