package q00452

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
			input{nums: [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}},
			output{ans: 2},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, findMinArrowShots(data.i.nums))
	}
}
