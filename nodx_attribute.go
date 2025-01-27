package nodx

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// ensure that nodeAttribute implements the Node interface.
var _ Node = (*nodeAttribute)(nil)

// nodeAttribute represents an HTML attribute.
type nodeAttribute struct {
	name  string
	value string
}

// newNodeAttribute creates a new HTML attribute.
func newNodeAttribute(name, value string) nodeAttribute {
	return nodeAttribute{
		name:  name,
		value: value,
	}
}

// Render writes the HTML attribute to the writer.
func (na nodeAttribute) Render(w io.Writer) error {
	if na.name == "" {
		return nil
	}
	_, err := fmt.Fprintf(w, " %s=\"%s\"", na.name, EscapeHTML(na.value))
	return err
}

// RenderString returns the HTML attribute as a string.
func (na nodeAttribute) RenderString() (string, error) {
	buf := &strings.Builder{}
	err := na.Render(buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// RenderBytes returns the HTML attribute as a byte slice.
func (na nodeAttribute) RenderBytes() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := na.Render(buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
