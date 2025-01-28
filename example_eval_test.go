package nodx

import "fmt"

func ExampleEval() {
	condition := 3 > 2

	node := Div(
		Class("container"),
		Eval(func() Node {
			if condition {
				return Text("Condition is true")
			}
			return Text("Condition is false")
		}),
	)

	fmt.Println(node)
	// Output:
	// <div class="container">Condition is true</div>
}
