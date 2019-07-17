package main

import (
	"log"
	"os"
)

func main() {
	log.Printf("Hello world")
	log.SetOutput(os.Stdout)
	log.Printf("This should be to stdout")

}
