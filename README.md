# Goldmark Latex

> Underscores `_` won't break LaTeX and Markdown any more

`goldmark-latex` is an extension for [goldmark](http://github.com/yuin/goldmark) that adds support for math blocks and inline math.

This is just a fork of [litao91/goldmark-mathjax](https://github.com/litao91/goldmark-mathjax) with more options for parsing and rendering. There are also some minor modifications (and I hope to modernizing that code in future commits).

(I decided to not make a direct fork because the name was misleading in referencing [MathJax](https://www.mathjax.org/) as the project isn't really used anywhere concretely inside the project. The goldmark extension just parsed latex code inside markdown and doesn't even add a CDN link to mathjax or something similar)

## Installation

```
go get github.com/aziis98/goldmark-latex
```

## Usage

The default config uses `$...$` for inline math and `$$...$$` for math blocks and renders the former to `<span class="math inline">...</span>` and the latter to `<div class="math block">...</span>`.

```go
markdown := goldmark.New(
    goldmark.WithExtensions(
        latex.NewLatex(),
    ),
)
```

All delimiters can be customized using the following functions, for example the following uses `\(...\)` and `\[...\]` and preserves them as is during render.

```go
markdown := goldmark.New(
    goldmark.WithExtensions(
        latex.NewLatex(
            latex.WithSourceInlineDelim(`\(`, `\)`),
            latex.WithSourceBlockDelim(`\[`, `\]`),
            latex.WithOutputInlineDelim(`\(`, `\)`),
            latex.WithOutputBlockDelim(`\[`, `\]`),
        ),
    ),
)
```

## TODOs

-   [ ] I ported this directly from `goldmark-mathjax` but I read too late that it was based on old goldmark code for [fenced blocks](https://github.com/yuin/goldmark/blob/master/parser/fcode_block.go) and [code spans](https://github.com/yuin/goldmark/blob/master/parser/code_span.go) that in the meantime change a bit.

    Sometime I'll update the parser code to not use the deprecated `util.DedentPosition`.

-   [ ] The old `goldmark-mathjax` actually had some code that compiled latex to SVGs using `pdflatex` and `pdf2svg`.

    Something interesting could be to use the "tex renderer" just for math blocks. Rendering inline math to SVGs has the problem that images can't wrap along lines, for those falling back to KaTeX or MathJax might be the best option.

## License

MIT
