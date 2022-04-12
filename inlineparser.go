package latex

import (
	"bytes"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var inlineLatexInfoKey = parser.NewContextKey()

type inlineLatexData struct{ indent int }

type inlineLatexParser struct {
	Delim StartEndDelimiters
}

func NewInlineLatexParser(delim StartEndDelimiters) parser.InlineParser {
	return &inlineLatexParser{delim}
}

func (s *inlineLatexParser) Trigger() []byte {
	return []byte{s.Delim.Begin[0]}
}

func (s *inlineLatexParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	line, _ := block.PeekLine()

	// check line starts with the starting delimiter
	if !bytes.HasPrefix(line, []byte(s.Delim.Begin)) {
		return nil
	}

	// advance begining of inline math piece
	block.Advance(len(s.Delim.Begin))

	node := NewInlineLatex()

outer:
	for {
		line, segment := block.PeekLine()

		// search "s.Delim.End"
		for i := 0; i < len(line); i++ {
			if bytes.HasPrefix(line[i:], []byte(s.Delim.End)) {
				// found end of inline math
				segment := segment.WithStop(segment.Start + i)
				if !segment.IsEmpty() {
					node.AppendChild(node, ast.NewRawTextSegment(segment))
				}
				block.Advance(i + len(s.Delim.End))
				break outer
			}
		}

		// still inside math but line has ended
		if !util.IsBlank(line) {
			node.AppendChild(node, ast.NewRawTextSegment(segment))
		}

		block.AdvanceLine()
	}

	// 	for {
	// 		line, segment := block.PeekLine()
	// 		if line == nil {
	// 			block.SetPosition(l, pos)
	// 			return ast.NewTextSegment(startSegment.WithStop(startSegment.Start + opener))
	// 		}
	// 		for i := 0; i < len(line); i++ {
	// 			c := line[i]
	// 			if c == '$' {
	// 				oldi := i
	// 				for ; i < len(line) && line[i] == '$'; i++ {
	// 				}
	// 				closure := i - oldi
	// 				if closure == opener && (i+1 >= len(line) || line[i+1] != '$') {
	// 					segment := segment.WithStop(segment.Start + i - closure)
	// 					if !segment.IsEmpty() {
	// 						node.AppendChild(node, ast.NewRawTextSegment(segment))
	// 					}
	// 					block.Advance(i)
	// 					goto end
	// 				}
	// 			}
	// 		}
	// 		if !util.IsBlank(line) {
	// 			node.AppendChild(node, ast.NewRawTextSegment(segment))
	// 		}
	// 		block.AdvanceLine()
	// 	}
	// end:

	if !node.IsBlank(block.Source()) {
		// trim first halfspace and last halfspace
		shouldTrimmed := true

		segment := node.FirstChild().(*ast.Text).Segment
		if !(!segment.IsEmpty() && block.Source()[segment.Start] == ' ') {
			shouldTrimmed = false
		}

		segment = node.LastChild().(*ast.Text).Segment
		if !(!segment.IsEmpty() && block.Source()[segment.Stop-1] == ' ') {
			shouldTrimmed = false
		}

		if shouldTrimmed {
			textBegin := node.FirstChild().(*ast.Text)
			beginSegment := textBegin.Segment
			textBegin.Segment = beginSegment.WithStart(beginSegment.Start + 1)

			textEnd := node.LastChild().(*ast.Text)
			endSegment := node.LastChild().(*ast.Text).Segment
			textEnd.Segment = endSegment.WithStop(endSegment.Stop - 1)
		}
	}

	return node
}
