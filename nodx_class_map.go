package nodx

import (
	"bytes"
	"io"
	"sort"
	"strings"
)

// ensure that ClassMap implements the Node interface.
var _ Node = (*ClassMap)(nil)

// ClassMap represents a map of class names with conditional rendering.
// The keys are the class names, and the values are boolean expressions
// (conditions) that determine whether each class should be included
// in the final output.
//
// Example:
//
//	isOdd := func(n int) bool { return n % 2 != 0 }
//	isEven := func(n int) bool { return n % 2 == 0 }
//
//	renderOdd := isOdd(3)   // true
//	renderEven := isEven(3) // false
//
//	cm := ClassMap{
//		"odd-class":  renderOdd,      // Included
//		"even-class": renderEven,     // Excluded
//		"always-on":  true,           // Always included
//	}
//
// This will render the class attribute as: class="odd-class always-on"
type ClassMap map[string]bool

// Render writes the "class" attribute to the provided writer.
// It includes only the class names with a `true` value, sorted alphabetically.
func (cm ClassMap) Render(w io.Writer) error {
	classes := []string{}

	for class, condition := range cm {
		if condition {
			classes = append(classes, class)
		}
	}

	sort.Strings(classes)
	classesStr := strings.Join(classes, " ")

	return Attr("class", classesStr).Render(w)
}

// RenderString returns the "class" attribute as a string.
// It includes only the class names with a `true` value, sorted alphabetically.
func (cm ClassMap) RenderString() (string, error) {
	buf := &strings.Builder{}
	err := cm.Render(buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// RenderBytes returns the "class" attribute as a byte slice.
// It includes only the class names with a `true` value, sorted alphabetically.
func (cm ClassMap) RenderBytes() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := cm.Render(buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
