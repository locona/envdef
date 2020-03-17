package main

import (
	"flag"
	"log"

	"github.com/locona/envdef"
)

func main() {
	defaultSource := ".env.sample"
	defaultDist := ".env"

	var source string
	flag.StringVar(&source, "s", defaultSource, "source .env file")
	var dist string
	flag.StringVar(&dist, "d", defaultDist, "distribution .env file")
	var overwrite bool
	flag.BoolVar(&overwrite, "o", false, "Whether to overwrite when the key already exists.")

	flag.Parse()

	result, err := envdef.Diff(source, dist, overwrite)
	if err != nil {
		log.Fatal(err)
	}

	err = result.Write()
	if err != nil {
		log.Fatal(err)
	}
	result.Print()
}
