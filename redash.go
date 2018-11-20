package main

import (
	"log"
	"os"

	"github.com/koooge/redash-cli/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "--doc" {
		err := doc.GenMarkdownTree(cmd.RootCmd(), "./doc/")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		cmd.Execute()
	}
}
