package logging

import (
	"log"
	"os"
)

// Debugf prints logging information to standard out.
func Debugf(format string, a ...interface{}) {
	log.SetOutput(os.Stdout)
	log.Printf(format, a...)
}
