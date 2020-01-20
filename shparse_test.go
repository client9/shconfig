package shconfig

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	confs := []string{
		"foo bar",
		"foo 123",
		"foo 123-456-789",
		"foo 123.456",
		"foo-bar",
		"foo; bar",
		"foo \"bar\"",
		`foo "bar"`,
		`foo "b\nar"`,
		`foo "b\x32ar"`,
		"foo bar; hello world",
		"foo { hello world; }",
		`foo 
		bar`,
	}

	for i, c := range confs {
		p := NewParser(c, fmt.Sprintf("conf #%d", i))
		for {
			args, err := p.Next()
			if err != nil {
				t.Errorf("conf %q failed: %v", c, err)
				break
			}
			if args == nil {
				break
			}
			fmt.Printf("ARGS: %v\n", args)
		}
	}
}
