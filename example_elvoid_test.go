package nodx

import "fmt"

func ExampleElVoid() {
	node := ElVoid("my-custom-element",
		Class("container"),
	)

	fmt.Println(node)
	// Output:
	// <my-custom-element class="container">
}
