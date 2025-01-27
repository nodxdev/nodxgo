package nodx

import (
	"fmt"
	"io"
)

// ensure that nodeText implements the Node interface.
var _ Node = (*nodeText)(nil)

// nodeText represents a text node.
//
// IMPORTANT: This node renders the text as is, without escaping it. It is the
// responsibility of the caller to ensure that the text is properly escaped.
type nodeText struct {
	text string
}

// newNodeText creates a new text node.
//
// IMPORTANT: This node renders the text as is, without escaping it. It is the
// responsibility of the caller to ensure that the text is properly escaped.
func newNodeText(text string) nodeText {
	return nodeText{
		text: text,
	}
}

// Render writes the text to the writer.
func (nt nodeText) Render(w io.Writer) error {
	if nt.text == "" {
		return nil
	}
	_, err := fmt.Fprintf(w, "%s", nt.text)
	return err
}

// RenderString returns the text as a string.
func (nt nodeText) RenderString() (string, error) {
	return nt.text, nil
}

// RenderBytes returns the text as a byte slice.
func (nt nodeText) RenderBytes() ([]byte, error) {
	return []byte(nt.text), nil
}
