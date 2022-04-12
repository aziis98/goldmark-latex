package latex

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type InlineLatexRenderer struct {
	Delim StartEndDelimiters
}

func NewInlineLatexRenderer(delim StartEndDelimiters) renderer.NodeRenderer {
	return &InlineLatexRenderer{delim}
}

func (r *InlineLatexRenderer) renderInlineLatex(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		// write begin delim
		_, _ = w.WriteString(r.Delim.Begin)

		// concat all children with a space
		for c := n.FirstChild(); c != nil; c = c.NextSibling() {
			segment := c.(*ast.Text).Segment
			value := segment.Value(source)
			w.Write(value)
		}

		return ast.WalkSkipChildren, nil
	}

	// write end delim
	_, _ = w.WriteString(r.Delim.End)

	return ast.WalkContinue, nil
}

func (r *InlineLatexRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindInlineLatex, r.renderInlineLatex)
}
