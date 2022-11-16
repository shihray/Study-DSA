package q00779

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	n int
	k int
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
			input{n: 1, k: 1},
			output{ans: 0},
		},
		{
			input{n: 2, k: 1},
			output{ans: 0},
		},
		{
			input{n: 2, k: 2},
			output{ans: 1},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, kthGrammar(data.i.n, data.i.k))
	}
}