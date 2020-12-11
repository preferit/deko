package deko

import . "github.com/gregoryv/web"

func NewSpecification(project *Element) *Specification {
	return &Specification{
		project: project,
	}
}

type Specification struct {
	project *Element
}

func (me *Specification) SaveAs(filename string) {
	page := NewPage(
		Html(Body(me.project)),
	)
	page.SaveAs(filename)
}

func Project(name string, goal *Element, v ...interface{}) *Element {
	return Article(
		H1(name),
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

// Requirements returns all requirements in the specification
func Requirements(spec *Element) []*Element {
	return nil
}
