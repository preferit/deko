package deko

import . "github.com/gregoryv/web"

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

		Version(n, "0.1.0-unreleased", "",
			"Add initial background and goal of project",
		),
	)
}

func Version(n *Hn, version, date string, changes ...interface{}) *Element {
	ul := Ul()
	for _, c := range changes {
		ul.With(Li(c))
	}
	return Section(
		n.H2(version, date), ul,
	)
}
