package deko

import (
	"regexp"
	"strings"

	"github.com/gregoryv/web"
	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/toc"
)

func NewSpecification(name string, goal, background interface{}) *Specification {
	return &Specification{
		name:       name,
		goal:       goal,
		background: background,
	}
}

type Specification struct {
	name       string
	goal       interface{}
	background interface{}
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

	// add open questions
	findQuestions(body)

	// fix all non html elements
	renameElement(body, "question", "span")
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

func findQuestions(root *Element) []*Element {
	res := make([]*Element, 0)

	web.WalkElements(root, func(e *web.Element) {
		if e.Name != "question" {
			return
		}
	})
	return res
}

func renameElement(root *Element, from, to string) {
	web.WalkElements(root, func(e *web.Element) {
		if e.Name == from {
			e.Name = to
		}
	})
}

func Goal(v ...interface{}) *Element {
	return Em(v...)
}

func Background(v ...interface{}) *Element {
	return Section(Class("background"), H2("Background")).With(v...)
}

func Question(v string) *Element {
	return NewElement("question", Class("question"), v)
}

var idChars = regexp.MustCompile(`\W`)

func genID(v string) string {
	txt := idChars.ReplaceAllString(v, "")
	return strings.ToLower(txt)
}
