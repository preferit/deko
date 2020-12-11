package deko

import (
	"testing"
)

// TestSpec generates additional source from the spec.
func Test_example_spec(t *testing.T) {
	spec := NewDeko()
	spec.SaveAs("docs/index.html")

}
