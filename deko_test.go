package deko

import (
	"testing"
)

func Test_generate_deko_specification(t *testing.T) {
	sre := NewDeko()
	sre.Page().SaveAs("docs/index.html")
}

func Test_all_requirements_have_ids(t *testing.T) {
	sre := NewDeko()
	if err := sre.CheckRequirements(); err != nil {
		t.Fatal(err)
	}
}
