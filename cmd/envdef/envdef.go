package main

import (
	"flag"
	"log"

	"github.com/locona/envdef"
)

func main() {
	defaultSource := ".env.sample"
	defaultDist := ".env"

	source := ""
	dist := ""
	overwrite := false

	flag.StringVar(&source, "s", defaultSource, "source .env file")
	flag.StringVar(&dist, "d", defaultDist, "distribution .env file")
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
