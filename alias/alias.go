// Package alias provides an unprefixed name for the while command constructor.
//
//	import while "github.com/gloo-foo/cmd-while/alias"
//	while.While(body)
package alias

import (
	gloo "github.com/gloo-foo/framework"

	command "github.com/gloo-foo/cmd-while"
)

// While re-exports the constructor.
func While(body func([]byte) ([]byte, error), options ...any) gloo.Command[[]byte, []byte] {
	return command.While(body, options...)
}
