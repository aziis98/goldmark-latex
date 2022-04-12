package latex

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type LatexBlockRenderer struct {
	Delim StartEndDelimiters
}

func NewLatexBlockRenderer(delim StartEndDelimiters) renderer.NodeRenderer {
	return &LatexBlockRenderer{delim}
}

func (r *LatexBlockRenderer) renderLatexBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*LatexBlock)
	if entering {
		_, _ = w.WriteString("<p>" + r.Delim.Begin)
		l := n.Lines().Len()
		for i := 0; i < l; i++ {
			line := n.Lines().At(i)
			w.Write(line.Value(source))
		}
	} else {
		_, _ = w.WriteString(r.Delim.End + "</p>\n")
	}
	return ast.WalkContinue, nil
}

func (r *LatexBlockRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindLatexBlock, r.renderLatexBlock)
}
