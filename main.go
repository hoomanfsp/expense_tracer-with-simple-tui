package main

import (
	"et_sui/database"
	"et_sui/ui"
	"log"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("error, using database !", err)
	}
	ui.Start(db)
}
