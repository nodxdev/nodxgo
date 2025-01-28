package nodx

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// ensure that group implements the Node interface.
var _ Node = (*nodeGroup)(nil)

// nodeGroup is a special Node that represents a group of nodes.
//
// When rendered directly, it will call Render on all the nodes in the group
// sequentially.
//
// When used as a child of another node, it will be expanded so that the nodes
// in the group become children of the group's parent.
//
// This is not an element nor an attribute and should be treated specially when
// needed.
type nodeGroup []Node

// newNodeGroup creates a new group of nodes.
func newNodeGroup(nodes ...Node) nodeGroup {
	return nodeGroup(nodes)
}

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

func (g nodeGroup) RenderString() (string, error) {
	buf := &strings.Builder{}
	err := g.Render(buf)
	return buf.String(), err
}

func (g nodeGroup) RenderBytes() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := g.Render(buf)
	return buf.Bytes(), err
}

func (g nodeGroup) IsElement() bool {
	return false
}

func (g nodeGroup) IsAttribute() bool {
	return false
}

func (g nodeGroup) String() string {
	str, _ := g.RenderString()
	return str
}
