package q02256

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
			input{nums: []int{2,5,3,9,5,3}},
			output{ans: 3},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, minimumAverageDifference(data.i.nums))
	}
}