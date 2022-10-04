package q00154

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
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
			input{nums: []int{1,3,5}},
			output{ans: 1},
		},
		{
			input{nums: []int{2,2,2,0,1}},
			output{ans: 0},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, findMin(data.i.nums))
	}
}