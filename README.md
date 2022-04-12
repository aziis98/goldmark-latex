# Goldmark Latex

> Underscores `_` won't break LaTeX and Markdown any more

`goldmark-latex` is an extension for [goldmark](http://github.com/yuin/goldmark) that adds support for math blocks and inline math.

This is just a fork of [litao91/goldmark-mathjax](https://github.com/litao91/goldmark-mathjax) with more options for parsing and rendering. There are also some minor modifications (and I hope to modernizing that code in future commits).

## Installation

```
go get github.com/aziis98/goldmark-latex
```

## Usage

The default config uses `$...$` for inline math and `$$...$$` for math blocks and renders the former to `<span class="math inline">...</span>` and the latter to `<div class="math block">...</span>`.

```go
markdown := goldmark.New(
    goldmark.WithExtensions(latex.DefaultConfig),
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

-   [ ] I ported this directly from `goldmark-mathjax` but I read too late that it is based on goldmark code for [fenced blocks](https://github.com/yuin/goldmark/blob/master/parser/fcode_block.go) and [code spans](https://github.com/yuin/goldmark/blob/master/parser/code_span.go). Sometime I'll update the parser code to not use the deprecated `util.DedentPosition` and be more similar to those two files.

-   [ ] The old `goldmark-mathjax` actually had some code that compiled latex to SVGs using `pdflatex` and `pdf2svg`. Something interesting would be to use the "tex renderer" just for block latex nodes (rendering inline math nodes to SVGs has the problem of not being able to wrap images along lines, for those falling back to KaTeX or MathJax might be the best option)

## License

MIT
