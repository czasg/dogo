package main

import (
	_ "go.uber.org/automaxprocs"
	"proj/cmd"
	"proj/lifecycle"
)

func main() {
	panic(cmd.App.ExecuteContext(lifecycle.RootContext))
}
