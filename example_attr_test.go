package nodx

import "fmt"

func ExampleAttr() {
	node := Div(
		Attr("my-custom-attr", "foo"),
	)

	fmt.Println(node)
	// Output:
	// <div my-custom-attr="foo"></div>
}
