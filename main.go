package main

import (
	"log"
	"os"

	"github.com/cbrand/ir-remote-backend/cli"
)

func main() {
	appCli := cli.GetApp()
	err := appCli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
