package nodx

import (
	"testing"

	"github.com/nodxdev/nodxgo/internal/assert"
)

func TestClassMap(t *testing.T) {
	t.Run("Render empty ClassMap", func(t *testing.T) {
		cm := ClassMap{}
		expected := `class=""`

		got, err := cm.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Render ClassMap with single true class", func(t *testing.T) {
		cm := ClassMap{
			"class1": true,
		}
		expected := `class="class1"`

		got, err := cm.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Render ClassMap with multiple true classes", func(t *testing.T) {
		cm := ClassMap{
			"class1": true,
			"class2": true,
			"class3": true,
		}
		expected := `class="class1 class2 class3"`

		got, err := cm.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Render ClassMap with mixed true and false classes", func(t *testing.T) {
		cm := ClassMap{
			"class1": true,
			"class2": false,
			"class3": true,
			"class4": false,
		}
		expected := `class="class1 class3"`

		got, err := cm.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Render ClassMap with conditions", func(t *testing.T) {
		isOdd := func(n int) bool { return n%2 != 0 }
		isEven := func(n int) bool { return n%2 == 0 }

		renderOdd := isOdd(3)
		renderEven := isEven(3)

		cm := ClassMap{
			"odd-class":  renderOdd,
			"even-class": renderEven,
			"always-on":  true,
		}
		expected := `class="always-on odd-class"`

		got, err := cm.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Render sorted ClassMap", func(t *testing.T) {
		cm := ClassMap{
			"z-class": true,
			"a-class": true,
			"m-class": true,
		}
		expected := `class="a-class m-class z-class"`

		got, err := cm.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("RenderBytes for ClassMap", func(t *testing.T) {
		cm := ClassMap{
			"class1": true,
			"class2": false,
			"class3": true,
		}
		expected := `class="class1 class3"`

		got, err := cm.RenderBytes()
		assert.NoError(t, err)
		assert.Equal(t, expected, string(got))
	})

	t.Run("Render to empty writer", func(t *testing.T) {
		cm := ClassMap{}
		expected := `class=""`

		got, err := cm.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("ClassMap with no true values", func(t *testing.T) {
		cm := ClassMap{
			"class1": false,
			"class2": false,
		}
		expected := `class=""`

		got, err := cm.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("ClassMap with special characters", func(t *testing.T) {
		cm := ClassMap{
			"class-1": true,
			"class_2": true,
			"3class":  true,
		}
		expected := `class="3class class-1 class_2"`

		got, err := cm.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Render inside a node", func(t *testing.T) {
		template := El("div", ClassMap{
			"class1": true,
			"class2": false,
		})
		expected := `<div class="class1"></div>`

		got, err := template.RenderString()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("Stringer interface", func(t *testing.T) {
		cm := ClassMap{
			"class1": true,
			"class2": false,
			"class3": true,
		}
		expected := `class="class1 class3"`

		got := cm.String()
		assert.Equal(t, expected, got)
	})
}
