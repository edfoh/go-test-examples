package example4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalFunc(t *testing.T) {
	res := InternalFunc()

	assert.Equal(t, 1, res)
}
