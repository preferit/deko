package deko

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gregoryv/web"
	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/toc"
)

type Specification struct {
	name         string
	goals        *Element
	currentState *Element
	changelog    *Element
	references   *Element
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
			me.goals,
			openQuestions,
			me.currentState,
		),
		me.changelog,
		me.references,
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
	refs := anchorDt(me.references)

	linkReferences(me.goals, refs)
	linkReferences(me.currentState, refs)

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

func linkReferences(dst *Element, refs map[string]string) {
	web.WalkElements(dst, func(e *web.Element) {
		for i, c := range e.Children {
			switch c := c.(type) {
			case string:
				lc := strings.ToLower(c)

			replace:
				for txt, refId := range refs {
					j := strings.Index(lc, txt)
					if j > -1 {
						k := j + len(txt)
						e.Children[i] = fmt.Sprintf(`%s<a href="#%s">%s</a>%s`,
							c[:j], refId, c[j:k], c[k:],
						)
						break replace
					}
				}
			}

		}
	})
}

// anchorDt creates ids for all dt elements and returns a map of map[txt]id
func anchorDt(root *Element) map[string]string {
	refs := make(map[string]string)

	web.WalkElements(root, func(e *web.Element) {
		if e.Name == "dt" {
			txt := e.Text()
			id := genID(txt)
			refs[strings.ToLower(txt)] = id
			e.With(Id(id))
		}
	})
	return refs
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
