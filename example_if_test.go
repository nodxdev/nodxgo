package nodx

import "fmt"

func ExampleIf() {
	condition := 3 > 2

	node := Div(
		Class("container"),
		If(condition, Text("Condition is true")),
	)

	fmt.Println(node)
	// Output:
	// <div class="container">Condition is true</div>
}
