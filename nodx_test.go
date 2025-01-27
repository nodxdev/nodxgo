package nodx

import (
	"testing"

	"github.com/nodxdev/nodxgo/internal/assert"
)

func TestGroup(t *testing.T) {
	t.Run("Group with no nodes", func(t *testing.T) {
		group := Group()
		expected := ""

		got, err := group.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Group with multiple nodes", func(t *testing.T) {
		group := Group(Text("Hello"), Text("World"))
		expected := "HelloWorld"

		got, err := group.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Group with some nil nodes", func(t *testing.T) {
		group := Group(Text("Hello"), nil, Text("World"))
		expected := "HelloWorld"

		got, err := group.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Group with all nil nodes", func(t *testing.T) {
		group := Group(nil, nil, nil)
		expected := ""

		got, err := group.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Group with mixed children", func(t *testing.T) {
		group := Group(
			El("div",
				Attr("class", "container"),
				Text("Hello"),
				Text("World"),
				nil,
			),
			Text("Hello"),
			nil,
			Text("World"),
		)
		expected := `<div class="container">HelloWorld</div>HelloWorld`

		got, err := group.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
}

func TestEl(t *testing.T) {
	t.Run("El with no children", func(t *testing.T) {
		el := El("div")
		expected := "<div></div>"

		got, err := el.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("El with children", func(t *testing.T) {
		el := El(
			"div",
			Attr("class", "container"),
			Text("Hello"),
			Text("World"),
		)
		expected := `<div class="container">HelloWorld</div>`

		got, err := el.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
}

func TestElVoid(t *testing.T) {
	t.Run("ElVoid with no children", func(t *testing.T) {
		elVoid := ElVoid("img")
		expected := "<img/>"

		got, err := elVoid.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("ElVoid with ignored children", func(t *testing.T) {
		elVoid := ElVoid(
			"img",
			Attr("class", "container"),
			Text("Ignored"),
		)
		expected := `<img class="container"/>`

		got, err := elVoid.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
}

func TestAttr(t *testing.T) {
	t.Run("Attr with value", func(t *testing.T) {
		attr := Attr("class", "container")
		expected := ` class="container"`

		got, err := attr.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Attr with empty value", func(t *testing.T) {
		attr := Attr("class", "")
		expected := ` class=""`

		got, err := attr.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
}

func TestText(t *testing.T) {
	t.Run("Text with content", func(t *testing.T) {
		text := Text("Hello <b>World</b>")
		expected := "Hello &lt;b&gt;World&lt;/b&gt;"

		got, err := text.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Text with empty content", func(t *testing.T) {
		text := Text("")
		expected := ""

		got, err := text.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
}

func TestTextf(t *testing.T) {
	t.Run("Textf with formatted content", func(t *testing.T) {
		text := Textf("Hello, %s!", "World")
		expected := "Hello, World!"

		got, err := text.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Textf with escaped content", func(t *testing.T) {
		text := Textf("Hello <b>%s</b>", "World")
		expected := "Hello &lt;b&gt;World&lt;/b&gt;"

		got, err := text.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Textf with empty content", func(t *testing.T) {
		text := Textf("")
		expected := ""

		got, err := text.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
}

func TestRaw(t *testing.T) {
	t.Run("Raw with content", func(t *testing.T) {
		raw := Raw("<script>alert('XSS');</script>")
		expected := "<script>alert('XSS');</script>"

		got, err := raw.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Raw with empty content", func(t *testing.T) {
		raw := Raw("")
		expected := ""

		got, err := raw.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
}

func TestRawf(t *testing.T) {
	t.Run("Rawf with formatted content", func(t *testing.T) {
		raw := Rawf("<div>%s</div>", "Content")
		expected := "<div>Content</div>"

		got, err := raw.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
}

func TestIf(t *testing.T) {
	t.Run("If condition true", func(t *testing.T) {
		node := If(true, Text("Hello"))
		expected := "Hello"

		got, err := node.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("If condition false", func(t *testing.T) {
		node := If(false, Text("Hello"))
		expected := ""

		got, err := node.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
}

func TestIfFunc(t *testing.T) {
	t.Run("IfFunc condition true", func(t *testing.T) {
		node := IfFunc(true, func() Node {
			return Text("Hello")
		})
		expected := "Hello"

		got, err := node.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("IfFunc condition false", func(t *testing.T) {
		node := IfFunc(false, func() Node {
			return Text("Hello")
		})
		expected := ""

		got, err := node.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
}

func TestMap(t *testing.T) {
	t.Run("Map with integers", func(t *testing.T) {
		slice := []int{1, 2, 3}
		node := Map(slice, func(i int) Node {
			return Textf("%d", i)
		})
		expected := "123"

		got, err := node.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Map with strings", func(t *testing.T) {
		slice := []string{"a", "b", "c"}
		node := Map(slice, func(s string) Node {
			return Text(s)
		})
		expected := "abc"

		got, err := node.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Map with empty slice", func(t *testing.T) {
		slice := []int{}
		node := Map(slice, func(i int) Node {
			return Textf("%d", i)
		})
		expected := ""

		got, err := node.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})
}
