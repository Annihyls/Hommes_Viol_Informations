/*
Package logger allow to redirect the message in the correct stream:

	stdout for regular logs
	stderr for errors.
*/
package logger

import (
	"log"
	"os"
)

var Stderr *log.Logger = log.New(os.Stderr, "Error: ", log.LstdFlags|log.Lmicroseconds)
var Stdout *log.Logger = log.New(os.Stdout, "Info:  ", log.LstdFlags|log.Lmicroseconds)
