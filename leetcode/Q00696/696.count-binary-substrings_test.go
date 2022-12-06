package q00696

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums string
}

type output struct {
	ans int
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{nums: "00110011"},
			output{ans: 6},
		},
		{
			input{nums: "10101"},
			output{ans: 4},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, countBinarySubstrings(data.i.nums))
	}
}