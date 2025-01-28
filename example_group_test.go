package nodx

import "fmt"

func ExampleGroup() {
	node := Group(
		DocType(),
		Html(
			Head(
				TitleEl(Text("Hello, World!")),
			),
			Body(Text("Hello, World!")),
		),
	)

	fmt.Println(node)
	// Output:
	// <!DOCTYPE html><html><head><title>Hello, World!</title></head><body>Hello, World!</body></html>
}
