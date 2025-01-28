package nodx

import (
	"bytes"
	"testing"

	"github.com/nodxdev/nodxgo/internal/assert"
)

func TestNodeGroup(t *testing.T) {
	t.Run("Render empty", func(t *testing.T) {
		group := newNodeGroup()
		buf := &bytes.Buffer{}

		err := group.Render(buf)
		assert.NoError(t, err)
		assert.Equal(t, "", buf.String())
	})

	t.Run("Render", func(t *testing.T) {
		group := newNodeGroup(newNodeText("Hello, World!"))
		expected := "Hello, World!"
		buf := &bytes.Buffer{}

		err := group.Render(buf)
		assert.NoError(t, err)
		assert.Equal(t, expected, buf.String())
	})

	t.Run("Render String", func(t *testing.T) {
		group := newNodeGroup(newNodeText("Hello, World!"))
		expected := "Hello, World!"
		got, err := group.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Render Bytes", func(t *testing.T) {
		group := newNodeGroup(newNodeText("Hello, World!"))
		expected := "Hello, World!"
		got, err := group.RenderBytes()
		assert.NoError(t, err)
		assert.Equal(t, expected, string(got))
	})

	t.Run("Render Stringer Interface", func(t *testing.T) {
		group := newNodeGroup(newNodeText("Hello, World!"))
		expected := "Hello, World!"
		assert.Equal(t, expected, group.String())
	})

	t.Run("Render group with multiple text nodes", func(t *testing.T) {
		group := newNodeGroup(
			newNodeText("Hello"),
			newNodeText(", "),
			newNodeText("World!"),
		)
		expected := "Hello, World!"
		buf := &bytes.Buffer{}

		err := group.Render(buf)
		assert.NoError(t, err)
		assert.Equal(t, expected, buf.String())
	})

	t.Run("Render group with mixed nodes", func(t *testing.T) {
		group := newNodeGroup(
			newNodeText("Hello, "),
			newNodeElement(false, "strong", newNodeText("World")),
			newNodeText("!"),
		)
		expected := "Hello, <strong>World</strong>!"
		buf := &bytes.Buffer{}

		err := group.Render(buf)
		assert.NoError(t, err)
		assert.Equal(t, expected, buf.String())
	})

	t.Run("Render group with nil node", func(t *testing.T) {
		group := newNodeGroup(
			newNodeText("Hello"),
			nil,
			newNodeText(", World!"),
		)
		expected := "Hello, World!"
		buf := &bytes.Buffer{}

		err := group.Render(buf)
		assert.NoError(t, err)
		assert.Equal(t, expected, buf.String())
	})

	t.Run("RenderString for group", func(t *testing.T) {
		group := newNodeGroup(
			newNodeText("Hello"),
			newNodeText(", "),
			newNodeText("World!"),
		)
		expected := "Hello, World!"

		got, err := group.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("RenderBytes for group", func(t *testing.T) {
		group := newNodeGroup(
			newNodeText("Hello"),
			newNodeText(", "),
			newNodeText("World!"),
		)
		expected := "Hello, World!"

		got, err := group.RenderBytes()
		assert.NoError(t, err)
		assert.Equal(t, expected, string(got))
	})

	t.Run("Group with nested groups", func(t *testing.T) {
		group := newNodeGroup(
			newNodeGroup(
				newNodeText("Hello, "),
				newNodeElement(false, "strong", newNodeText("World")),
			),
			newNodeText("!"),
		)
		expected := "Hello, <strong>World</strong>!"

		buf := &bytes.Buffer{}
		err := group.Render(buf)
		assert.NoError(t, err)
		assert.Equal(t, expected, buf.String())
	})

	t.Run("Empty group RenderString", func(t *testing.T) {
		group := newNodeGroup()
		expected := ""

		got, err := group.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Empty group RenderBytes", func(t *testing.T) {
		group := newNodeGroup()
		expected := ""

		got, err := group.RenderBytes()
		assert.NoError(t, err)
		assert.Equal(t, expected, string(got))
	})

	t.Run("Expand group as div child element", func(t *testing.T) {
		template := El(
			"div",
			newNodeGroup(
				newNodeText("Hello, World!"),
			),
		)
		expected := "<div>Hello, World!</div>"

		got, err := template.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Expand group as div child attribute", func(t *testing.T) {
		template := El(
			"div",
			newNodeGroup(
				Attr("class", "test"),
			),
			newNodeText("Hello, World!"),
		)
		expected := `<div class="test">Hello, World!</div>`

		got, err := template.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Expand group as div child element and attribute", func(t *testing.T) {
		template := El(
			"div",
			newNodeGroup(
				Attr("class", "test"),
				newNodeText("Hello, World!"),
			),
		)
		expected := `<div class="test">Hello, World!</div>`

		got, err := template.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
}
