package deko

import (
	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/toc"
)

func NewSpecification(name string, goal, background *Element) *Specification {
	return &Specification{
		name:       name,
		goal:       goal,
		background: background,
	}
}

type Specification struct {
	name       string
	goal       *Element
	background *Element
}

func (me *Specification) SaveAs(filename string) {
	nav := Nav()
	body := Body(
		H1(me.name),
		nav,
		Article(
			H2("Goal"),
			me.goal,
			me.background,
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

func Goal(v ...interface{}) *Element {
	return Wrap(v...)
}

func Background(v ...interface{}) *Element {
	return Section(Class("background"), H2("Background")).With(v...)
}

func Question(v ...interface{}) *Element {
	return Span(Class("question")).With(v...)
}
