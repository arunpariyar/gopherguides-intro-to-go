package cli

import (
	"context"
	"fmt"
	"io"
	"os"
)

func (oi IO) Stdin() io.Reader {
	if oi.In == nil {
		return os.Stdin
	}
	return oi.In
}

type App struct {
	IO
}

func (app *App) Main(ctx context.Context, pwd string, args []string) error {

	if app == nil {
		return fmt.Errorf("app is nil")
	}

	if len(args) == 0 {
		return app.Usage(app.Stdout())
	}

	fmt.Println("app.Main")
	fmt.Println("args:", args)
	fmt.Println("pwd:", pwd)

	<-ctx.Done()
	return nil
}

func (app *App) Usage(w io.Writer) error {
	fmt.Fprintln(w, "Usage: ugotnews <command> [options][<args>...]")
	fmt.Fprintln(w, "---------------")

	//TODO: pring sub-commandsß
	return nil
}
