package command_test

import (
	"bytes"
	"testing"

	command "github.com/gloo-foo/cmd-while"
	"github.com/gloo-foo/testable"
)

func TestWhile_UpperCase(t *testing.T) {
	cmd := command.While(func(line []byte) ([]byte, error) {
		return bytes.ToUpper(line), nil
	})
	lines, err := testable.TestLines(cmd, "hello\nworld\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 2 || lines[0] != "HELLO" || lines[1] != "WORLD" {
		t.Errorf("got %q, want [\"HELLO\", \"WORLD\"]", lines)
	}
}

func TestWhile_EmptyInput(t *testing.T) {
	cmd := command.While(func(line []byte) ([]byte, error) {
		return line, nil
	})
	lines, err := testable.TestLines(cmd, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(lines) != 0 {
		t.Errorf("got %d lines, want 0", len(lines))
	}
}
