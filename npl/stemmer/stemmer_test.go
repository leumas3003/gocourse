package stemmer

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	inCI = os.Getenv("CI") != ""
)

func TestInCI(t *testing.T) {
	if !inCI {
		t.Skip("not in CI")
	}
	// See also testing.Short()

	// TODO
}

func TestStemmer(t *testing.T) {
	testCases := []struct {
		word string
		stem string
	}{
		{"works", "work"},
		{"worked", "work"},
		{"working", "work"},
	}

	for _, tc := range testCases {
		t.Run(tc.word, func(t *testing.T) {
			stem := Stem(tc.word)
			require.Equal(t, tc.stem, stem)
		})
	}
}
