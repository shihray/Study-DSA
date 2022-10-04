package q00287

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums   []int
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
			input{nums: []int{3,1,3,4,2}},
			output{ans: 3},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, findDuplicate(data.i.nums))
	}
}