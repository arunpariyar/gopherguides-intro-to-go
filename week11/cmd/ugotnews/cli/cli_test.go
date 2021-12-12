package cli

import (
	"bytes"
	"context"
	"fmt"
	"testing"
)

type MockCommander struct{
	Name string
	Pwd string
	Args []string
}


func (c *MockCommander) Main(ctx context.Context, pwd string, args []string) error {
	if c == nil {
		return fmt.Errorf("commander is nil")
	}
	c.Pwd = pwd
	c.Args = args
	return nil
}

func Test_App_SubCommand_Routing(t *testing.T){

	fooCmd := &MockCommander{
		Name: "foo",
	}
	
	args := []string{"foo", "bar", "baz"}

	barCmd := &MockCommander{
		Name: "bar",
	}

	table := []struct {
		name string
		c    *MockCommander
		args []string
		err  bool
	}{
		{
			name: "all good",
			c:    fooCmd,
			args: args,
		},
		{
			name: "nil commander",
			args: args,
			err:  true,
		},
		{
			name: "unknown command",
			c:    barCmd,
			args: args,
			err:  true,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {

			app := &App{
				Commands: map[string]Commander{},
			}

			if tt.c != nil {
				app.Commands[tt.c.Name] = tt.c
			}

			err := app.Main(context.Background(), "testdata", tt.args)
			if tt.err {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tt.c.Pwd != "testdata" {
				t.Fatalf("expected %q, got %q", "testdata", tt.c.Pwd)
			}

			exp := tt.args[1:]
			act := tt.c.Args

			if len(exp) != len(act) {
				t.Fatalf("expected %d args, got %d", len(exp), len(act))
			}

			for i, a := range exp {
				if a != act[i] {
					t.Fatalf("expected %q, got %q", a, act[i])
				}
			}

		})
	}

}

func Test_App_Main_Usage(t *testing.T) {
	t.Parallel()

	app := &App{}

	bb := &bytes.Buffer{}

	// assign the bytes buffer to the
	// standard output of the app
	app.Out = bb

	err := app.Main(context.Background(), ".", []string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	act := bb.String()
	exp := "Usage: ugotnews <command> [options][<args>...]\n---------------\n"

	if act != exp {
		t.Fatalf("expected %q, got %q", exp, act)
	}

}


