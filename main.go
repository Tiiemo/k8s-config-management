package main

import (
	"flag"
	"os"
)

var sNewdir = flag.String("NewDir", "", "Name of the directory to create")

func main() {
	flag.Parse()
	os.Mkdir(*sNewdir, os.ModePerm)
}
