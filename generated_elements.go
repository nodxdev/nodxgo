package nodx

// Code generated by NodX. DO NOT EDIT.

// A defines a hyperlink.
//
//	func ExampleA() {
//		node := nodx.A(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <a id="1">Hello World!</a>
//	}
func A(children ...Node) Node {
	return El("a", children...)
}

// Abbr defines an abbreviation or acronym.
//
//	func ExampleAbbr() {
//		node := nodx.Abbr(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <abbr id="1">Hello World!</abbr>
//	}
func Abbr(children ...Node) Node {
	return El("abbr", children...)
}

// Address defines contact information for the author/owner of a document.
//
//	func ExampleAddress() {
//		node := nodx.Address(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <address id="1">Hello World!</address>
//	}
func Address(children ...Node) Node {
	return El("address", children...)
}

// Applet defines an embedded applet (deprecated).
//
//	func ExampleApplet() {
//		node := nodx.Applet(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <applet id="1">Hello World!</applet>
//	}
func Applet(children ...Node) Node {
	return El("applet", children...)
}

// Area defines an area inside an image-map.
//
//	func ExampleArea() {
//		node := nodx.Area(
//			nodx.Id( "1"),
//		)
//		fmt.Println(node)
//		// Output: <area id="1">
//	}
func Area(children ...Node) Node {
	return ElVoid("area", children...)
}

// Article defines an article.
//
//	func ExampleArticle() {
//		node := nodx.Article(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <article id="1">Hello World!</article>
//	}
func Article(children ...Node) Node {
	return El("article", children...)
}

// Aside defines content aside from the page content.
//
//	func ExampleAside() {
//		node := nodx.Aside(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <aside id="1">Hello World!</aside>
//	}
func Aside(children ...Node) Node {
	return El("aside", children...)
}

// Audio defines embedded sound content.
//
//	func ExampleAudio() {
//		node := nodx.Audio(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <audio id="1">Hello World!</audio>
//	}
func Audio(children ...Node) Node {
	return El("audio", children...)
}

// B defines bold text.
//
//	func ExampleB() {
//		node := nodx.B(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <b id="1">Hello World!</b>
//	}
func B(children ...Node) Node {
	return El("b", children...)
}

// Base specifies the base URL for all relative URLs in a document.
//
//	func ExampleBase() {
//		node := nodx.Base(
//			nodx.Id( "1"),
//		)
//		fmt.Println(node)
//		// Output: <base id="1">
//	}
func Base(children ...Node) Node {
	return ElVoid("base", children...)
}

// Basefont specifies a default color, size, and font for all text in a document (deprecated).
//
//	func ExampleBasefont() {
//		node := nodx.Basefont(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <basefont id="1">Hello World!</basefont>
//	}
func Basefont(children ...Node) Node {
	return El("basefont", children...)
}

// Bdi isolates a part of text that might be formatted in a different direction from other text.
//
//	func ExampleBdi() {
//		node := nodx.Bdi(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <bdi id="1">Hello World!</bdi>
//	}
func Bdi(children ...Node) Node {
	return El("bdi", children...)
}

// Bdo overrides the current text direction.
//
//	func ExampleBdo() {
//		node := nodx.Bdo(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <bdo id="1">Hello World!</bdo>
//	}
func Bdo(children ...Node) Node {
	return El("bdo", children...)
}

// Blockquote defines a section that is quoted from another source.
//
//	func ExampleBlockquote() {
//		node := nodx.Blockquote(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <blockquote id="1">Hello World!</blockquote>
//	}
func Blockquote(children ...Node) Node {
	return El("blockquote", children...)
}

// Body defines the body of the document.
//
//	func ExampleBody() {
//		node := nodx.Body(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <body id="1">Hello World!</body>
//	}
func Body(children ...Node) Node {
	return El("body", children...)
}

// Br inserts a single line break.
//
//	func ExampleBr() {
//		node := nodx.Br(
//			nodx.Id( "1"),
//		)
//		fmt.Println(node)
//		// Output: <br id="1">
//	}
func Br(children ...Node) Node {
	return ElVoid("br", children...)
}

// Button defines a clickable button.
//
//	func ExampleButton() {
//		node := nodx.Button(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <button id="1">Hello World!</button>
//	}
func Button(children ...Node) Node {
	return El("button", children...)
}

