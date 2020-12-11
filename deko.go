package deko

import . "github.com/gregoryv/web"

func NewDeko() *Element {
	p := Project("deko",
		Goal("Simplify time keeping between consultants and customers"),

		Background(),
	)

	return p
}
