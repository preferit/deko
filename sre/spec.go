// Package sre provides ways to describe software requirements
// specifications.
package sre

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gregoryv/web"
	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/toc"
)

type Specification struct {
	Name         string
	LastUpdate   string
	Goals        *Element
	CurrentState *Element
	Changelog    *Element
	References   *Element
}

func (me *Specification) Page() *Page {
	nav := Nav()
	mainGoal := findFirstChild(me.Goals, "maingoal")
	body := Body(
		Div(Class("timestamp"), "Last update: ", me.LastUpdate),

		Div(Class("doctitle"), me.Name),
		P(mainGoal),
		nav,
		Article(
			me.Goals,
			me.CurrentState,
		),
		me.References,
		me.Changelog,
	)

	// fix all non html elements
	renameElement(body, "question", "h5")
	renameElement(body, "requirement", "div")
	renameElement(me.Goals, "maingoal", "em")
	renameElement(me.Goals, "goal", "wrapper")
	renameElement(me.CurrentState, "issue", "div")

	// include h5 so questions are visible in the top navigation
	toc.MakeTOC(nav, body, "h1", "h2", "h5")

	refs := anchorDt(me.References)
	LinkAll(me.Goals, refs)
	LinkAll(me.CurrentState, refs)

	// make sure titles not part of the navigation are easy referenced
	toc.GenerateIDs(body, "h3", "h4")
	toc.GenerateAnchors(body, "h3", "h4")

	return NewPage(
		Html(
			Head(
				Style(theme()),
			),
			body,
		),
	)
}

func groupQuestions(dst, from *Element) {
	questions := findQuestions(from)
	if len(questions) > 0 {
		ul := Ul()
		for _, q := range questions {
			qid := genID(q.Text())
			q.With(Id(qid))
			ul.With(Li(
				A(Href("#"+qid), Class("question"), q.Text())),
			)
		}
		dst.With(H2("Open questions"), ul)
	}
}

// anchorDt creates ids for all dt elements and returns a map of map[txt]id
func anchorDt(root *Element) map[string]string {
	refs := make(map[string]string)

	web.WalkElements(root, func(e *web.Element) {
		if e.Name == "dt" {
			txt := e.Text()
			id := genID(txt)
			refs[strings.ToLower(txt)] = "#" + id
			e.With(Id(id))
		}
	})
	return refs
}

func findFirstChild(root *Element, name string) (found *Element) {
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

func Question(v string) *Element {
	return Table(Class("question"),
		Tr(
			Td(
				Class("question"),
				NewElement("question", v),
			),
			Td(" "),
			Td(
				Class("answer"),
				A(
					Href(
						fmt.Sprintf("mailto:gregory@preferit.se?subject=%s", v),
					),
					"Answer",
				),
			),
		),
	)
}

func Features(v ...interface{}) *Element {
	ul := Ul(Class("features"))
	for _, el := range v {
		ul.With(Li(el))
	}
	return ul
}

func Feature(v string) *Element {
	return NewElement("Feature", Class("Feature"), v)
}

func Issues(v ...interface{}) *Element {
	ul := Ul(Class("issues"))
	for _, el := range v {
		ul.With(Li(el))
	}
	return ul
}

func Issue(v string) *Element {
	return NewElement("issue", Class("issue"), v)
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

// CheckRequirements
func (me *Specification) CheckRequirements() error {
	missingId := make([]string, 0)
	var found int
	web.WalkElements(me.CurrentState, func(e *web.Element) {
		if e.Name == requirementTag {
			found++
			if !e.HasAttr("id") {
				txt := e.Text()
				txt = strings.ReplaceAll(txt, "\t", "")
				txt = strings.ReplaceAll(txt, "\n", " ")
				missingId = append(missingId, txt)
			}
		}
	})
	if len(missingId) > 0 {
		var wb strings.Builder
		for _, r := range missingId {
			wb.WriteString(newRID()) // Propose a new rid
			wb.WriteString(" - ")
			wb.WriteString(r)
			wb.WriteString("\n")
		}
		return fmt.Errorf("Requirements missing ids\n%v", wb.String())
	}
	if found == 0 {
		return fmt.Errorf("No requirements specified")
	}
	return nil
}

func Requirements(v ...*Requirement) *Element {
	ul := Ul()
	for _, req := range v {
		id := req.ID
		el := NewElement(requirementTag, Class("requirement"), req.Txt, Attr("title", id))
		if id != "" {
			el.With(Id(id))
		} else {
			el.With(Span(Class("warn"), " (Missing ID)"))
		}
		ul.With(Li(el))
	}
	return ul
}

type Requirement struct {
	ID  string
	Txt string
}

const requirementTag = "requirement"

var idChars = regexp.MustCompile(`\W`)

func genID(v string) string {
	txt := idChars.ReplaceAllString(v, "")
	return strings.ToLower(txt)
}
