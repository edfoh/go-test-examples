package example3_test

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/edfoh/go-test-examples/example3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	update = flag.Bool("update", false, "update the golden files of this test")
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestGenerate(t *testing.T) {
	testCases := []struct {
		filename string
		markdown *example3.Markdown
		wantErr  error
	}{
		{
			filename: "valid_markdown.golden",
			markdown: &example3.Markdown{
				Title: "Introduction",
				Body:  "Hi this is an intro.",
				References: []string{
					"https://example.com",
					"https://google.com",
				},
			},
		},
	}

	for _, testCase := range testCases {
		filepath := fmt.Sprintf("testdata/%s", testCase.filename)

		t.Run(fmt.Sprintf("testcase: %s", filepath), func(t *testing.T) {
			gen, err := example3.Generate(testCase.markdown)

			updateGoldenFileIfFlagSet(t, filepath, gen)
			wantGenerated := getGoldenFile(t, filepath)

			if testCase.wantErr != nil {
				require.Equal(t, testCase.wantErr, err)
				return
			}

			assert.Equal(t, wantGenerated, gen)
		})
	}
}

func getGoldenFile(t *testing.T, filepath string) string {
	t.Helper()
	b, err := ioutil.ReadFile(filepath)
	require.NoError(t, err)
	return string(b)
}

func updateGoldenFileIfFlagSet(t *testing.T, filepath string, generated string) {
	t.Helper()
	if *update {
		ioutil.WriteFile(filepath, []byte(generated), 0644)
	}
}
