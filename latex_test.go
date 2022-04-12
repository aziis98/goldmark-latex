package latex_test

import (
	"bytes"
	"strings"
	"testing"

	latex "github.com/aziis98/goldmark-latex"
	"github.com/stretchr/testify/assert"
	"github.com/yuin/goldmark"
)

type testGoldmarkLatex struct {
	config *latex.Latex

	label, source, output string
}

func dedent(level int, s string) string {
	lines := strings.Split(s, "\n")

	// remove first line if blank
	if strings.TrimSpace(lines[0]) == "" {
		lines = lines[1:]
	}

	// remove last line if blank
	if strings.TrimSpace(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-1]
	}

	// remove the first {level} runes from every line
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) >= level {
			lines[i] = lines[i][level:]
		}
	}

	return strings.Join(lines, "\n")
}

func TestLatex(t *testing.T) {
	tests := []testGoldmarkLatex{
		{
			label: "Inline",
			source: dedent(4, `
				lorem $1 + 2$ ipsum
			`),
			output: dedent(4, `
				<p>lorem <div class="inline">1 + 2</div> ipsum</p>
			`),
		},
		{
			label: "InlineCustomSourceDelim",
			config: latex.NewLatex(
				latex.WithSourceInlineDelim(`\(`, `\)`),
			),
			source: dedent(4, `
				lorem \(1 + 2\) ipsum
			`),
			output: dedent(4, `
				<p>lorem <div class="inline">1 + 2</div> ipsum</p>
			`),
		},
		{
			label: "InlineCustomSourceOutputDelim",
			config: latex.NewLatex(
				latex.WithSourceInlineDelim(`<<`, `>>`),
				latex.WithOutputInlineDelim(`<custom-latex>`, `</custom-latex>`),
			),
			source: dedent(4, `
				lorem <<1 + 2>> ipsum
			`),
			output: dedent(4, `
				<p>lorem <custom-latex>1 + 2</custom-latex> ipsum</p>
			`),
		},
		{
			label: "LongInline",
			source: dedent(4, `
				lorem $1+2+3
				+4+5+6$ ipsum.
			`),
			output: dedent(4, `
				<p>lorem <div class="inline">1+2+3
				+4+5+6</div> ipsum.</p>
			`),
		},
		{
			label: "Block",
			source: dedent(4, `
				lorem
				
				$$
				1+2
				$$
				
				ipsum
			`),
			output: dedent(4, `
				<p>lorem</p>
				<p><div class="block">1+2
				</div></p>
				<p>ipsum</p>
			`),
		},
		{
			label: "Complex",
			source: dedent(4, `
				- lorem ipsum

				  $$
				  1 + 2 + 3
				  $$
			`),
			output: dedent(4, `
				<ul>
				<li>
				<p>lorem ipsum</p>
				<p><div class="block">1 + 2 + 3
				</div></p>
				</li>
				</ul>
			`),
		},
		{
			// I still don't really understand what goldmark "padding" is but at least this test case works as intended
			label: "Padding?",
			source: dedent(4, `
				- lorem ipsum

				  $$
				    1 + 2
				      + 3
				  $$
			`),
			output: dedent(4, `
				<ul>
				<li>
				<p>lorem ipsum</p>
				<p><div class="block">  1 + 2
				    + 3
				</div></p>
				</li>
				</ul>
			`),
		},
		{
			// ok even this works as intended
			label: "Padding??",
			source: dedent(4, `
				- lorem ipsum

				    $$
				  1 + 2
				  $$
			`),
			output: dedent(4, `
				<ul>
				<li>
				<p>lorem ipsum</p>
				<p><div class="block">1 + 2
				</div></p>
				</li>
				</ul>
			`),
		},
	}

	for _, test := range tests {
		t.Run(test.label, func(t *testing.T) {
			latexCfg := latex.DefaultConfig
			if test.config != nil {
				latexCfg = test.config
			}

			md := goldmark.New(
				goldmark.WithExtensions(latexCfg),
			)

			markdown := goldmark.New(
				goldmark.WithExtensions(
					latex.NewLatex(
						latex.WithSourceInlineDelim()
					),
				),
			)

			t.Logf("\n=== Source ===\n%s\n==============", test.source)

			var buf bytes.Buffer
			if err := md.Convert([]byte(test.source), &buf); err != nil {
				t.Fatal(err)
			}

			t.Logf("\n=== Output ===\n%s\n==============", buf.String())

			assert.Equal(t, test.output, strings.TrimSpace(buf.String()))
		})
	}
}
