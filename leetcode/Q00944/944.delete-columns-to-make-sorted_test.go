package q00944

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	strs []string
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
			input{strs: []string{"zyx", "wvu", "tsr"}},
			output{ans: 3},
		},
		{
			input{strs: []string{"cba", "daf", "ghi"}},
			output{ans: 1},
		},
		{
			input{strs: []string{"a", "b"}},
			output{ans: 0},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, minDeletionSize(data.i.strs))
	}
}
