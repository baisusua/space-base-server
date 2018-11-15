package main

import (
	db "space-base-server/database"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":8000")
}
