package main

import (
	_ "go.uber.org/automaxprocs"
	"miniblog/internal/miniblog"
	"os"
)

func main() {
	command := miniblog.NewMinBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
