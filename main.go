package main

import (
	"github.com/sean9999/hermeti"
)

func main() {

	app := new(state)
	cli := hermeti.NewRealCli(app)
	cli.Run()

}
