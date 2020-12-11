package deko

import . "github.com/gregoryv/web"

func NewChangelog(n *Hn) *Element {
	return Article(
		n.H1("Changelog"),
		Version("unreleased", "",
			"Add initial background and goal of project",
		),
	)
}

func Version(version, date string, changes ...interface{}) *Element {
	ul := Ul()
	for _, c := range changes {
		ul.With(Li(c))
	}
	return Section(
		"[", version, "] ", date, ul,
	)
}
