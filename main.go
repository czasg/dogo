package main

import (
	_ "go.uber.org/automaxprocs"
	"proj/cmd"
	"proj/lifecycle"
)

func main() {
	if err := cmd.App.ExecuteContext(lifecycle.RootContext); err != nil {
		panic(err)
	}
}
