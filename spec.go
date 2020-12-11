package deko

import . "github.com/gregoryv/web"

func Project(name string, goal *Element, v ...interface{}) *Element {
	return Article(
		H1(name),
		"Goal: ", goal,
	).With(v...)

}

func Goal(v ...interface{}) *Element {
	return Span(Class("goal")).With(v...)
}

func Background(v ...interface{}) *Element {
	return Span(Class("background")).With(v...)
}

// Requirements returns all requirements in the specification
func Requirements(spec *Element) []*Element {
	return nil
}
