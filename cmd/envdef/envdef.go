package main

import (
	"flag"
	"log"

	"github.com/locona/envdef"
)

var (
	defaultSource = ".env.sample"
	defaultDist   = ".env"
)

func main() {
	var source string
	flag.StringVar(&source, "s", defaultSource, "source .env file")
	var dist string
	flag.StringVar(&dist, "d", defaultDist, "distribution .env file")
	flag.Parse()

	result, err := envdef.Diff(source, dist)
	if err != nil {
		log.Fatal(err)
	}

	result.Write()
	result.Print()
}
