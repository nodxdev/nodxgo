package nodx

import (
	"strings"
	"testing"

	"github.com/nodxdev/nodxgo/internal/assert"
)

func TestNodeText(t *testing.T) {
	t.Run("Text render", func(t *testing.T) {
		text := "Hello, World!"
		node := newNodeText(text)
		buf := &strings.Builder{}

		err := node.Render(buf)
		assert.NoError(t, err)
		assert.Equal(t, text, buf.String())
	})

	t.Run("Text render string", func(t *testing.T) {
		text := "Hello, World!"
		node := newNodeText(text)

		str, err := node.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, text, str)
	})

	t.Run("Text render bytes", func(t *testing.T) {
		text := "Hello, World!"
		node := newNodeText(text)

		bytes, err := node.RenderBytes()
		assert.NoError(t, err)
		assert.Equal(t, text, string(bytes))
	})

	t.Run("Empty text render", func(t *testing.T) {
		node := newNodeText("")
		buf := &strings.Builder{}

		err := node.Render(buf)
		assert.NoError(t, err)
		assert.Equal(t, "", buf.String())
	})

	t.Run("Not initialized nodeText", func(t *testing.T) {
		node := nodeText{}
		buf := &strings.Builder{}

		err := node.Render(buf)
		assert.NoError(t, err)
		assert.Equal(t, "", buf.String())
	})
}
