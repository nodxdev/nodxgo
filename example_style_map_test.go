package nodx

import "fmt"

func ExampleStyleMap() {
	shouldHide := 3 > 2

	node := Div(
		StyleMap{
			"color: red":             true,
			"border: 1px solid blue": false,
			"display: none":          shouldHide,
		},
		Text("Hello, World!"),
	)

	fmt.Println(node)
	// Output:
	// <div style="color: red; display: none">Hello, World!</div>
}
