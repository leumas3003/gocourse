package nlp

import (
	"regexp"
	"strings"

	"github.com/leumas3003/npl/stemmer"
)

var (
	wordRe = regexp.MustCompile(`[[:alpha:]]+`)
)

// Tokenize returns a list of tokens (lower case) found in text.
func Tokenize(text string) []string {
	words := wordRe.FindAllString(text, -1)
	// var tokens []string
	// 75 percentils of sentences have 15 or less tokens
	tokens := make([]string, 0, 15)
	for _, w := range words {
		token := strings.ToLower(w)
		token = stemmer.Stem(token)
		if token != "" {
			tokens = append(tokens, token)
		}
	}
	return tokens
}