// Canvas used to draw graphics on the fly via scripting.
//
//	func ExampleCanvas() {
//		node := nodx.Canvas(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <canvas id="1">Hello World!</canvas>
//	}
func Canvas(children ...Node) Node {
	return El("canvas", children...)
}

// Caption defines a table caption.
//
//	func ExampleCaption() {
//		node := nodx.Caption(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <caption id="1">Hello World!</caption>
//	}
func Caption(children ...Node) Node {
	return El("caption", children...)
}

// Center defines centered text (deprecated).
//
//	func ExampleCenter() {
//		node := nodx.Center(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <center id="1">Hello World!</center>
//	}
func Center(children ...Node) Node {
	return El("center", children...)
}

// CiteEl defines the title of a work.
//
//	func ExampleCiteEl() {
//		node := nodx.CiteEl(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <cite id="1">Hello World!</cite>
//	}
func CiteEl(children ...Node) Node {
	return El("cite", children...)
}

// CodeEl defines a piece of computer code.
//
//	func ExampleCodeEl() {
//		node := nodx.CodeEl(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <code id="1">Hello World!</code>
//	}
func CodeEl(children ...Node) Node {
	return El("code", children...)
}

// Col specifies column properties for each column within a colgroup element.
//
//	func ExampleCol() {
//		node := nodx.Col(
//			nodx.Id( "1"),
//		)
//		fmt.Println(node)
//		// Output: <col id="1">
//	}
func Col(children ...Node) Node {
	return ElVoid("col", children...)
}

// Colgroup specifies a group of one or more columns in a table for formatting.
//
//	func ExampleColgroup() {
//		node := nodx.Colgroup(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <colgroup id="1">Hello World!</colgroup>
//	}
func Colgroup(children ...Node) Node {
	return El("colgroup", children...)
}

// DataEl links the given content with a machine-readable translation.
//
//	func ExampleDataEl() {
//		node := nodx.DataEl(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <data id="1">Hello World!</data>
//	}
func DataEl(children ...Node) Node {
	return El("data", children...)
}

// Datalist specifies a list of pre-defined options for input controls.
//
//	func ExampleDatalist() {
//		node := nodx.Datalist(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <datalist id="1">Hello World!</datalist>
//	}
func Datalist(children ...Node) Node {
	return El("datalist", children...)
}

// Dd defines a description/value of a term in a description list.
//
//	func ExampleDd() {
//		node := nodx.Dd(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <dd id="1">Hello World!</dd>
//	}
func Dd(children ...Node) Node {
	return El("dd", children...)
}

// Del defines text that has been deleted from a document.
//
//	func ExampleDel() {
//		node := nodx.Del(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <del id="1">Hello World!</del>
//	}
func Del(children ...Node) Node {
	return El("del", children...)
}

// Details defines additional details that the user can view or hide.
//
//	func ExampleDetails() {
//		node := nodx.Details(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <details id="1">Hello World!</details>
//	}
func Details(children ...Node) Node {
	return El("details", children...)
}

// Dfn represents the defining instance of a term.
//
//	func ExampleDfn() {
//		node := nodx.Dfn(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <dfn id="1">Hello World!</dfn>
//	}
func Dfn(children ...Node) Node {
	return El("dfn", children...)
}

// Dialog defines a dialog box or window.
//
//	func ExampleDialog() {
//		node := nodx.Dialog(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <dialog id="1">Hello World!</dialog>
//	}
func Dialog(children ...Node) Node {
	return El("dialog", children...)
}

// DirEl defines a directory list (deprecated).
//
//	func ExampleDirEl() {
//		node := nodx.DirEl(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <dir id="1">Hello World!</dir>
//	}
func DirEl(children ...Node) Node {
	return El("dir", children...)
}

// Div defines a section or a division in a document.
//
//	func ExampleDiv() {
//		node := nodx.Div(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <div id="1">Hello World!</div>
//	}
func Div(children ...Node) Node {
	return El("div", children...)
}

// Dl defines a description list.
//
//	func ExampleDl() {
//		node := nodx.Dl(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <dl id="1">Hello World!</dl>
//	}
func Dl(children ...Node) Node {
	return El("dl", children...)
}

