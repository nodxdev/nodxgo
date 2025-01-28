package nodx

import "fmt"

func ExampleClassMap() {
	shouldHide := 3 > 2

	node := Div(
		ClassMap{
			"container": true,
			"other":     false,
			"hidden":    shouldHide,
		},
		Text("Hello, World!"),
	)

	fmt.Println(node)
	// Output:
	// <div class="container hidden">Hello, World!</div>
}
