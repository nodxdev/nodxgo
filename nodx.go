package nodx

import (
	"io"
)

// Node is the interface that wraps the basic Render methods used in the
// text, attribute, and element nodes.
type Node interface {
	Render(w io.Writer) error
	RenderString() (string, error)
	RenderBytes() ([]byte, error)
}
