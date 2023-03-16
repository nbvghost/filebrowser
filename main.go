package main

import (
	"flag"
	"github.com/filebrowser/filebrowser/v2/cmd"
)

func main() {
	flag.Parse()
	cmd.Execute()
}
