package q00542


import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums [][]int
}

type output struct {
	ans [][]int
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
				{0,0,0},
				{0,1,0},
				{0,0,0},
			}},
			output{ans: [][]int{
				{0,0,0},
				{0,1,0},
				{0,0,0},
			}},
		},
		{
			input{nums: [][]int{
				{0,0,0},
				{0,1,0},
				{1,1,1},
			}},
			output{ans: [][]int{
				{0,0,0},
				{0,1,0},
				{1,2,1},
			}},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, updateMatrix(data.i.nums))
	}
}