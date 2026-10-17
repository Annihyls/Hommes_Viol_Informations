package logger

import (
	"log"
	"os"
)

var Stderr *log.Logger = log.New(os.Stderr, "Error: ", log.LstdFlags | log.Lmicroseconds)
var Stdout *log.Logger = log.New(os.Stdout, "Infos: ", log.LstdFlags | log.Lmicroseconds)
