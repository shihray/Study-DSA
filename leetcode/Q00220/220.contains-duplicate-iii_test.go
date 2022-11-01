package q00220

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums      []int
	indexDiff int
	valueDiff int
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
			input{
				nums:      []int{1, 2, 3, 1},
				indexDiff: 3,
				valueDiff: 0,
			},
			output{ans: true},
		},
		{
			input{
				nums:      []int{1, 5, 9, 1, 5, 9},
				indexDiff: 2,
				valueDiff: 3,
			},
			output{ans: false},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, containsNearbyAlmostDuplicate(data.i.nums, data.i.indexDiff, data.i.valueDiff))
	}
}
