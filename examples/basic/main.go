package main

import (
	"bytes"
	"fmt"
	"log"

	latex "github.com/aziis98/goldmark-latex"
	"github.com/lithammer/dedent"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func main() {
	md := goldmark.New(
		goldmark.WithExtensions(
			latex.NewLatex(),
			extension.NewTypographer(),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	// todo more control on the parsing process
	var output bytes.Buffer
	source := dedent.Dedent(`
		# Title

		Math like "$a_0$" along _markdown_ -- just -- works!
	
		Inline math $\frac{1}{2}$

		$$
		\mathbb{E}(X) = \int x \mathrm d F(x) = 
		\left\{ 
		\begin{aligned} 
			\sum_x x f(x) \; & \text{ if } X \text{ is discrete} \\ 
			\int x f(x) \mathrm d x \; & \text{ if } X \text{ is continuous }
		\end{aligned} \right.
		$$

		-   Even nested

			$$
			1
			$$
			
			-   list item
				
				$$
				2
				$$
				
				-   work (mostly)
				
					$$
					3
					$$
	`)

	if err := md.Convert([]byte(source), &output); err != nil {
		log.Fatal(err)
	}

	fmt.Println(output.String())
}
