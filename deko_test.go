package deko

import (
	"testing"
)

// TestSpec generates additional source from the sre.
func Test_example_spec(t *testing.T) {
	sre := NewDeko()
	sre.SaveAs("docs/index.html")
}

func Test_all_requirements_have_ids(t *testing.T) {
	sre := NewDeko()
	err := sre.CheckRequirements()
	if err != nil {
		t.Fatal(err)
	}
}
