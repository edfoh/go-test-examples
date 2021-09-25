package example2_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/edfoh/go-test-examples/example2"
	"github.com/stretchr/testify/assert"
)

func beforeTest() {
	fmt.Println("before test!")
}

func afterTest() {
	fmt.Println("after test!")
}

func TestMain(m *testing.M) {
	beforeTest()
	ret := m.Run()
	afterTest()
	os.Exit(ret)
}

func TestSum_1(t *testing.T) {
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
			res := example2.Sum(testCase.a, testCase.b)
			assert.Equal(t, testCase.wantResult, res)
		})
	}
}
