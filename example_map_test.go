package nodx

import "fmt"

func ExampleMap() {
	items := []string{"Item 1", "Item 2", "Item 3"}

	node := Ul(
		Map(items, func(item string) Node {
			return Li(Text(item))
		}),
	)

	fmt.Println(node)
	// Output:
	// <ul><li>Item 1</li><li>Item 2</li><li>Item 3</li></ul>
}
