package nlp

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

type TestCase []struct {
	Text   string   `yaml:"text"`
	Tokens []string `yaml:"tokens"`
}

// Exercise: Load the test cases from tokenize_cases.yml
//THIS IS MY CODE, MIKI USED os.OPEN to read the file, both
func TestYamlCases(t *testing.T) {
	testcases := TestCase{}
	data, err := ioutil.ReadFile("tokenize_cases.yml")
	if err != nil {
		log.Fatalf("error during file reading")
	}
	err = yaml.Unmarshal([]byte(data), &testcases)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for _, tc := range testcases {
		t.Run(tc.Text, func(t *testing.T) {
			log.Println(tc.Text)
			tokens := Tokenize(tc.Text)
			require.Equal(t, tc.Tokens, tokens)
		})
	}
}

func TestTokenizeTable(t *testing.T) {
	var testCases = []struct {
		text   string
		tokens []string
	}{
		{"", nil},
		{"Who's on first?", []string{"who", "s", "on", "first"}},
	}

	for _, tc := range testCases {
		t.Run(tc.text, func(t *testing.T) { // run sub test
			tokens := Tokenize(tc.text)
			require.Equal(t, tc.tokens, tokens)
		})
	}

}

func TestTokenize(t *testing.T) {
	text := "What's on second?"
	tokens := Tokenize(text)
	expected := []string{"what", "s", "on", "second"}
	require.Equal(t, expected, tokens)
	/* Code before testify
	if !reflect.DeepEqual(expected, tokens) {
		t.Fatalf("%q: expected: %#v, got %#v", text, expected, tokens)
	}
	*/
}
