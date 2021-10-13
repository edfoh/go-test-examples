// +build all

package example6_test

import (
	"testing"
	"time"

	"github.com/edfoh/go-test-examples/example5"
	"github.com/stretchr/testify/assert"
)

var (
	fakeNow = time.Date(2021, 11, 11, 0, 0, 0, 0, time.UTC)
)

func TestNowPlusOneHour(t *testing.T) {
	stubNowCleanup := testStubNow(t)
	defer stubNowCleanup()

	res := example5.NowPlusOneHour()

	assert.Equal(t, fakeNow.Add(1*time.Hour), res)
}

func testStubNow(t *testing.T) func() {
	t.Helper()

	example5.Now = func() time.Time {
		return fakeNow
	}

	return func() {
		example5.Now = func() time.Time {
			return time.Now()
		}
	}
}
