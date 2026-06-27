// Package alias provides an unprefixed name for the while command constructor.
//
//	import while "github.com/gloo-foo/cmd-while/alias"
//	while.While(body)
package alias

import command "github.com/gloo-foo/cmd-while"

// While re-exports the constructor.
var While = command.While
