package main

import (
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	log.Printf("This should be to standard out")
	log.SetOutput(os.Stderr)
	log.Printf("This should be to standard error")
}
