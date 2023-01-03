package q01962

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
	k    int
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
			input{nums: []int{4, 3, 6, 7}, k: 3},
			output{ans: 12},
		},
		{
			input{nums: []int{5, 4, 9}, k: 2},
			output{ans: 12},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, minStoneSum(data.i.nums, data.i.k))
	}
}
