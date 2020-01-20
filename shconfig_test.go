package shconfig

import (
	"fmt"
	"testing"
)

type dispatcher struct {
	gotHello int
}

func (d *dispatcher) ConfCall(args []string) error {
	switch args[0] {
	case "hello":
		return RequireString0(args, func() error {
			d.gotHello += 1
			return nil
		})
	default:
		return fmt.Errorf("Unknown command %q", args[0])

	}
	return nil
}

func (d *dispatcher) ConfObject(args []string) (Dispatcher, error) {
	return nil, nil
}

func TestString0(t *testing.T) {
	conf := "hello"
	dis := dispatcher{}
	if err := Parse(&dis, conf); err != nil {
		t.Errorf("test0 failed expected nil got %v", err)
	}
	if dis.gotHello != 1 {
		t.Error("test0 call failed")
	}

	conf = "hello world"

	if err := Parse(&dis, conf); err == nil {
		t.Error("test0 failed - expected error got nil")
	}
	/*
		conf = "hello world; hello world"

		if err := Parse(&dis, conf); err != nil {
			t.Errorf("test0 failed expected nil got %v", err)
		}
		if dis.gotHello != 3 {
			t.Error("test0 call failed")
		}
	*/
}
