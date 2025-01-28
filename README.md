# NodX for Go

<a href="https://pkg.go.dev/github.com/nodxdev/nodxgo">
  <img src="https://pkg.go.dev/badge/github.com/nodxdev/nodxgo" alt="Go Reference"/>
</a>
<a href="https://goreportcard.com/report/nodxdev/nodxgo">
  <img src="https://goreportcard.com/badge/nodxdev/nodxgo" alt="Go Report Card"/>
</a>
<a href="https://github.com/nodxdev/nodxgo/releases/latest">
  <img src="https://img.shields.io/github/release/nodxdev/nodxgo.svg" alt="Release Version"/>
</a>
<a href="LICENSE">
  <img src="https://img.shields.io/github/license/nodxdev/nodxgo.svg" alt="License"/>
</a>
<a href="https://github.com/nodxdev/nodxgo">
  <img src="https://img.shields.io/github/stars/nodxdev/nodxgo?style=flat&label=github+stars"/>
</a>

---

NodX is a modern and developer-friendly Go template engine for generating
**safe**, **clean**, and **maintainable** HTML. Designed for maximum
productivity and easy maintenance, it combines **simplicity**, **type safety**
and **robust formatting**, making it the perfect fit for your Go-based web
projects.

## Key Features

- **Type Safety 🛡️**: Fully typed APIs ensure you write error-free HTML, even at
  scale.
- **Robust Formatting 🧹**: Works seamlessly with `go fmt` to keep your code
  clean and consistent.
- **Zero Dependencies 📦**: Lightweight and fast, with no external dependencies.
- **DX in mind 🧠**: If you can write HTML and Go, you can write NodX.
- **Security by Default 🔒**: Automatically escapes unsafe text to prevent XSS
  vulnerabilities.

## Quick Start

Install the library:

```bash
# Go 1.22 or later is required
go get github.com/nodxdev/nodxgo
```

Start building your HTML with intuitive, type-safe functions:

```go
package main

import (
  "os"
  "github.com/nodxdev/nodxgo"
)

func main() {
  // you can fetch this from a database or some other source
  happiness := 100
  hideContainer := false

  myTemplate := nodx.Group(
    nodx.DocType(),
    nodx.Html(
      nodx.Head(
        nodx.TitleEl(nodx.Text("My NodX Page")),
      ),
      nodx.Body(
        nodx.Div(
          nodx.ClassMap{
            "container": true,
            "hidden":    hideContainer,
          },
          nodx.H1(
            nodx.Class("title"),
            nodx.Textf("Welcome to %s!", "NodX"),
          ),
          nodx.P(nodx.Text("This is a type-safe HTML generator for Go.")),
          nodx.If(happiness > 90, nodx.P(nodx.Textf("With NodX, you will be %d%% happy!", happiness))),
        ),
      ),
    ),
  )

  _ = myTemplate.Render(os.Stdout)
  // or
  // str, err := myTemplate.RenderString()
  // byt, err := myTemplate.RenderBytes()
}
```

### Output:

```html
<!DOCTYPE html>
<html>
  <head>
    <title>My NodX Page</title>
  </head>
  <body>
    <div class="container">
      <h1 class="title">Welcome to NodX!</h1>
      <p>This is a type-safe HTML generator for Go.</p>
      <p>With NodX, you will be 100% happy!</p>
    </div>
  </body>
</html>
```

## Key Concepts

### 1. **Elements made easy**

Every HTML tag is a function! Just call the function with child elements,
attributes, groups or text.

```go
nodx.Div(
  nodx.Class("container"),
  nodx.H1(nodx.Text("Hello, NodX!")),
  nodx.P(nodx.Text("Build clean and safe HTML effortlessly.")),
)
```

### 2. **Attributes with helpers**

Attributes like `class`, `src`, and `alt` have their own functions for
simplicity.

```go
nodx.Img(
  nodx.Src("image.jpg"),
  nodx.Alt("A beautiful image"),
)
```

### 3. **Dynamic class management**

Use `ClassMap` to conditionally apply classes based on your logic.

```go
nodx.Div(
  nodx.ClassMap{
    "visible": true,
    "hidden":  false,
  },
  nodx.Text("Conditional classes made simple!"),
)
```

### 4. **Fully typed server rendered components**

You can create your own components to avoid code duplication and keep your code
clean.

```go
func button(text string) nodx.Node {
  return nodx.Button(
    nodx.Class("btn"),
    nodx.Text(text),
  )
}

func main() {
  myTemplate := nodx.Div(
    button("Click me 1!"),
    button("Click me 2!"),
    button("Click me 3!"),
  )

  _ := myTemplate.Render(os.Stdout)

  /*
    Output:
    <div>
      <button class="btn">Click me 1!</button>
      <button class="btn">Click me 2!</button>
      <button class="btn">Click me 3!</button>
    </div>
  */
}
```

### 5. **Advanced features**

Please refer to the
[Full Documentation](https://pkg.go.dev/github.com/nodxdev/nodxgo) to read more
about all the included features.

- **Custom components**: You can create your own components to avoid code
  duplication and keep your code clean.
- **Conditional rendering**: You can use the `nodx.If` and `nodx.IfFunc`
  function to conditionally render elements based on your logic.
- **Map**: You can use the `nodx.Map` function to loop over a list of items and
  render them as a list of Nodes.
- **Custom elements and attributes**: You can use `nodx.El` and `nodx.Attr` to
  create your own elements and attributes if they are not included in the
  library.
- **Component library**: You can write your own component library and publish it
  for others to use.
- **Dynamic classes and styles**: You can use the `nodx.ClassMap` and
  `nodx.StyleMap` to conditionally apply classes and styles based on your own
  logic.

## Why Choose NodX?

- **Expressive and Clear Code**: Great helpers (`Div`, `H1`, `P`, etc.) that
  mirror HTML semantics, making your Go code as readable as HTML.
- **Battle-Tested for Safety**: Text is escaped automatically to protect your
  app from XSS vulnerabilities.
- **Lightweight and Efficient**: With no dependencies, you can focus on building
  without bloat.
- **Perfect for Modern Go Developers**: Designed with Go's simplicity and
  elegance in mind.

## License

NodX is open-source and available under the [MIT License](LICENSE). Feel free to
use it in your personal or commercial projects.

---

Start building safe, elegant, and maintainable HTML templates in Go today with
**NodX**!
