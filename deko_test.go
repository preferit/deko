package deko

import (
	"testing"
)

// TestSpec generates additional source from the sore.
func Test_example_spec(t *testing.T) {
	sore := NewDeko()
	sore.SaveAs("docs/index.html")
}

func Test_all_requirements_have_ids(t *testing.T) {
	sore := NewDeko()
	err := sore.CheckRequirements()
	if err != nil {
		t.Fatal(err)
	}
}
