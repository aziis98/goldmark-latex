package latex

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

type InlineLatex struct {
	ast.BaseInline
}

func (n *InlineLatex) Inline() {}

func (n *InlineLatex) IsBlank(source []byte) bool {
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		text := c.(*ast.Text).Segment
		if !util.IsBlank(text.Value(source)) {
			return false
		}
	}
	return true
}

func (n *InlineLatex) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

var KindInlineLatex = ast.NewNodeKind("InlineLatex")

func (n *InlineLatex) Kind() ast.NodeKind {
	return KindInlineLatex
}

func NewInlineLatex() *InlineLatex {
	return &InlineLatex{
		BaseInline: ast.BaseInline{},
	}
}
