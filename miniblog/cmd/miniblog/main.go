package main

import (
	_ "go.uber.org/automaxprocs"
	"miniblog/internal/miniblog"
	"os"
)

func main() {
	command := miniblog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

// https://juejin.cn/book/7176608782871429175/section/7179876382296506428
