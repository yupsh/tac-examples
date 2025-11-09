package tac_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/tac"
)

func ExampleTac_basic() {
	// echo "First\nSecond\nThird" | tac
	gloo.MustRun(
		Tac(strings.NewReader("First\nSecond\nThird")),
	)
	// Output:
	// Third
	// Second
	// First
}
