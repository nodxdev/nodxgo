// Package nodx is a modern and developer-friendly Go template engine for
// generating safe, clean, and maintainable HTML. Designed for maximum
// productivity and easy maintenance, it combines simplicity, type safety
// and robust formatting, making it the perfect fit for your Go-based web
// projects.
package nodx

import (
	"fmt"
	"io"
)

// Node is the interface that wraps the basic Render methods used in the
// different NodX node types.
type Node interface {
	// Render writes the node HTML to the writer.
	Render(w io.Writer) error

	// RenderString returns the node HTML as a string.
	RenderString() (string, error)

	// RenderBytes returns the node HTML as a byte slice.
	RenderBytes() ([]byte, error)
}

// Group combines multiple nodes into a single node without wrapping them in any HTML tag.
// It renders the child nodes directly.
func Group(nodes ...Node) Node {
	return newNodeGroup(nodes...)
}

// El creates a new Node representing an HTML element with the given name.
func El(name string, children ...Node) Node {
	return newNodeElement(false, name, children...)
}

// ElVoid creates a new Node representing an HTML void element with the given name.
// Void elements are self-closing, such as <img> or <input>.
func ElVoid(name string, children ...Node) Node {
	return newNodeElement(true, name, children...)
}

// Attr creates a new Node representing an HTML attribute with a name and value.
// The value is HTML-escaped to prevent XSS attacks.
func Attr(name string, value string) Node {
	return newNodeAttribute(name, value)
}

// Text creates a new Node representing an escaped HTML text node.
// The value is HTML-escaped to prevent XSS attacks.
func Text(value string) Node {
	return newNodeText(EscapeHTML(value))
}

// Textf creates a new Node representing an escaped HTML text node.
// The value is formatted with fmt.Sprintf and then HTML-escaped to prevent XSS attacks.
func Textf(format string, a ...any) Node {
	return Text(fmt.Sprintf(format, a...))
}

// Raw creates a new Node representing a raw HTML text node.
// The value is not escaped, so the caller must ensure the content is safe.
// Useful for rendering raw HTML, like <script> or <style> tags.
func Raw(value string) Node {
	return newNodeText(value)
}

// Rawf creates a new Node representing a raw HTML text node.
// The value is formatted with fmt.Sprintf and not escaped, so ensure its safety.
// Useful for rendering raw HTML, like <script> or <style> tags.
func Rawf(format string, a ...any) Node {
	return Raw(fmt.Sprintf(format, a...))
}

// If renders a Node based on the provided boolean condition.
// If the condition is true, the node is rendered; otherwise, it renders nothing.
func If(condition bool, node Node) Node {
	if condition {
		return node
	}
	return Group()
}

// IfFunc executes and renders the result of a function based on the provided boolean condition.
// If the condition is true, the function is executed and rendered; otherwise, nothing is executed nor rendered.
func IfFunc(condition bool, function func() Node) Node {
	if condition {
		return function()
	}
	return Group()
}

// Map transforms a slice of any type into a group of nodes by applying a function to each element.
// Returns a single Node that contains all the resulting nodes.
func Map[T any](slice []T, function func(T) Node) Node {
	nodes := make([]Node, len(slice))
	for i, item := range slice {
		nodes[i] = function(item)
	}
	return Group(nodes...)
}

// DocType Defines the document type.
//
// This is a special element and should be used inside a nodx.Group.
//
// Output: <!DOCTYPE html>
func DocType() Node {
	return Raw("<!DOCTYPE html>")
}
