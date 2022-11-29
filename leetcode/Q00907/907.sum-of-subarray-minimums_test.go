package q00907

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
			input{nums: []int{3,1,2,4}},
			output{ans: 17},
		},
		{
			input{nums: []int{11,81,94,43,3}},
			output{ans: 444},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, sumSubarrayMins(data.i.nums))
	}
}