package tac_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/tac"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleTac_fromFile_basic() {
	// tac testdata/text.txt
	gloo.MustRun(
		Tac(gloo.File("testdata/text.txt")),
	)
	// Output:
	// Third
	// Second
	// First
}
