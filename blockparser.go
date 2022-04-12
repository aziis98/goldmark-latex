package latex

import (
	"bytes"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var latexBlockInfoKey = parser.NewContextKey()

type latexBlockData struct{ indent int }

type latexBlockParser struct {
	Delim StartEndDelimiters
}

func NewLatexBlockParser(delim StartEndDelimiters) parser.BlockParser {
	return &latexBlockParser{delim}
}

func (b *latexBlockParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	line, _ := reader.PeekLine()

	indent := pc.BlockOffset()
	if indent == -1 {
		return nil, parser.NoChildren
	}

	if !bytes.HasPrefix(line[indent:], []byte(b.Delim.Begin)) {
		return nil, parser.NoChildren
	}

	pc.Set(latexBlockInfoKey, &latexBlockData{indent})

	return NewLatexBlock(), parser.NoChildren
}

func (b *latexBlockParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	line, segment := reader.PeekLine()

	data := pc.Get(latexBlockInfoKey).(*latexBlockData)

	w, pos := util.IndentWidth(line, 0)
	if w < 4 {
		if bytes.HasPrefix(line[pos:], []byte(b.Delim.End)) &&
			util.IsBlank(line[pos+len(b.Delim.End):]) {
			reader.Advance(segment.Stop - segment.Start - segment.Padding)
			return parser.Close
		}
	}

	pos, padding := util.DedentPosition(line, 0, data.indent)
	seg := text.NewSegmentPadding(segment.Start+pos, segment.Stop, padding)

	node.Lines().Append(seg)
	reader.AdvanceAndSetPadding(segment.Stop-segment.Start-pos-1, padding)
	return parser.Continue | parser.NoChildren
}

func (b *latexBlockParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
	pc.Set(latexBlockInfoKey, nil)
}

func (b *latexBlockParser) CanInterruptParagraph() bool {
	return true
}

func (b *latexBlockParser) CanAcceptIndentedLine() bool {
	return false
}

func (b *latexBlockParser) Trigger() []byte {
	return nil
}
