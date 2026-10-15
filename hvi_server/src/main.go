package main

import (
        "hvi_server/database"
	"log"
)

func main() {
        db.Init()
        db.Close()
        log.Println("Hello World!")
}
