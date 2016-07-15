package main

import (
	"crypto/rand"
	"log"
)

type randomDataGenerator struct {
	size                 int
	repeatFreq           int
	repeatText           string
	previousContentsFreq int
	maxTextLen           int
}

func (r randomDataGenerator) Run() []string {
	if r.size < 1 {
		log.Fatalf("must specify number of random data items needed")
	}
	d := make([]string, r.size)
	rs := randomStringGenerator{r.maxTextLen}
	for i := range d {
		if (i+1)%r.repeatFreq == 0 {
			d[i] = r.repeatText
		} else if (i+1)%r.previousContentsFreq == 0 {
			d[i] = r.concatPrev(d)
		} else {
			d[i] = rs.randString()
		}
		if len(d[i]) > r.maxTextLen {
			d[i] = d[i][:r.maxTextLen]
		}
	}
	return d
}

func (r randomDataGenerator) concatPrev(data []string) string {
	var s string
	for i := 0; i < len(data); i++ {
		s += data[i]
	}
	return s
}

type randomStringGenerator struct {
	length int
}

func (r randomStringGenerator) randString() string {
	return string(r.randBytes())[:r.length]
}

func (r randomStringGenerator) randBytes() []byte {
	b := make([]byte, r.length)
	rand.Read(b)
	return b
}
