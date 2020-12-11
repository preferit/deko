package deko

import (
	"regexp"
	"strings"
	"time"

	"github.com/gregoryv/web"
	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/toc"
)

type Specification struct {
	name       string
	goals      *Element
	background interface{}
	changelog  interface{}
}

func (me *Specification) SaveAs(filename string) {
	nav := Nav()
	openQuestions := Wrap()
	mainGoal := FindFirstChild(me.goals, "maingoal")
	body := Body(
		Div(Class("timestamp"), "Last update: ", time.Now().Format("2006-01-02 15:04")),

		H1(me.name),
		mainGoal,
		nav,
		Article(
			H2("Goals"),
			me.goals,
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
	renameElement(body, "requirement", "div")
	renameElement(me.goals, "maingoal", "em")
	renameElement(me.goals, "goal", "wrapper")

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

func FindFirstChild(root *Element, name string) (found *Element) {
	web.WalkElements(root, func(e *web.Element) {
		if e.Name == name {
			if found == nil {
				found = e
			}
		}
	})
	return
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

func MainGoal(v string) *Element {
	return NewElement("maingoal", Class("goal main-goal"), v)
}

func Goal(v string) *Element {
	return NewElement("goal", Class("goal"), v)
}

func Background(v ...interface{}) *Element {
	return Section(Class("background")).With(v...)
}

func Question(v string) *Element {
	return NewElement("question", Class("question"), v)
}

func Requirements(v ...interface{}) *Element {
	ul := Ul()
	for _, req := range v {
		ul.With(Li(req))
	}
	return ul
}

func Requirement(v string) *Element {
	return NewElement("requirement", Class("requirement"), v)
}

var idChars = regexp.MustCompile(`\W`)

func genID(v string) string {
	txt := idChars.ReplaceAllString(v, "")
	return strings.ToLower(txt)
}
