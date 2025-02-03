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
	name     string
	hasValue bool
	value    string
}

func newNodeAttribute(name string, value ...string) nodeAttribute {
	val := ""
	hasValue := false
	if len(value) > 0 {
		val = value[0]
		hasValue = true
	}

	return nodeAttribute{
		name:     name,
		value:    val,
		hasValue: hasValue,
	}
}

func (na nodeAttribute) Render(w io.Writer) error {
	if na.name == "" {
		return nil
	}

	if !na.hasValue {
		_, err := fmt.Fprint(w, na.name)
		if err != nil {
			return fmt.Errorf("failed to render %s attribute: %w", na.name, err)
		}
	}

	if na.hasValue {
		_, err := fmt.Fprintf(w, "%s=\"%s\"", na.name, EscapeHTML(na.value))
		if err != nil {
			return fmt.Errorf("failed to render %s attribute: %w", na.name, err)
		}
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
