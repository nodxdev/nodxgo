package nodx

import "fmt"

func ExampleText() {
	node := Div(
		Text("<script>alert('Hello, World!')</script>"),
	)

	fmt.Println(node)
	// Output:
	// <div>&lt;script&gt;alert(&#39;Hello, World!&#39;)&lt;/script&gt;</div>
}

func ExampleTextf() {
	node := Div(
		Textf("<script>alert('Hello, %s!')</script>", "World"),
	)

	fmt.Println(node)
	// Output:
	// <div>&lt;script&gt;alert(&#39;Hello, World!&#39;)&lt;/script&gt;</div>
}

func ExampleRaw() {
	node := Div(
		Raw("<script>alert('Hello, World!')</script>"),
	)

	fmt.Println(node)
	// Output:
	// <div><script>alert('Hello, World!')</script></div>
}

func ExampleRawf() {
	node := Div(
		Rawf("<script>alert('Hello, %s!')</script>", "World"),
	)

	fmt.Println(node)
	// Output:
	// <div><script>alert('Hello, World!')</script></div>
}
