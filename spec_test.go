package deko

import (
	"fmt"
	"testing"
)

// TestSpec generates additional source from the spec.
func TestSpec(t *testing.T) {
	s := NewSpecification()
	for _, r := range s.Requirements() {
		fmt.Println(r)
	}
}
