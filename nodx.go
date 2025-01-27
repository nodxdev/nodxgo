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

// Group creates a new group of nodes.
//
// This helper function allows for grouping multiple nodes into a single node without
// adding extra HTML elements.
//
// The resulting node renders its child nodes directly without any wrapping tag.
func Group(nodes ...Node) Node {
	return newNodeGroup(nodes...)
}

// El creates a new nodx.Node representing an HTML element.
//
// Useful for creating HTML elements not included in this library.
func El(name string, children ...Node) Node {
	return newNodeElement(false, name, children...)
}

// ElVoid creates a new nodx.Node representing an HTML void element.
//
// Useful for creating HTML elements not included in this library.
func ElVoid(name string, children ...Node) Node {
	return newNodeElement(true, name, children...)
}

// Attr creates a new nodx.Node representing an HTML attribute.
//
// Useful for creating HTML attributes not included in this library.
func Attr(name string, value string) Node {
	return newNodeAttribute(name, value)
}

// Text creates a new nodx.Node representing an HTML text node.
// The value is escaped to ensure it is safe to use in HTML
// preventing XSS attacks.
func Text(value string) Node {
	return newNodeText(EscapeHTML(value))
}

// Raw creates a new nodx.Node representing an HTML raw text node.
// The value is not escaped, so it is the responsibility of the caller to
// ensure that the value is safe to use in HTML.
//
// Useful when you need to render raw HTML such as <script> tags.
func Raw(value string) Node {
	return newNodeText(value)
}
