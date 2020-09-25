package firsttask

import (
	"database/sql"
)


// Categories is ...
type Categories struct {
	DB *sql.DB
}


// Get is ...
func (categories *Categories) Get() []Item {
	items := []Item{}
	rows, _ := categories.DB.Query(`
		SELECT * FROM firstTask
	`)

	var id int
	var category string
	var amount float64
	var parentId int

	for rows.Next() {
			rows.Scan(&id, &category, &amount, &parentId)

			payload := Item {
				ID: id, 
				CATEGORY: category, 
				AMOUNT: amount, 
				PARENTID: parentId,
			}

			items = append(items, payload)
	}
	return items
}


// Add is ...
func (categories *Categories) Add(item Item) {
	//// INSERT INTO DB
	stmt, _ := categories.DB.Prepare(`
			INSERT INTO firstTask ("id", "category", "amount", "parentId") 
														VALUES ('1', 'cat1', '2000', '')
														VALUES ('2', 'cat2', '1000', '1')
														VALUES ('3', 'cat3', '3000', '1')
														VALUES ('4', 'cat4', '1500', '2')
														VALUES ('5', 'cat5', '1100', '3')
														VALUES ('6', 'cat6', '900', '2')
	`)
	stmt.Exec()
}

// CategorySums is ...
func CategorySums(db *sql.DB) *Categories {

	///////////// CREATE DB ///////////////////

	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "firstTask" (
			"id"	INTEGER NOT NULL UNIQUE,
			"category"	TEXT NOT NULL,
			"amount"	NUMERIC NOT NULL,
			"parentId"	INTEGER,
			PRIMARY KEY("id" AUTOINCREMENT)
		);
	`)
	stmt.Exec()

	return &Categories { // returns a Categories referrence
		DB: db,
	}
}