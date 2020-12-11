package deko

import (
	"testing"

	. "github.com/gregoryv/web"
)

// TestSpec generates additional source from the spec.
func Test_example_spec(t *testing.T) {
	project := NewDeko()

	spec := NewPage(
		Html(Body(project)),
	)
	spec.SaveAs("specification.md")
}
