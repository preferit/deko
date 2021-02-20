package deko

import . "github.com/gregoryv/web"

const LastUpdate = "2021-02-12 07:55" //time.Now().Format("2006-01-02 15:04"))

func NewChangelog() *Element {
	return Article(
		H1("Changelog"),

		P(`All notable changes to this project are documented
		here. Further details are found at `,

			A(
				Href("https://github.com/preferit/deko/commits/main"),
				"deko/commits.",
			),
		),

		newSection("0.1.0-unreleased", "",
			"Answered: What features does current log process provide?",
			"Elicit requirements of the current state",
			"Add initial background and goal of project",
		),
	)
}

func newSection(version, date string, changes ...interface{}) *Element {
	ul := Ul()
	for _, c := range changes {
		ul.With(Li(c))
	}
	return Section(
		H2(version, date), ul,
	)
}
