package latex

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type StartEndDelimiters struct {
	Begin, End string
}

type InlineAndBlockDelimiters struct {
	Inline, Block StartEndDelimiters
}

type SourceAndOutputDelimiters struct {
	Source, Output InlineAndBlockDelimiters
}

type Latex struct {
	SourceAndOutputDelimiters
}

// NewLatex creates a goldmark plugin for parsing and rendering inline and block math. By default delimiters are "$" and "$" for inline and "$$" and "$$" for block elements. The default output delimiters are "span.math.inline" for inline math nodes and "div.math.block" for math block nodes.
func NewLatex(opts ...Option) *Latex {
	instance := &Latex{
		SourceAndOutputDelimiters{
			Source: InlineAndBlockDelimiters{
				Inline: StartEndDelimiters{`$`, `$`},
				Block:  StartEndDelimiters{`$$`, `$$`},
			},
			Output: InlineAndBlockDelimiters{
				Inline: StartEndDelimiters{`<span class="math inline">`, `</span>`},
				Block:  StartEndDelimiters{`<div class="math block">`, `</div>`},
			},
		},
	}

	for _, opt := range opts {
		opt.SetOption(instance)
	}

	return instance
}

func (e *Latex) Extend(m goldmark.Markdown) {
	// Parsers
	m.Parser().AddOptions(parser.WithBlockParsers(
		util.Prioritized(NewLatexBlockParser(e.Source.Block), 701),
	))
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(NewInlineLatexParser(e.Source.Inline), 501),
	))

	// Renderers
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewLatexBlockRenderer(e.Output.Block), 501),
		util.Prioritized(NewInlineLatexRenderer(e.Output.Inline), 502),
	))
}

//
// Options
//

type OptionFunc func(*Latex)

func (fn OptionFunc) SetOption(instance *Latex) {
	fn(instance)
}

type Option interface {
	SetOption(*Latex)
}

func WithSourceInlineDelim(start, end string) Option {
	return OptionFunc(func(m *Latex) {
		m.Source.Inline.Begin = start
		m.Source.Inline.End = end
	})
}

func WithSourceBlockDelim(start, end string) Option {
	return OptionFunc(func(m *Latex) {
		m.Source.Block.Begin = start
		m.Source.Block.End = end
	})
}

func WithOutputInlineDelim(start, end string) Option {
	return OptionFunc(func(m *Latex) {
		m.Output.Inline.Begin = start
		m.Output.Inline.End = end
	})
}

func WithOutputBlockDelim(start, end string) Option {
	return OptionFunc(func(m *Latex) {
		m.Output.Block.Begin = start
		m.Output.Block.End = end
	})
}
