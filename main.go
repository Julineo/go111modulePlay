package main

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	_, err := io.Copy(os.Stdout, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
}
