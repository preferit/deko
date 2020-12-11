package deko

import (
	"regexp"
	"strings"

	"github.com/gregoryv/web"
	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/toc"
)

func NewSpecification(name string, goal, background interface{}) *Specification {
	spec := &Specification{
		name:       name,
		background: background,
	}
	spec.goal.main = goal
	return spec
}

type Specification struct {
	name string
	goal struct {
		main      interface{}
		secondary []interface{}
	}
	background interface{}
	changelog  interface{}
}

func (me *Specification) SaveAs(filename string) {
	nav := Nav()
	openQuestions := Wrap()
	body := Body(
		H1(me.name),
		me.goal.main,
		nav,
		Article(
			H2("Goals"),
			me.goal.main,
			openQuestions,
			H2("Background"),
			me.background,
		),
		me.changelog,
	)

	// add open questions
	questions := findQuestions(body)
	if len(questions) > 0 {
		ul := Ul()
		for _, q := range questions {
			qid := genID(q.Text())
			q.With(Id(qid))
			ul.With(Li(
				A(Href("#"+qid), Class("question"), q.Text())),
			)
		}
		openQuestions.With(H2("Open questions"), ul)
	}

	// fix all non html elements
	renameElement(body, "question", "h4")

	toc.MakeTOC(nav, body, "h2", "h3", "h4")
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
		if e.Name == "question" {
			res = append(res, e)
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
	return Section(Class("background")).With(v...)
}

func Question(v string) *Element {
	return NewElement("question", Class("question"), v)
}

var idChars = regexp.MustCompile(`\W`)

func genID(v string) string {
	txt := idChars.ReplaceAllString(v, "")
	return strings.ToLower(txt)
}