// Dt defines a term/name in a description list.
//
//	func ExampleDt() {
//		node := nodx.Dt(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <dt id="1">Hello World!</dt>
//	}
func Dt(children ...Node) Node {
	return El("dt", children...)
}

// Em defines emphasized text.
//
//	func ExampleEm() {
//		node := nodx.Em(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <em id="1">Hello World!</em>
//	}
func Em(children ...Node) Node {
	return El("em", children...)
}

// Embed defines a container for an external application or interactive content.
//
//	func ExampleEmbed() {
//		node := nodx.Embed(
//			nodx.Id( "1"),
//		)
//		fmt.Println(node)
//		// Output: <embed id="1">
//	}
func Embed(children ...Node) Node {
	return ElVoid("embed", children...)
}

// Fieldset groups related elements in a form.
//
//	func ExampleFieldset() {
//		node := nodx.Fieldset(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <fieldset id="1">Hello World!</fieldset>
//	}
func Fieldset(children ...Node) Node {
	return El("fieldset", children...)
}

// Figcaption defines a caption for a figure element.
//
//	func ExampleFigcaption() {
//		node := nodx.Figcaption(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <figcaption id="1">Hello World!</figcaption>
//	}
func Figcaption(children ...Node) Node {
	return El("figcaption", children...)
}

// Figure specifies self-contained content.
//
//	func ExampleFigure() {
//		node := nodx.Figure(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <figure id="1">Hello World!</figure>
//	}
func Figure(children ...Node) Node {
	return El("figure", children...)
}

// Font defines font, color, and size for text (deprecated).
//
//	func ExampleFont() {
//		node := nodx.Font(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <font id="1">Hello World!</font>
//	}
func Font(children ...Node) Node {
	return El("font", children...)
}

// Footer defines a footer for a document or section.
//
//	func ExampleFooter() {
//		node := nodx.Footer(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <footer id="1">Hello World!</footer>
//	}
func Footer(children ...Node) Node {
	return El("footer", children...)
}

// FormEl defines an HTML form for user input.
//
//	func ExampleFormEl() {
//		node := nodx.FormEl(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <form id="1">Hello World!</form>
//	}
func FormEl(children ...Node) Node {
	return El("form", children...)
}

// Frame defines a window (a frame) in a frameset (deprecated).
//
//	func ExampleFrame() {
//		node := nodx.Frame(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <frame id="1">Hello World!</frame>
//	}
func Frame(children ...Node) Node {
	return El("frame", children...)
}

// Frameset defines a set of frames (deprecated).
//
//	func ExampleFrameset() {
//		node := nodx.Frameset(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <frameset id="1">Hello World!</frameset>
//	}
func Frameset(children ...Node) Node {
	return El("frameset", children...)
}

// H1 defines HTML headings level 1.
//
//	func ExampleH1() {
//		node := nodx.H1(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <h1 id="1">Hello World!</h1>
//	}
func H1(children ...Node) Node {
	return El("h1", children...)
}

// H2 defines HTML headings level 2.
//
//	func ExampleH2() {
//		node := nodx.H2(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <h2 id="1">Hello World!</h2>
//	}
func H2(children ...Node) Node {
	return El("h2", children...)
}

// H3 defines HTML headings level 3.
//
//	func ExampleH3() {
//		node := nodx.H3(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <h3 id="1">Hello World!</h3>
//	}
func H3(children ...Node) Node {
	return El("h3", children...)
}

// H4 defines HTML headings level 4.
//
//	func ExampleH4() {
//		node := nodx.H4(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <h4 id="1">Hello World!</h4>
//	}
func H4(children ...Node) Node {
	return El("h4", children...)
}

// H5 defines HTML headings level 5.
//
//	func ExampleH5() {
//		node := nodx.H5(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <h5 id="1">Hello World!</h5>
//	}
func H5(children ...Node) Node {
	return El("h5", children...)
}

// H6 defines HTML headings level 6.
//
//	func ExampleH6() {
//		node := nodx.H6(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <h6 id="1">Hello World!</h6>
//	}
func H6(children ...Node) Node {
	return El("h6", children...)
}

// Head contains metadata/information for the document.
//
//	func ExampleHead() {
//		node := nodx.Head(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <head id="1">Hello World!</head>
//	}
func Head(children ...Node) Node {
	return El("head", children...)
}

