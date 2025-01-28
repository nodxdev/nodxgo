package nodx

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// ensure that nodeElement implements the Node interface.
var _ Node = (*nodeElement)(nil)

// nodeElement represents an HTML element.
type nodeElement struct {
	isVoid   bool
	name     string
	children []Node
}

// newNodeElement creates a new HTML element.
func newNodeElement(isVoid bool, name string, children ...Node) nodeElement {
	return nodeElement{
		isVoid:   isVoid,
		name:     name,
		children: children,
	}
}

// Render writes the HTML element and all its children to the writer.
func (ne nodeElement) Render(w io.Writer) error {
	if ne.name == "" {
		return nil
	}

	_, err := fmt.Fprintf(w, "<%s", ne.name)
	if err != nil {
		return err
	}

	// Render only nodeAttributes
	for _, child := range ne.children {
		if _, ok := child.(nodeAttribute); ok {
			err = child.Render(w)
			if err != nil {
				return err
			}
		}
		if _, ok := child.(ClassMap); ok {
			err = child.Render(w)
			if err != nil {
				return err
			}
		}
	}

	_, err = fmt.Fprintf(w, ">")
	if err != nil {
		return err
	}

	if !ne.isVoid {
		// Render only nodeElements and nodeTexts
		for _, child := range ne.children {
			if _, ok := child.(nodeElement); ok {
				err = child.Render(w)
				if err != nil {
					return err
				}
				continue
			}
			if _, ok := child.(nodeText); ok {
				err = child.Render(w)
				if err != nil {
					return err
				}
			}
		}

		_, err = fmt.Fprintf(w, "</%s>", ne.name)
		if err != nil {
			return err
		}
	}

	return nil
}

// RenderString returns the HTML element and all its children as a string.
func (ne nodeElement) RenderString() (string, error) {
	buf := &strings.Builder{}
	err := ne.Render(buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// RenderBytes returns the HTML element and all its children as a byte slice.
func (ne nodeElement) RenderBytes() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := ne.Render(buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
