package nlp_test

import (
	"fmt"

	nlp "github.com/leumas3003/npl"
)

func ExampleTokenize() {
	text := "Who's on first?"
	tokens := nlp.Tokenize(text)
	fmt.Println(tokens)
	// Output:
	// [who s on first]
}
