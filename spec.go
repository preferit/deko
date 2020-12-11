package deko

import (
	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/toc"
)

func NewSpecification(name string, project *Element) *Specification {
	return &Specification{
		name:    name,
		project: project,
	}
}

type Specification struct {
	name    string
	project *Element
}

func (me *Specification) SaveAs(filename string) {
	nav := Nav()
	body := Body(
		H1(me.name),
		nav,
		Article(
			me.project,
		),
	)
	toc.MakeTOC(nav, body, "h2", "h3")
	page := NewPage(
		Html(
			Head(
				Style(theme()),
			),
			body,
		),
	)
	page.SaveAs(filename)
}

func Project(goal *Element, v ...interface{}) *Element {
	return Wrap(
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

func Question(v ...interface{}) *Element {
	return Span(Class("question")).With(v...)
}
