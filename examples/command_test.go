package while_test

import (
	"io"
	"os"
	"strings"

	gloo "github.com/gloo-foo/framework"

	while "github.com/gloo-foo/cmd-while"
)

func ExampleWhile() {
	// echo "1\n2\n3" | while read line; do echo "Line: $line"; done
	src := gloo.ByteReaderSource([]io.Reader{strings.NewReader("1\n2\n3\n")})
	cmd := while.While(func(line []byte) ([]byte, error) {
		return append([]byte("Line: "), line...), nil
	})
	if _, err := gloo.Run(src, gloo.ByteWriteTo(os.Stdout), cmd); err != nil {
		panic(err)
	}
	// Output:
	// Line: 1
	// Line: 2
	// Line: 3
}
