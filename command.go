package command

import (
	"context"

	"github.com/destel/rill"
	gloo "github.com/gloo-foo/framework"
)

// While returns a Command that executes a body function for each input line,
// emitting the results.
func While(body func([]byte) ([]byte, error), _ ...any) gloo.Command[[]byte, []byte] {
	return gloo.FuncCommand[[]byte, []byte](func(_ context.Context, in gloo.Stream[[]byte]) gloo.Stream[[]byte] {
		return gloo.WrapFrom(rill.OrderedMap(in.Chan(), 1, body), in)
	})
}
