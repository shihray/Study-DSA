package q00055

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
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
			input{nums: []int{2, 3, 1, 1, 4}},
			output{ans: true},
		},
		{
			input{nums: []int{3, 2, 1, 0, 4}},
			output{ans: false},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, canJump(data.i.nums))
	}
}
