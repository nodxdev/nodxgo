package nodx

import (
	"io"
)

// Node is the interface that wraps the basic Render methods used in the
// text, attribute, and element nodes.
type Node interface {
	// Render writes the node HTML to the writer.
	Render(w io.Writer) error

	// RenderString returns the node HTML as a string.
	RenderString() (string, error)

	// RenderBytes returns the node HTML as a byte slice.
	RenderBytes() ([]byte, error)
}
