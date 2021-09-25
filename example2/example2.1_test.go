package example2_test

import (
	"testing"

	"github.com/edfoh/go-test-examples/example2"
	"github.com/stretchr/testify/assert"
)

func testSetup(t *testing.T) func(t *testing.T) {
	t.Log("test setup!")
	return func(t *testing.T) {
		t.Log("test teardown!")
	}
}

func testCaseSetup(t *testing.T) func(t *testing.T) {
	t.Log("test case setup!")
	return func(t *testing.T) {
		t.Log("test case teardown!")
	}
}

func TestSum(t *testing.T) {
	testTeardown := testSetup(t)
	defer testTeardown(t)

	testCases := []struct {
		name       string
		a          int
		b          int
		wantResult int
	}{
		{
			name:       "1+1=2",
			a:          1,
			b:          1,
			wantResult: 2,
		},
		{
			name:       "2+2=4",
			a:          2,
			b:          2,
			wantResult: 4,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCaseTeardown := testCaseSetup(t)
			defer testCaseTeardown(t)

			res := example2.Sum(testCase.a, testCase.b)
			assert.Equal(t, testCase.wantResult, res)
		})
	}
}
