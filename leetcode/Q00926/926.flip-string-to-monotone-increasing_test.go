package q00926

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums string
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
			input{nums: "00110"},
			output{ans: 1},
		},
		{
			input{nums: "010110"},
			output{ans: 2},
		},
		{
			input{nums: "00011000"},
			output{ans: 2},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, minFlipsMonoIncr(data.i.nums))
	}
}
