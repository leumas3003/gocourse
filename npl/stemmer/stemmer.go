package stemmer

import (
	"strings"
)

var (
	suffixes = []string{"s", "ing", "ed"}
)

// Stem returns the stemmed version of word
// works -> work, working -> work, worked -> work
func Stem(word string) string {
	for _, suffix := range suffixes {
		if strings.HasSuffix(word, suffix) {
			return word[:len(word)-len(suffix)]
		}
	}
	return word
}
