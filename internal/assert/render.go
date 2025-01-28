package assert

import (
	"io"
	"strings"
	"testing"
)

// Node with basic Render methods.
//
// Redeclared to avoid import cycle.
type Node interface {
	Render(w io.Writer) error
	RenderString() (string, error)
	RenderBytes() ([]byte, error)
	String() string
}

// Render renders a NodX node with all the Render methods and
// asserts that the output is equal to the expected string.
func Render(t *testing.T, expected string, node Node) {
	t.Helper()

	buf := &strings.Builder{}
	err := node.Render(buf)
	NoError(t, err)
	Equal(t, expected, buf.String())

	gotStr, err := node.RenderString()
	NoError(t, err)
	Equal(t, expected, gotStr)

	gotBytes, err := node.RenderBytes()
	NoError(t, err)
	Equal(t, expected, string(gotBytes))

	gotStr = node.String()
	Equal(t, expected, gotStr)
}
