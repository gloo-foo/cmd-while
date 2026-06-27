package alias_test

import (
	"bytes"
	"slices"
	"testing"

	while "github.com/gloo-foo/cmd-while/alias"
	"github.com/gloo-foo/testable"
)

// The alias package re-exports the While constructor under an unprefixed name.
// A mis-wired re-export (While bound to the wrong function) compiles cleanly, so
// only behavior can prove the wiring. The test exercises the re-export and
// asserts that the supplied body runs over every input line in order.

func TestAlias_WhileRunsBodyOverEveryLine(t *testing.T) {
	cmd := while.While(func(line []byte) ([]byte, error) {
		return bytes.ToUpper(line), nil
	})
	lines, err := testable.TestLines(cmd, "hello\nworld\n")
	if err != nil {
		t.Fatal(err)
	}
	if want := []string{"HELLO", "WORLD"}; !slices.Equal(lines, want) {
		t.Fatalf("got %q, want %q", lines, want)
	}
}
