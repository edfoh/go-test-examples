package example1_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/edfoh/go-test-examples/example1"
)

func TestDeduct_SubTests(t *testing.T) {
	total := 100

	t.Run("deduction <= total", func(t *testing.T) {
		t.Run("deduction == total", func(t *testing.T) {
			bal, err := example1.Deduct(total, 100)
			require.NoError(t, err)
			assert.Equal(t, 0, bal)
		})
		t.Run("deduction < total", func(t *testing.T) {
			bal, err := example1.Deduct(total, 99)
			require.NoError(t, err)
			assert.Equal(t, 1, bal)
		})
	})

	t.Run("deduction > total", func(t *testing.T) {
		_, err := example1.Deduct(total, 101)
		assert.EqualError(t, err, "deduction must be less than or equal to total")
	})
}

func TestDeduct_Table(t *testing.T) {
	testCases := []struct {
		name        string
		total       int
		deduction   int
		wantBalance int
		wantErr     error
	}{
		{
			name:        "deduction == total",
			total:       100,
			deduction:   100,
			wantBalance: 0,
		},
		{
			name:        "deduction < total",
			total:       100,
			deduction:   99,
			wantBalance: 1,
		},
		{
			name:      "deduction > total",
			total:     100,
			deduction: 101,
			wantErr:   errors.New("deduction must be less than or equal to total"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			bal, err := example1.Deduct(testCase.total, testCase.deduction)

			if testCase.wantErr != nil {
				require.Equal(t, testCase.wantErr, err)
				return
			}

			assert.Equal(t, testCase.wantBalance, bal)
		})
	}
}

func TestDeduct_Closure(t *testing.T) {
	deductEquals := func(total, deduction, wantBalance int, wantErr error) func(t *testing.T) {
		return func(t *testing.T) {
			bal, err := example1.Deduct(total, deduction)

			if wantErr != nil {
				require.Equal(t, wantErr, err)
				return
			}

			assert.Equal(t, wantBalance, bal)
		}
	}

	total := 100
	t.Run("deduction == total", deductEquals(total, 100, 0, nil))
	t.Run("deduction < total", deductEquals(total, 99, 1, nil))
	t.Run("deduction > total", deductEquals(total, 101, 0, errors.New("deduction must be less than or equal to total")))
}
