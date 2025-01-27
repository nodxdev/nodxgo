package nodx

import (
	"bytes"
	"testing"

	"github.com/nodxdev/nodxgo/internal/assert"
)

func TestNodeText(t *testing.T) {
	t.Run("Text render", func(t *testing.T) {
		text := "Hello, World!"

		node := newNodeText(text)
		buf := new(bytes.Buffer)

		err := node.Render(buf)
		assert.NoError(t, err)

		assert.Equal(t, buf.String(), text)
	})

	t.Run("Text render string", func(t *testing.T) {
		text := "Hello, World!"

		node := newNodeText(text)
		str, err := node.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, str, text)
	})

	t.Run("Text render bytes", func(t *testing.T) {
		text := "Hello, World!"

		node := newNodeText(text)
		bytes, err := node.RenderBytes()
		assert.NoError(t, err)
		assert.Equal(t, string(bytes), text)
	})

	t.Run("Empty text render", func(t *testing.T) {
		node := newNodeText("")
		buf := new(bytes.Buffer)
		err := node.Render(buf)
		assert.NoError(t, err)
		assert.Equal(t, buf.String(), "")
	})
}
