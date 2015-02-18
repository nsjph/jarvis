package plugin

import (
	irc "github.com/thoj/go-ircevent"
	"math/rand"
	"regexp"
)

type Random struct {
	Regexp *regexp.Regexp
}

func random(n int) int {
	return rand.Intn(n)
}

func newRandom(prepend string) *Random {
	r := new(Random)
	r.Regexp = ""
	return r
}
