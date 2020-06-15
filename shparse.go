package shconfig

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

type Parser struct {
	scan *scanner.Scanner
}

func NewParser(s string, name string) *Parser {
	scan := scanner.Scanner{}
	scan.Init(strings.NewReader(s))
	scan.Filename = name
	scan.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanChars | scanner.ScanStrings | scanner.ScanRawStrings | scanner.ScanComments | scanner.SkipComments
	p := &Parser{
		scan: &scan,
	}
	p.RawMode(false)
	return p
}
func (p *Parser) RawMode(on bool) {
	if on {
		// this basically makes everything that is NOT a newline an "identifier"
		// basically copies the line as-in
		p.scan.Whitespace = 0
		p.scan.IsIdentRune = func(ch rune, i int) bool { return ch != '\n' }
		return
	}

	// white is usual stuff but not "\n" since that's end of line
	// and what to deal with it in a different way
	p.scan.Whitespace = 1<<'\t' | 1<<'\r' | 1<<' '
	p.scan.IsIdentRune = func(ch rune, i int) bool {
		if ch < 33 {
			return false
		}

		// cannot be another special token for whitespace
		// or string quoting. (otherwise infinite loops)
		return !(ch == ' ' || ch == '"' || ch == '`')
	}
}

func (p *Parser) Next() ([]string, error) {
	args := []string{}
	for tok := p.scan.Scan(); tok != scanner.EOF; tok = p.scan.Scan() {
		val := p.scan.TokenText()
		vlen := len(val)
		if vlen > 0 && val[0] == '"' && val[vlen-1] == '"' {
			raw, err := strconv.Unquote(val)
			if err != nil {
				return nil, fmt.Errorf("quote error for %s %v", val, err)
			}
			val = raw
		}
		if vlen > 0 && val[0] == '`' && val[vlen-1] == '`' {
			val = val[1 : vlen-1]
		}
		if val == ";" || val == "\n" {
			if len(args) > 0 {
				return args, nil
			}
			continue
		}
		if val == "{" {
			args = append(args, val)
			return args, nil
		}

		// in raw mode, the "}" could be "      }"
		// so we need to trim
		if strings.TrimSpace(val) == "}" {
			args = append(args, strings.TrimSpace(val))
			return args, nil
		}
		args = append(args, val)
	}
	if len(args) > 0 {
		return args, nil
	}
	return nil, nil
}
