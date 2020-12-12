package deko

import (
	"math/rand"
	"time"
)

func RID() string {
	return "r" + newID(3, 2)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func newID(prefixLen, digitLen int) string {
	alfabets := prefixAlfabets(prefixLen)
	alf := make([]string, 0)
	for _, a := range alfabets {
		alf = append(alf, a)
	}
	for i := 0; i < digitLen; i++ {
		alf = append(alf, digits)
	}
	return NewWord(alf...)
}

func startWithVowel() bool {
	return true // rand.Intn(2) == 1
}

func prefixAlfabets(size int) []string {
	if size >= len(order) {
		size = len(order) - 1
	}
	alfabets := order[:size]
	if startWithVowel() {
		alfabets = order[1 : size+1]
	}
	return alfabets
}

func NewWord(alfabets ...string) string {
	buf := make([]byte, len(alfabets))
	for i, alfabet := range alfabets {
		buf[i] = alfabet[rand.Intn(len(alfabet))]
	}
	return string(buf)
}

const (
	digits           = "0123456789"
	vowels           = "aeioy"
	firstConsonants  = "bcdfghjklmnpqrstvz"
	secondConsonants = "bcdfghjklmnpqrstvxz"
)

var (
	// dictates max len of prefix
	order = []string{
		firstConsonants,
		vowels,
		secondConsonants,
		vowels,
		secondConsonants,
		vowels,
		secondConsonants,
	}
)
