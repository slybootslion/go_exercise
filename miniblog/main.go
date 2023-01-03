package main

import (
	"miniblog/internal/miniblog"
	"os"

	_ "go.uber.org/automaxprocs"
)

func main() {
	command := miniblog.NewMinBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