// Header defines a header for a document or section.
//
//	func ExampleHeader() {
//		node := nodx.Header(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <header id="1">Hello World!</header>
//	}
func Header(children ...Node) Node {
	return El("header", children...)
}

// Hr defines a thematic change in the content.
//
//	func ExampleHr() {
//		node := nodx.Hr(
//			nodx.Id( "1"),
//		)
//		fmt.Println(node)
//		// Output: <hr id="1">
//	}
func Hr(children ...Node) Node {
	return ElVoid("hr", children...)
}

// Html defines the root of an HTML document.
//
//	func ExampleHtml() {
//		node := nodx.Html(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <html id="1">Hello World!</html>
//	}
func Html(children ...Node) Node {
	return El("html", children...)
}

// I defines a part of text in an alternate voice or mood.
//
//	func ExampleI() {
//		node := nodx.I(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <i id="1">Hello World!</i>
//	}
func I(children ...Node) Node {
	return El("i", children...)
}

// Iframe defines an inline frame.
//
//	func ExampleIframe() {
//		node := nodx.Iframe(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <iframe id="1">Hello World!</iframe>
//	}
func Iframe(children ...Node) Node {
	return El("iframe", children...)
}

// Img defines an image.
//
//	func ExampleImg() {
//		node := nodx.Img(
//			nodx.Id( "1"),
//		)
//		fmt.Println(node)
//		// Output: <img id="1">
//	}
func Img(children ...Node) Node {
	return ElVoid("img", children...)
}

// Input defines an input control.
//
//	func ExampleInput() {
//		node := nodx.Input(
//			nodx.Id( "1"),
//		)
//		fmt.Println(node)
//		// Output: <input id="1">
//	}
func Input(children ...Node) Node {
	return ElVoid("input", children...)
}

// Ins defines a text that has been inserted into a document.
//
//	func ExampleIns() {
//		node := nodx.Ins(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <ins id="1">Hello World!</ins>
//	}
func Ins(children ...Node) Node {
	return El("ins", children...)
}

// Kbd defines keyboard input.
//
//	func ExampleKbd() {
//		node := nodx.Kbd(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <kbd id="1">Hello World!</kbd>
//	}
func Kbd(children ...Node) Node {
	return El("kbd", children...)
}

// LabelEl defines a label for an input element.
//
//	func ExampleLabelEl() {
//		node := nodx.LabelEl(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <label id="1">Hello World!</label>
//	}
func LabelEl(children ...Node) Node {
	return El("label", children...)
}

// Legend defines a caption for a fieldset element.
//
//	func ExampleLegend() {
//		node := nodx.Legend(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <legend id="1">Hello World!</legend>
//	}
func Legend(children ...Node) Node {
	return El("legend", children...)
}

// Li defines a list item.
//
//	func ExampleLi() {
//		node := nodx.Li(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <li id="1">Hello World!</li>
//	}
func Li(children ...Node) Node {
	return El("li", children...)
}

// Link defines the relationship between a document and an external resource.
//
//	func ExampleLink() {
//		node := nodx.Link(
//			nodx.Id( "1"),
//		)
//		fmt.Println(node)
//		// Output: <link id="1">
//	}
func Link(children ...Node) Node {
	return ElVoid("link", children...)
}

// Main specifies the main content of a document.
//
//	func ExampleMain() {
//		node := nodx.Main(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <main id="1">Hello World!</main>
//	}
func Main(children ...Node) Node {
	return El("main", children...)
}

// MapEl defines an image-map.
//
//	func ExampleMapEl() {
//		node := nodx.MapEl(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <map id="1">Hello World!</map>
//	}
func MapEl(children ...Node) Node {
	return El("map", children...)
}

// Mark defines marked or highlighted text.
//
//	func ExampleMark() {
//		node := nodx.Mark(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <mark id="1">Hello World!</mark>
//	}
func Mark(children ...Node) Node {
	return El("mark", children...)
}

// Meta defines metadata about an HTML document.
//
//	func ExampleMeta() {
//		node := nodx.Meta(
//			nodx.Id( "1"),
//		)
//		fmt.Println(node)
//		// Output: <meta id="1">
//	}
func Meta(children ...Node) Node {
	return ElVoid("meta", children...)
}

