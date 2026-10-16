package main

import (
        "hvi_server/database"
	"log"
)

func main() {
        database.Init()
        database.Close()
        log.Println("Hello World!")
}
