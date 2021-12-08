package cli

import (
	"context"
	"fmt"
)

type App struct{}

func (app *App) Main(ctx context.Context, pwd string, args []string) error {
	fmt.Println("app.Main")
	fmt.Println("args:", args)
	fmt.Println("pwd:", pwd)
	
	<-ctx.Done()
	return nil
}