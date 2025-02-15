package main

import (
	"flag"

	"github.com/leandrosavn/aclow"
)

func main() {
	src := flag.String("src", "", "source log file")
	dest := flag.String("dest", "", "dest test file")

	flag.Parse()

	aclow.GenerateTests(*src, *dest)
}
