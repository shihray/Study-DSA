package q00841

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums [][]int
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
			input{nums: [][]int{
				{1},
				{2},
				{3},
				{},
			}},
			output{ans: true},
		},
		{
			input{nums: [][]int{
				{1,3},
				{3,0,1},
				{2},
				{0},
			}},
			output{ans: false},
		},
		{
			input{nums: [][]int{
				{2},
				{},
				{1},
			}},
			output{ans: true},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, canVisitAllRooms(data.i.nums))
	}
}