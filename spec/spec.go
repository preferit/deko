package spec

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
	Name         string
	Goals        *Element
	CurrentState *Element
	Changelog    *Element
	References   *Element
}

func (me *Specification) SaveAs(filename string) {
	nav := Nav()
	openQuestions := Wrap()
	mainGoal := FindFirstChild(me.Goals, "maingoal")
	body := Body(
		Div(Class("timestamp"), "Last update: ", time.Now().Format("2006-01-02 15:04")),

		H1(me.Name),
		mainGoal,
		nav,
		Article(
			me.Goals,
			openQuestions,
			me.CurrentState,
		),
		me.Changelog,
		me.References,
	)

	// add open questions
	groupQuestions(openQuestions, body)

	// fix all non html elements
	renameElement(body, "question", "h4")
	renameElement(body, "requirement", "div")
	renameElement(me.Goals, "maingoal", "em")
	renameElement(me.Goals, "goal", "wrapper")
	renameElement(me.CurrentState, "issue", "div")
	refs := anchorDt(me.References)

	linkReferences(me.Goals, refs)
	linkReferences(me.CurrentState, refs)

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

func Question(v string) *Element {
	return NewElement("question", Class("question"), v)
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
		if e.Name == REQ {
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
			wb.WriteString(RID())
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
		el := NewElement(REQ, Class("requirement"), req.Txt, Attr("title", id))
		if id != "" {
			el.With(Id(id))
		} else {
			el.With(Span(Class("warn"), " (Missing ID)"))
		}
		ul.With(Li(el))
	}
	return ul
}

func NewRequirement(v string) *Requirement {
	return &Requirement{Txt: v}
}

type Requirement struct {
	ID  string
	Txt string
}

const REQ = "requirement"

var idChars = regexp.MustCompile(`\W`)

func genID(v string) string {
	txt := idChars.ReplaceAllString(v, "")
	return strings.ToLower(txt)
}
