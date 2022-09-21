package q00394

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	str string
}

type output struct {
	ans string
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{str: "3[a]2[bc]"},
			output{ans: "aaabcbc"},
		},
		{
			input{str: "3[a2[c]]"},
			output{ans: "accaccacc"},
		},
		{
			input{str: "2[abc]3[cd]ef"},
			output{ans: "abcabccdcdcdef"},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, decodeString(data.i.str))
	}
}