package main

import (
	"log"
	"os"

	"github.com/sjsanc/encard/internal/cli"
)

func main() {
	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
