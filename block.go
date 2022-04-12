package latex

import "github.com/yuin/goldmark/ast"

type LatexBlock struct {
	ast.BaseBlock
}

var KindLatexBlock = ast.NewNodeKind("LatexBlock")

func NewLatexBlock() *LatexBlock {
	return &LatexBlock{}
}

func (n *LatexBlock) Dump(source []byte, level int) {
	m := map[string]string{}
	ast.DumpHelper(n, source, level, m, nil)
}

func (n *LatexBlock) Kind() ast.NodeKind {
	return KindLatexBlock
}

func (n *LatexBlock) IsRaw() bool {
	return true
}
