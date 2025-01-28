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

func newNodeAttribute(name, value string) nodeAttribute {
	return nodeAttribute{
		name:  name,
		value: value,
	}
}

func (na nodeAttribute) Render(w io.Writer) error {
	if na.name == "" {
		return nil
	}

	_, err := fmt.Fprintf(w, "%s=\"%s\"", na.name, EscapeHTML(na.value))
	if err != nil {
		return fmt.Errorf("failed to render %s attribute: %w", na.name, err)
	}

	return nil
}

func (na nodeAttribute) RenderString() (string, error) {
	buf := &strings.Builder{}
	err := na.Render(buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (na nodeAttribute) RenderBytes() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := na.Render(buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (na nodeAttribute) IsElement() bool {
	return false
}

func (na nodeAttribute) IsAttribute() bool {
	return true
}

func (na nodeAttribute) String() string {
	str, _ := na.RenderString()
	return str
}
