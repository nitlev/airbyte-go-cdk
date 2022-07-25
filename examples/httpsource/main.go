package main

import (
	"log"
	"os"

	"github.com/nitlev/airbyte"
	"github.com/nitlev/airbyte/examples/httpsource/apisource"
)

func main() {
	hsrc := apisource.NewAPISource("https://api.bitstrapped.com")
	runner := airbyte.NewSourceRunner(hsrc, os.Stdout)
	err := runner.Start()
	if err != nil {
		log.Fatal(err)
	}
}
