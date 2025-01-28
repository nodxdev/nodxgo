package nodx

import (
	"testing"

	"github.com/nodxdev/nodxgo/internal/assert"
)

func TestStyleMap(t *testing.T) {
	t.Run("Render empty StyleMap", func(t *testing.T) {
		sm := StyleMap{}
		expected := ` style=""`

		got, err := sm.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Render StyleMap with single true style", func(t *testing.T) {
		sm := StyleMap{
			"border: 1px solid black": true,
		}
		expected := ` style="border: 1px solid black"`

		got, err := sm.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Render StyleMap with multiple true styles", func(t *testing.T) {
		sm := StyleMap{
			"border: 1px solid black": true,
			"margin: 10px":            true,
			"color: red":              true,
		}
		expected := ` style="border: 1px solid black; color: red; margin: 10px"`

		got, err := sm.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Render StyleMap with mixed true and false styles", func(t *testing.T) {
		sm := StyleMap{
			"border: 1px solid black": true,
			"padding: 20px":           false,
			"margin: 10px":            true,
			"color: red":              false,
		}
		expected := ` style="border: 1px solid black; margin: 10px"`

		got, err := sm.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Render sorted StyleMap", func(t *testing.T) {
		sm := StyleMap{
			"z-index: 1":  true,
			"border: 1px": true,
			"margin: 5px": true,
		}
		expected := ` style="border: 1px; margin: 5px; z-index: 1"`

		got, err := sm.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("RenderBytes for StyleMap", func(t *testing.T) {
		sm := StyleMap{
			"border: 1px solid black": true,
			"margin: 10px":            true,
		}
		expected := ` style="border: 1px solid black; margin: 10px"`

		got, err := sm.RenderBytes()
		assert.NoError(t, err)
		assert.Equal(t, expected, string(got))
	})

	t.Run("StyleMap with no true values", func(t *testing.T) {
		sm := StyleMap{
			"border: 1px solid black": false,
			"margin: 10px":            false,
		}
		expected := ` style=""`

		got, err := sm.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("StyleMap with special characters in styles", func(t *testing.T) {
		sm := StyleMap{
			"content: 'Hello, World!'": true,
			"font-size: 16px":          true,
		}
		expected := ` style="content: &#39;Hello, World!&#39;; font-size: 16px"`

		got, err := sm.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("StyleMap with conditional styles", func(t *testing.T) {
		showBorder := true
		showPadding := false

		sm := StyleMap{
			"border: 1px solid black": showBorder,
			"padding: 20px":           showPadding,
			"margin: 5px":             true,
		}
		expected := ` style="border: 1px solid black; margin: 5px"`

		got, err := sm.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Render inside a node", func(t *testing.T) {
		template := El("div", StyleMap{
			"border: 1px solid black": true,
			"margin: 5px":             false,
		})
		expected := `<div style="border: 1px solid black"></div>`

		got, err := template.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Stringer interface", func(t *testing.T) {
		sm := StyleMap{
			"border: 1px solid black": true,
			"padding: 20px":           false,
			"margin: 5px":             true,
		}
		expected := ` style="border: 1px solid black; margin: 5px"`

		got := sm.String()
		assert.Equal(t, expected, got)
	})
}
