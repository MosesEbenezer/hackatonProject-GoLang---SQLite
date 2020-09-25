package main

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"goFirst/platform/firsttask"
)

func main() {
	db, _ := sql.Open("sqlite3", "./firstTask.db")
	cat := firsttask.CategorySums(db)

	// cat.Add(firsttask.Item{
	// 	CATEGORY: "Helloe", // to adjust better later
	// })

	items := cat.Get()
		fmt.Println(items)
}
