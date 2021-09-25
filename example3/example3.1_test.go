package example3_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/edfoh/go-test-examples/example3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSONToMap(t *testing.T) {
	testCases := []struct {
		filename   string
		wantResult map[string]interface{}
		wantErrMsg string
	}{
		{
			filename: "valid.json",
			wantResult: map[string]interface{}{
				"id": "123",
			},
		},
		{
			filename:   "invalid.json",
			wantErrMsg: "invalid character '}' looking for beginning of object key string",
		},
	}

	for _, testCase := range testCases {
		filepath := fmt.Sprintf("testdata/%s", testCase.filename)

		t.Run(fmt.Sprintf("testcase: %s", filepath), func(t *testing.T) {
			b, err := ioutil.ReadFile(filepath)
			require.NoError(t, err)

			res, err := example3.JSONToMap(b)

			if testCase.wantErrMsg != "" {
				require.EqualError(t, err, testCase.wantErrMsg)
				return
			}

			assert.Equal(t, testCase.wantResult, res)
		})
	}
}
