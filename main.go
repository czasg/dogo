package main

import (
	_ "go.uber.org/automaxprocs"
	"proj/cmd"
)

func main() {
	panic(cmd.App.Execute())
}
