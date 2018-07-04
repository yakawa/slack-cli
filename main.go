package main

import (
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/yakawa/slack-cli/cmd"
	"os"
)

func main() {
	cmd.Execute(
		rwi.New(
			rwi.WithReader(os.Stdin),
			rwi.WithWriter(os.Stdout),
			rwi.WithErrorWriter(os.Stderr),
		),
		os.Args[1:],
	).Exit()
}
