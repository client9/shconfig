package shconfig

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
	"unicode"
)

type parser struct {
	scan *scanner.Scanner
}

func NewParser(s string, name string) *parser {
	scan := scanner.Scanner{}
	scan.Init(strings.NewReader(s))
	scan.Whitespace = 1<<'\t' | 1<<'\r' | 1<<' '
	scan.IsIdentRune = func(ch rune, i int) bool {
		return ch == '-' || ch == '.' || (unicode.IsDigit(ch)) || unicode.IsLetter(ch)
	}

	scan.Filename = name
	scan.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanChars | scanner.ScanStrings | scanner.ScanRawStrings | scanner.ScanComments | scanner.SkipComments

	return &parser{
		scan: &scan,
	}
}

func (p *parser) Next() ([]string, error) {
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
		if val == "}" {
			args = append(args, val)
			return args, nil
		}
		args = append(args, val)
	}
	if len(args) > 0 {
		return args, nil
	}
	return nil, nil
}
