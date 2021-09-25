package example1_test

import (
	"testing"

	"github.com/edfoh/go-test-examples/example1"
	"github.com/stretchr/testify/assert"
)

type mockBalanceGetter struct {
	GetBalanceFunc func(accountID int) int
}

func (m *mockBalanceGetter) GetBalance(accountID int) int {
	return m.GetBalanceFunc(accountID)
}

func TestAdd(t *testing.T) {
	wantAccountID := 99

	testCases := []struct {
		name           string
		accountID      int
		amount         int
		getBalanceFunc func(accountID int) int
		wantBalance    int
	}{
		{
			name:      "amount and balance > 0",
			accountID: wantAccountID,
			amount:    1,
			getBalanceFunc: func(accountID int) int {
				assert.Equal(t, wantAccountID, accountID)
				return 1
			},
			wantBalance: 2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockBalanceGetter := &mockBalanceGetter{
				GetBalanceFunc: testCase.getBalanceFunc,
			}

			bal := example1.Add(mockBalanceGetter, wantAccountID, testCase.amount)

			assert.Equal(t, testCase.wantBalance, bal)
		})
	}

}
