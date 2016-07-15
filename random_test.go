package main

import (
	"bytes"
	"testing"
)

func TestrandomDataGenerator(t *testing.T) {
	var prevContents string
	tests := []randomDataGenerator{
		randomDataGenerator{
			size:                 100,
			repeatFreq:           5,
			repeatText:           "Nothing to see here",
			previousContentsFreq: 11,
			// must be greater than repeatText length
			maxTextLen: 200,
		},
	}

	for _, r := range tests {
		d := r.Run()

		if len(d) != r.size {
			t.Errorf("size\nExpected: %q\nGot: %q", r.size, len(d))

		}

		for i, item := range d {
			if len(item) > r.maxTextLen {
				t.Fatalf("maxTextLen\nExpected: %d\nGot: %d", r.maxTextLen, len(item))

			}
			if (i+1)%r.repeatFreq == 0 {
				if item != r.repeatText {
					t.Errorf("repeatText\nExpected: %q\nGot: %q", r.repeatText, item)
				}
			} else if (i+1)%r.previousContentsFreq == 0 {
				prevContents = ""
				for j := 0; j <= i; j++ {
					prevContents += d[j]
				}
				if len(prevContents) > r.maxTextLen {
					prevContents = prevContents[:r.maxTextLen]
				}
				if item != prevContents {
					t.Fatalf("Previous Contents\nExpected: %q\nGot: %q", prevContents, item)
				}
			} else {
				if item == r.repeatText {
					t.Errorf("Not repeatText\nNot expected: %q\nGot: %q", r.repeatText, item)
				} else if item == "" {
					t.Errorf("RandomData is Blank\nGot: %q", item)
				}
			}
		}
	}
}

func TestRandomString(t *testing.T) {
	r := randomStringGenerator{20}
	randBytes := r.randBytes()
	if bytes.Equal(make([]byte, r.length), randBytes) {
		t.Errorf("randBytes()\nExpected: [% x]\nGot: [% x]", randBytes, make([]byte, r.length))
	}

	randStr := r.randString()
	if len(randStr) != r.length {
		t.Errorf("randStr() length\nExpected: %d\nGot: %d", r.length, len(randStr))
	}
}