// Meter defines a scalar measurement within a known range.
//
//	func ExampleMeter() {
//		node := nodx.Meter(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <meter id="1">Hello World!</meter>
//	}
func Meter(children ...Node) Node {
	return El("meter", children...)
}

// Nav defines navigation links.
//
//	func ExampleNav() {
//		node := nodx.Nav(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <nav id="1">Hello World!</nav>
//	}
func Nav(children ...Node) Node {
	return El("nav", children...)
}

// Noframes defines an alternate content for users that do not support frames (deprecated).
//
//	func ExampleNoframes() {
//		node := nodx.Noframes(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <noframes id="1">Hello World!</noframes>
//	}
func Noframes(children ...Node) Node {
	return El("noframes", children...)
}

// Noscript defines an alternate content for users that do not support client-side scripts.
//
//	func ExampleNoscript() {
//		node := nodx.Noscript(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <noscript id="1">Hello World!</noscript>
//	}
func Noscript(children ...Node) Node {
	return El("noscript", children...)
}

// Object defines an embedded object.
//
//	func ExampleObject() {
//		node := nodx.Object(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <object id="1">Hello World!</object>
//	}
func Object(children ...Node) Node {
	return El("object", children...)
}

// Ol defines an ordered list.
//
//	func ExampleOl() {
//		node := nodx.Ol(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <ol id="1">Hello World!</ol>
//	}
func Ol(children ...Node) Node {
	return El("ol", children...)
}

// Optgroup defines a group of related options in a drop-down list.
//
//	func ExampleOptgroup() {
//		node := nodx.Optgroup(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <optgroup id="1">Hello World!</optgroup>
//	}
func Optgroup(children ...Node) Node {
	return El("optgroup", children...)
}

// Option defines an option in a drop-down list.
//
//	func ExampleOption() {
//		node := nodx.Option(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <option id="1">Hello World!</option>
//	}
func Option(children ...Node) Node {
	return El("option", children...)
}

// Output defines the result of a calculation.
//
//	func ExampleOutput() {
//		node := nodx.Output(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <output id="1">Hello World!</output>
//	}
func Output(children ...Node) Node {
	return El("output", children...)
}

// P defines a paragraph.
//
//	func ExampleP() {
//		node := nodx.P(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <p id="1">Hello World!</p>
//	}
func P(children ...Node) Node {
	return El("p", children...)
}

// Param defines a parameter for an object.
//
//	func ExampleParam() {
//		node := nodx.Param(
//			nodx.Id( "1"),
//		)
//		fmt.Println(node)
//		// Output: <param id="1">
//	}
func Param(children ...Node) Node {
	return ElVoid("param", children...)
}

// Picture defines a container for multiple image resources.
//
//	func ExamplePicture() {
//		node := nodx.Picture(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <picture id="1">Hello World!</picture>
//	}
func Picture(children ...Node) Node {
	return El("picture", children...)
}

// Pre defines preformatted text.
//
//	func ExamplePre() {
//		node := nodx.Pre(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <pre id="1">Hello World!</pre>
//	}
func Pre(children ...Node) Node {
	return El("pre", children...)
}

// Progress represents the progress of a task.
//
//	func ExampleProgress() {
//		node := nodx.Progress(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <progress id="1">Hello World!</progress>
//	}
func Progress(children ...Node) Node {
	return El("progress", children...)
}

// Q defines a short quotation.
//
//	func ExampleQ() {
//		node := nodx.Q(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <q id="1">Hello World!</q>
//	}
func Q(children ...Node) Node {
	return El("q", children...)
}

// Rb used to delimit the base text component of a ruby annotation.
//
//	func ExampleRb() {
//		node := nodx.Rb(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <rb id="1">Hello World!</rb>
//	}
func Rb(children ...Node) Node {
	return El("rb", children...)
}

// Rp defines what to show in browsers that do not support ruby annotations.
//
//	func ExampleRp() {
//		node := nodx.Rp(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <rp id="1">Hello World!</rp>
//	}
func Rp(children ...Node) Node {
	return El("rp", children...)
}

// Rt defines an explanation/pronunciation of characters (for East Asian typography).
//
//	func ExampleRt() {
//		node := nodx.Rt(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <rt id="1">Hello World!</rt>
//	}
func Rt(children ...Node) Node {
	return El("rt", children...)
}

