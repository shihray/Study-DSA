package q00931

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums [][]int
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
			input{nums: [][]int{
				{2,1,3},
				{6,5,4},
				{7,8,9},
			}},
			output{ans: 13},
		},
		{
			input{nums: [][]int{
				{-19,57},
				{-40,-5},
			}},
			output{ans: -59},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, minFallingPathSum(data.i.nums))
	}
}