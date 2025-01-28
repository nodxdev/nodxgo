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

func (ne nodeElement) Render(w io.Writer) error {
	if ne.name == "" {
		return nil
	}

	_, err := fmt.Fprintf(w, "<%s", ne.name)
	if err != nil {
		return err
	}

	for _, child := range ne.children {
		if child == nil {
			continue
		}
		if child.IsAttribute() {
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
		for _, child := range ne.children {
			if child == nil {
				continue
			}
			if child.IsElement() {
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

func (ne nodeElement) RenderString() (string, error) {
	buf := &strings.Builder{}
	err := ne.Render(buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (ne nodeElement) RenderBytes() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := ne.Render(buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (ne nodeElement) IsElement() bool {
	return true
}

func (ne nodeElement) IsAttribute() bool {
	return false
}

func (ne nodeElement) String() string {
	str, _ := ne.RenderString()
	return str
}
