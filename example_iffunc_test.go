package nodx

import "fmt"

func ExampleIfFunc() {
	condition := 3 > 2

	node := Div(
		Class("container"),
		IfFunc(condition, func() Node {
			// You can add your own go code here
			return Text("Condition is true")
		}),
	)

	fmt.Println(node)
	// Output:
	// <div class="container">Condition is true</div>
}
