package sre

import (
	"math/rand"
	"time"
)

// NewRID returns an mnemonic identifier suitable for a requirement. 6
// characters long, starting with 'r' followed by 3 letters and 2
// digits. The id is fairly easy to pronounce but is not unique. It's
// up to the caller to validate it's uniqeness in a specification when
// required.
func newRID() string {
	return "r" + newID(3, 2)
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
	return newWord(alf...)
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

// startWithVowel returns true always as the only id generator is
// NewRID which prefixes each id with the letter 'r'
func startWithVowel() bool {
	return true
}

func newWord(alfabets ...string) string {
	buf := make([]byte, len(alfabets))
	for i, alfabet := range alfabets {
		buf[i] = alfabet[rand.Intn(len(alfabet))]
	}
	return string(buf)
}

const (
	digits           = "0123456789"
	vowels           = "aeiouy"
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

func init() {
	rand.Seed(time.Now().UnixNano())
}