// Rtc defines a ruby text container for multiple rt elements.
//
//	func ExampleRtc() {
//		node := nodx.Rtc(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <rtc id="1">Hello World!</rtc>
//	}
func Rtc(children ...Node) Node {
	return El("rtc", children...)
}

// Ruby defines a ruby annotation.
//
//	func ExampleRuby() {
//		node := nodx.Ruby(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <ruby id="1">Hello World!</ruby>
//	}
func Ruby(children ...Node) Node {
	return El("ruby", children...)
}

// S defines text that is no longer correct.
//
//	func ExampleS() {
//		node := nodx.S(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <s id="1">Hello World!</s>
//	}
func S(children ...Node) Node {
	return El("s", children...)
}

// Samp defines sample output from a computer program.
//
//	func ExampleSamp() {
//		node := nodx.Samp(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <samp id="1">Hello World!</samp>
//	}
func Samp(children ...Node) Node {
	return El("samp", children...)
}

// Script defines a client-side script.
//
//	func ExampleScript() {
//		node := nodx.Script(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <script id="1">Hello World!</script>
//	}
func Script(children ...Node) Node {
	return El("script", children...)
}

// Section defines a section in a document.
//
//	func ExampleSection() {
//		node := nodx.Section(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <section id="1">Hello World!</section>
//	}
func Section(children ...Node) Node {
	return El("section", children...)
}

// Select defines a drop-down list.
//
//	func ExampleSelect() {
//		node := nodx.Select(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <select id="1">Hello World!</select>
//	}
func Select(children ...Node) Node {
	return El("select", children...)
}

// SlotEl defines a placeholder inside a web component.
//
//	func ExampleSlotEl() {
//		node := nodx.SlotEl(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <slot id="1">Hello World!</slot>
//	}
func SlotEl(children ...Node) Node {
	return El("slot", children...)
}

// Small defines smaller text.
//
//	func ExampleSmall() {
//		node := nodx.Small(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <small id="1">Hello World!</small>
//	}
func Small(children ...Node) Node {
	return El("small", children...)
}

// Source defines multiple media resources for media elements.
//
//	func ExampleSource() {
//		node := nodx.Source(
//			nodx.Id( "1"),
//		)
//		fmt.Println(node)
//		// Output: <source id="1">
//	}
func Source(children ...Node) Node {
	return ElVoid("source", children...)
}

// SpanEl defines a section in a document.
//
//	func ExampleSpanEl() {
//		node := nodx.SpanEl(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <span id="1">Hello World!</span>
//	}
func SpanEl(children ...Node) Node {
	return El("span", children...)
}

// Strike defines strikethrough text (deprecated).
//
//	func ExampleStrike() {
//		node := nodx.Strike(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <strike id="1">Hello World!</strike>
//	}
func Strike(children ...Node) Node {
	return El("strike", children...)
}

// Strong defines important text.
//
//	func ExampleStrong() {
//		node := nodx.Strong(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <strong id="1">Hello World!</strong>
//	}
func Strong(children ...Node) Node {
	return El("strong", children...)
}

// StyleEl defines style information for a document.
//
//	func ExampleStyleEl() {
//		node := nodx.StyleEl(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <style id="1">Hello World!</style>
//	}
func StyleEl(children ...Node) Node {
	return El("style", children...)
}

// Sub defines subscripted text.
//
//	func ExampleSub() {
//		node := nodx.Sub(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <sub id="1">Hello World!</sub>
//	}
func Sub(children ...Node) Node {
	return El("sub", children...)
}

// SummaryEl defines a visible heading for a details element.
//
//	func ExampleSummaryEl() {
//		node := nodx.SummaryEl(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <summary id="1">Hello World!</summary>
//	}
func SummaryEl(children ...Node) Node {
	return El("summary", children...)
}

// Sup defines superscripted text.
//
//	func ExampleSup() {
//		node := nodx.Sup(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <sup id="1">Hello World!</sup>
//	}
func Sup(children ...Node) Node {
	return El("sup", children...)
}

// Svg defines a container for SVG graphics.
//
//	func ExampleSvg() {
//		node := nodx.Svg(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <svg id="1">Hello World!</svg>
//	}
func Svg(children ...Node) Node {
	return El("svg", children...)
}

