package cli

import (
	"bytes"
	"context"
	"testing"
)

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
