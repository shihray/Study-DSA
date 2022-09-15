package q00150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums   []string
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
			input{nums: []string{"2","1","+","3","*"}},
			output{ans: 9},
		},
		{
			input{nums: []string{"4","13","5","/","+"}},
			output{ans: 6},
		},
		{
			input{nums: []string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}},
			output{ans: 22},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, evalRPN(data.i.nums))
	}
}