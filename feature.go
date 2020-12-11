package deko

import (
	"regexp"
	"strings"

	. "github.com/gregoryv/web"
)

func Features(v ...interface{}) *Element {
	ul := Ul(Class("features"))
	for _, f := range v {
		ul.With(Li(f))
	}
	return ul
}

func Feature(name string) *Element {
	return A(Href("#"+FeatureIdFrom(name)), name)
}

func FeatureIdFrom(name string) string {
	txt := idChars.ReplaceAllString(name, "")
	return strings.ToLower(txt)
}

var idChars = regexp.MustCompile(`\W`)
