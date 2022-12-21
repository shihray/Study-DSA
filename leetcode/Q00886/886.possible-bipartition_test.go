package q00886

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	num int
	dislikes [][]int
}

type output struct {
	ans bool
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{
					num: 4,
					dislikes: [][]int{
					{1,2},{1,3},{2,4},
				},
			},
			output{ans: true},
		},
		{
			input{
					num: 3,
					dislikes: [][]int{
					{1,2},{1,3},{2,3},
				},
			},
			output{ans: false},
		},
		{
			input{
					num: 5,
					dislikes: [][]int{
					{1,2},{2,3},{3,4},{4,5},{1,5},
				},
			},
			output{ans: false},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, possibleBipartition(data.i.num, data.i.dislikes))
	}
}