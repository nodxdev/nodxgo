package nodx

import "fmt"

func ExampleEl() {
	node := El("my-custom-element",
		Text("Hello, World!"),
	)

	fmt.Println(node)
	// Output:
	// <my-custom-element>Hello, World!</my-custom-element>
}