// Table defines a table.
//
//	func ExampleTable() {
//		node := nodx.Table(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <table id="1">Hello World!</table>
//	}
func Table(children ...Node) Node {
	return El("table", children...)
}

// Tbody groups the body content in a table.
//
//	func ExampleTbody() {
//		node := nodx.Tbody(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <tbody id="1">Hello World!</tbody>
//	}
func Tbody(children ...Node) Node {
	return El("tbody", children...)
}

// Td defines a cell in a table.
//
//	func ExampleTd() {
//		node := nodx.Td(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <td id="1">Hello World!</td>
//	}
func Td(children ...Node) Node {
	return El("td", children...)
}

// Template defines a template.
//
//	func ExampleTemplate() {
//		node := nodx.Template(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <template id="1">Hello World!</template>
//	}
func Template(children ...Node) Node {
	return El("template", children...)
}

// Textarea defines a multiline input control.
//
//	func ExampleTextarea() {
//		node := nodx.Textarea(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <textarea id="1">Hello World!</textarea>
//	}
func Textarea(children ...Node) Node {
	return El("textarea", children...)
}

// Tfoot groups the footer content in a table.
//
//	func ExampleTfoot() {
//		node := nodx.Tfoot(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <tfoot id="1">Hello World!</tfoot>
//	}
func Tfoot(children ...Node) Node {
	return El("tfoot", children...)
}

// Th defines a header cell in a table.
//
//	func ExampleTh() {
//		node := nodx.Th(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <th id="1">Hello World!</th>
//	}
func Th(children ...Node) Node {
	return El("th", children...)
}

// Thead groups the header content in a table.
//
//	func ExampleThead() {
//		node := nodx.Thead(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <thead id="1">Hello World!</thead>
//	}
func Thead(children ...Node) Node {
	return El("thead", children...)
}

// Time defines a specific time.
//
//	func ExampleTime() {
//		node := nodx.Time(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <time id="1">Hello World!</time>
//	}
func Time(children ...Node) Node {
	return El("time", children...)
}

// TitleEl defines a title for the document.
//
//	func ExampleTitleEl() {
//		node := nodx.TitleEl(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <title id="1">Hello World!</title>
//	}
func TitleEl(children ...Node) Node {
	return El("title", children...)
}

// Tr defines a row in a table.
//
//	func ExampleTr() {
//		node := nodx.Tr(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <tr id="1">Hello World!</tr>
//	}
func Tr(children ...Node) Node {
	return El("tr", children...)
}

// Track defines text tracks for media elements.
//
//	func ExampleTrack() {
//		node := nodx.Track(
//			nodx.Id( "1"),
//		)
//		fmt.Println(node)
//		// Output: <track id="1">
//	}
func Track(children ...Node) Node {
	return ElVoid("track", children...)
}

// Tt defines teletype text (deprecated).
//
//	func ExampleTt() {
//		node := nodx.Tt(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <tt id="1">Hello World!</tt>
//	}
func Tt(children ...Node) Node {
	return El("tt", children...)
}

// U defines text that should be stylistically different from normal text.
//
//	func ExampleU() {
//		node := nodx.U(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <u id="1">Hello World!</u>
//	}
func U(children ...Node) Node {
	return El("u", children...)
}

// Ul defines an unordered list.
//
//	func ExampleUl() {
//		node := nodx.Ul(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <ul id="1">Hello World!</ul>
//	}
func Ul(children ...Node) Node {
	return El("ul", children...)
}

// Var defines a variable.
//
//	func ExampleVar() {
//		node := nodx.Var(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <var id="1">Hello World!</var>
//	}
func Var(children ...Node) Node {
	return El("var", children...)
}

// Video defines embedded video content.
//
//	func ExampleVideo() {
//		node := nodx.Video(
//			nodx.Id("1"),
//			nodx.Text("Hello World!"),
//		)
//		fmt.Println(node)
//		// Output: <video id="1">Hello World!</video>
//	}
func Video(children ...Node) Node {
	return El("video", children...)
}

// Wbr defines a possible line-break.
//
//	func ExampleWbr() {
//		node := nodx.Wbr(
//			nodx.Id( "1"),
//		)
//		fmt.Println(node)
//		// Output: <wbr id="1">
//	}
func Wbr(children ...Node) Node {
	return ElVoid("wbr", children...)
}
