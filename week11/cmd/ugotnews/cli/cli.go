package cli

import (
	"context"
	"fmt"
	"io"
)

type App struct {
	IO //embedded IO to allow IO instead of std io
	Commands map[string]Commander
}

func (app *App) Main(ctx context.Context, pwd string, args []string) error {

	if app == nil {
		return fmt.Errorf("app is nil")
	}

	if len(args) == 0 {
		return app.Usage(app.Stdout())
	}

	cmd, ok := app.Commands[args[0]]
	if !ok {
		return fmt.Errorf("command %q not found", args[0])
	}

	return cmd.Main(ctx, pwd, args[1:])

	// fmt.Println("app.Main")
	// fmt.Println("args:", args)
	// fmt.Println("pwd:", pwd)

	// <-ctx.Done()
	// return nil
}

func (app *App) Usage(w io.Writer) error {
	fmt.Fprintln(w, "Usage: ugotnews <command> [options][<args>...]")
	fmt.Fprintln(w, "---------------")

	//TODO: pring sub-commandsß
	return nil
}
