package tac_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/tac"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleTac_fromFile_basic() {
	// tac testdata/text.txt
	yup.MustRun(
		Tac(yup.File("testdata/text.txt")),
	)
	// Output:
	// Third
	// Second
	// First
}

