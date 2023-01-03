package q01834

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums [][]int
}

type output struct {
	ans []int
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{nums: [][]int{{19, 13}, {16, 9}, {21, 10}, {32, 25}, {37, 4}, {49, 24}, {2, 15}, {38, 41}, {37, 34}, {33, 6}, {45, 4}, {18, 18}, {46, 39}, {12, 24}}},
			output{ans: []int{6, 1, 2, 9, 4, 10, 0, 11, 5, 13, 3, 8, 12, 7}},
		},
		{
			input{nums: [][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}}},
			output{ans: []int{0, 2, 3, 1}},
		},
		{
			input{nums: [][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}}},
			output{ans: []int{4, 3, 2, 0, 1}},
		},
		{
			input{nums: [][]int{{7, 1}, {6, 3}, {1, 3}}},
			output{ans: []int{2,1,0}},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, getOrder(data.i.nums))
	}
}
