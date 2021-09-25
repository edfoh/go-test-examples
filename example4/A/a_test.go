//package A
package A_test

import (
	"testing"

	"github.com/edfoh/go-test-examples/example4/A"
	"github.com/edfoh/go-test-examples/example4/B"
	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	x := B.AnotherFunc()
	//res := Func1(x)
	res := A.Func1(x)

	assert.Equal(t, 1+x, res)
}
