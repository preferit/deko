package deko

import (
	"regexp"
	"strings"

	. "github.com/gregoryv/web"
)

func NewSpecification(project *Element) *Specification {
	return &Specification{
		project: project,
	}
}

type Specification struct {
	project *Element
}

func (me *Specification) SaveAs(filename string) {
	page := NewPage(
		Html(
			Head(
				Style(theme()),
			),
			Body(me.project)),
	)
	page.SaveAs(filename)
}

func Project(name string, goal *Element, v ...interface{}) *Element {
	return Article(
		H1(name),
		H2("Goal"),
		P(goal),
	).With(v...)

}

func Goal(v ...interface{}) *Element {
	return Wrap(v...)
}

func Background(v ...interface{}) *Element {
	return Section(Class("background"), H2("Background")).With(v...)
}

func Features(v ...interface{}) *Element {
	ul := Ul(Class("features"))
	for _, f := range v {
		ul.With(Li(f))
	}
	return ul
}

func Feature(name string) *Element {
	return A(Href("#"+FeatureIdFrom(name)), name)
}

func FeatureIdFrom(name string) string {
	txt := idChars.ReplaceAllString(name, "")
	return strings.ToLower(txt)
}

var idChars = regexp.MustCompile(`\W`)
