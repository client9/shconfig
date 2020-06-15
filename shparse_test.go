package shconfig

import (
	"fmt"
	"reflect"
	"testing"
	"reflect"
)

func TestParse(t *testing.T) {
<<<<<<< HEAD
	confs := []struct {
		name string
		in   string
		out  []string
=======
	confs := []struct{
		name string
		in string
		out []string
>>>>>>> f0c03c8... support string multi-line litterals
	}{
		{"normal", "foo bar", []string{"foo", "bar"}},
		{"int", "foo 123", []string{"foo", "123"}},
		{"ident-dash", "foo 123-456-789", []string{"foo", "123-456-789"}},
		{"floats", "foo 123.456", []string{"foo", "123.456"}},
		{"dash", "foo-bar", []string{"foo-bar"}},
		{"semi-colon", "foo; bar", []string{"foo"}},
		{"quoted", "foo \"bar\"", []string{"foo", "bar"}},
		{"quoted2", `foo "bar"`, []string{"foo", "bar"}},
<<<<<<< HEAD
		{"quoted-with-newline", `foo "b\nar"`, []string{"foo", "b\nar"}},
=======
		{"quoted-with-newline", `foo "b\nar"`, []string{ "foo", "b\nar" }},
>>>>>>> f0c03c8... support string multi-line litterals
		{"quoted-with-hex", `foo "b\x32ar"`, []string{"foo", "b\x32ar"}},
		{"compound", "foo bar; hello world", []string{"foo", "bar"}},
		{"scope", "foo { hello world; }", []string{"foo", "{"}},
		{"string-litteral", "foo `bar`", []string{"foo", "bar"}},
		{"embedded-newline", "foo\nbar", []string{"foo"}},
		{"string-litteral2", "foo `\nbar\n`", []string{"foo", "\nbar\n"}},
	}

	for i, c := range confs {
		p := NewParser(c.in, fmt.Sprintf("conf #%d", i))
		var args []string
		var err error
		for {
			args, err = p.Next()
			if err != nil {
				t.Errorf("conf %q failed: %v", c.in, err)
				break
			}
			if args == nil {
				break
			}
<<<<<<< HEAD
			if !reflect.DeepEqual(args, c.out) {
=======
			if ! reflect.DeepEqual(args, c.out) {
>>>>>>> f0c03c8... support string multi-line litterals
				t.Errorf("Case %d %s: Expected %v got %v", i, c.name, c.out, args)
			}
			break
		}
	}
}
