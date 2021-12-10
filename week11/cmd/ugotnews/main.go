package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
	"week11/cmd/ugotnews/cli"
)

func main() {
	//start a base context for the main
	ctx := context.Background()
	//intergrate cancel on interrupt ie. ctrl+c
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	//staring an anynomous go routine to sleep for a few section to allowing the tasks below
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()
	//getting the details of present working directory
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	//constructing app from the cli
	app := &cli.App{}
	//running the Main function
	err = app.Main(ctx, pwd, os.Args[1:])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	<-ctx.Done()
}
