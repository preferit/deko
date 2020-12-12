package deko

import (
	"testing"
)

// TestSpec generates additional source from the spec.
func Test_example_spec(t *testing.T) {
	spec := NewDeko()
	spec.SaveAs("docs/index.html")
}

func Test_all_requirements_have_ids(t *testing.T) {
	spec := NewDeko()
	err := spec.CheckRequirements()
	if err != nil {
		t.Fatal(err)
	}
}
