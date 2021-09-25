package B

import "github.com/edfoh/go-test-examples/example4/A"

func AnotherFunc() int {
	a := A.Func2()

	return a + 1
}
