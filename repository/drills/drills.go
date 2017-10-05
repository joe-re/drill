package drillsRepository

import (
	"database/sql"
	"log"
	"time"

	"github.com/joe-re/drill/database"
)

type Drill struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

func CreateTable(db *sql.DB) {
	var q = ""
	q = "CREATE TABLE drills ("
	q += " id INTEGER PRIMARY KEY AUTOINCREMENT"
	q += ", name VARCHAR(255) NOT NULL"
	q += ", created_at TIMESTAMP DEFAULT (DATETIME('now','localtime'))"
	q += ")"
	database.DbExec(db, q)
}

func Create(db *sql.DB, projectName string) {
	q := "INSERT INTO drills"
	q += " (name)"
	q += " VALUES"
	q += " (?)"
	database.DbExec(db, q, projectName)
}

func Destroy(db *sql.DB, drillId int) {
	q := "DELETE FROM drills WHERE id = ?"
	database.DbExec(db, q, drillId)
}

func All(db *sql.DB) []Drill {
	rows := database.Query(db, "SELECT * FROM drills")
	results := []Drill{}
	for rows.Next() {
		results = append(results, toDrill(rows))
	}
	defer rows.Close()
	return results
}

func Find(db *sql.DB, id int) Drill {
	rows := database.Query(db, "SELECT * FROM drills WHERE id = ?", id)
	if !rows.Next() {
		return Drill{}
	}
	drill := toDrill(rows)
	rows.Close()
	return drill
}

func toDrill(rows *sql.Rows) Drill {
	var id int
	var name string
	var createdAt time.Time
	if err := rows.Scan(&id, &name, &createdAt); err != nil {
		log.Fatal(err)
	}
	return Drill{id, name, createdAt}
}
