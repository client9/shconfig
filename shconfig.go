package shconfig

import (
	"fmt"
)

func RequireString0(args []string, fn func() error) error {
	if len(args) != 1 {
		return fmt.Errorf("%s: expected no args, got %d", args[0], len(args))
	}
	return fn()
}

func RequireString1(args []string, fn func(string) error) error {
	if len(args) != 2 {
		return fmt.Errorf("%s: expected 1 args, got %d", args[0], len(args))
	}
	return fn(args[1])
}

func RequireString2(args []string, fn func(string, string) error) error {
	if len(args) != 3 {
		return fmt.Errorf("%s: expected 2 args, got %d", args[0], len(args))
	}
	return fn(args[1], args[2])
}

type Dispatcher interface {
	ConfCall(args []string) error
	ConfObject(args []string) (Dispatcher, error)
}

func Parse(root Dispatcher, conf string) error {
	stack := []Dispatcher{root}
	scan := NewParser(conf, "conf")
	for {
		args, err := scan.Next()
		if err != nil {
			return err
		}
		if args == nil {
			break
		}
		if args[len(args)-1] == "{" {
			args = args[:len(args)-1]
			d, err := stack[len(stack)-1].ConfObject(args)
			if err != nil {
				return fmt.Errorf("%s: config err %s", "TODO", err)
			}
			stack = append(stack, d)
			continue
		}
		if len(args) == 1 && args[0] == "}" {
			if len(stack) == 0 {
				return fmt.Errorf("%s: fell off edge", "TODO")
			}
			stack = stack[:len(stack)-1]
			continue
		}
		err = stack[len(stack)-1].ConfCall(args)
		if err != nil {
			return fmt.Errorf("%s config err %s", "TODO", err)
		}
	}

	return nil
}
