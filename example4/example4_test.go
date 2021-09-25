package example4_test

import (
	"testing"

	"github.com/edfoh/go-test-examples/example4"
	"github.com/stretchr/testify/assert"
)

func TestInternalFunc(t *testing.T) {
	res := example4.InternalFunc()

	assert.Equal(t, 1, res)
}
