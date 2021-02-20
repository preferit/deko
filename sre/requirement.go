package sre

import (
	. "github.com/gregoryv/web"
)

// Requirements returns an ul element with listed requirements.
// Requirements without id are classed as warn.
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

// NewRequirement returns a new requirement with no id. This is useful
// during the elicitation process.
func NewRequirement(v string) *Requirement {
	return &Requirement{Txt: v}
}

type Requirement struct {
	ID  string
	Txt string
}

const requirementTag = "requirement"
