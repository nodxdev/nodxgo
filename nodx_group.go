package nodx

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// ensure that group implements the Node interface.
var _ Node = (*nodeGroup)(nil)

// nodeGroup represents a group of nodes.
type nodeGroup []Node

// newNodeGroup creates a new group of nodes.
func newNodeGroup(nodes ...Node) nodeGroup {
	return nodeGroup(nodes)
}

// Render writes the group of nodes to the writer.
func (g nodeGroup) Render(w io.Writer) error {
	for _, node := range g {
		if node == nil {
			continue
		}
		if err := node.Render(w); err != nil {
			return fmt.Errorf("failed to render group node: %w", err)
		}
	}
	return nil
}

// RenderString returns the group of nodes as a string.
func (g nodeGroup) RenderString() (string, error) {
	buf := &strings.Builder{}
	err := g.Render(buf)
	return buf.String(), err
}

// RenderBytes returns the group of nodes as a byte slice.
func (g nodeGroup) RenderBytes() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := g.Render(buf)
	return buf.Bytes(), err
}
