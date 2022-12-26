package q00739

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
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
			input{nums: []int{73, 74, 75, 71, 69, 72, 76, 73}},
			output{ans: []int{1, 1, 4, 2, 1, 1, 0, 0}},
		},
		{
			input{nums: []int{30, 40, 50, 60}},
			output{ans: []int{1, 1, 1, 0}},
		},
		{
			input{nums: []int{30, 60, 90}},
			output{ans: []int{1, 1, 0}},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, dailyTemperatures(data.i.nums))
	}
}
