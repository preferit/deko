package deko

import (
	"testing"
)

// TestSpec generates additional source from the spec.
func Test_example_spec(t *testing.T) {
	project := NewDeko()

	spec := NewSpecification(project)
	spec.SaveAs("docs/specification.md")
	spec.SaveAs("docs/specification.html")
}
