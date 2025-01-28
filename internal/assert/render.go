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

// RenderNoSpaces is the same as Render, but it removes all the white
// spaces and similar characters from the expected and the output to
// make the test more readable without the need to remove spaces
// manually when writing the test.
func RenderNoSpaces(t *testing.T, expected string, node Node) {
	t.Helper()

	removeSpaces := func(s string) string {
		for _, char := range []string{"\n", "\t", "\r", " "} {
			s = strings.ReplaceAll(s, char, "")
		}
		return s
	}

	expected = removeSpaces(expected)

	buf := &strings.Builder{}
	err := node.Render(buf)
	NoError(t, err)
	Equal(t, expected, removeSpaces(buf.String()))

	gotStr, err := node.RenderString()
	NoError(t, err)
	Equal(t, expected, removeSpaces(gotStr))

	gotBytes, err := node.RenderBytes()
	NoError(t, err)
	Equal(t, expected, removeSpaces(string(gotBytes)))

	gotStr = node.String()
	Equal(t, expected, removeSpaces(gotStr))
}
