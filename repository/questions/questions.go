package questionsRepository

import (
	"database/sql"
	"log"
	"time"

	"github.com/joe-re/drill/database"
)

type Question struct {
	ID        int
	DrillID   int
	Content   string
	Answear   string
	CreatedAt time.Time
}

func CreateTable(db *sql.DB) {
	var q = ""
	q = "CREATE TABLE questions ("
	q += " id INTEGER PRIMARY KEY AUTOINCREMENT"
	q += ", drill_id INTEGER NOT NULL"
	q += ", content TEXT NOT NULL"
	q += ", answear TEXT NOT NULL"
	q += ", created_at TIMESTAMP DEFAULT (DATETIME('now','localtime'))"
	q += ")"
	database.DbExec(db, q)
}

func Create(db *sql.DB, drillId int, content string, answear string) {
	q := "INSERT INTO questions"
	q += " (drill_id, content, answear)"
	q += " VALUES"
	q += " (?, ?, ?)"
	database.DbExec(db, q, drillId, content, answear)
}

func All(db *sql.DB) []Question {
	rows := database.Query(db, "SELECT * FROM questions")
	results := []Question{}
	for rows.Next() {
		results = append(results, toQuestion(rows))
	}
	defer rows.Close()
	return results
}

func Find(db *sql.DB, id int) Question {
	rows := database.Query(db, "SELECT * FROM drills WHERE id = ?", id)
	if !rows.Next() {
		return Question{}
	}
	question := toQuestion(rows)
	rows.Close()
	return question
}

func FindByDrillId(db *sql.DB, drillId int) []Question {
	rows := database.Query(db, "SELECT * FROM questions WHERE drill_id = ?", drillId)
	results := []Question{}
	for rows.Next() {
		results = append(results, toQuestion(rows))
	}
	defer rows.Close()
	return results
}

func toQuestion(rows *sql.Rows) Question {
	var id int
	var drillId int
	var content string
	var answear string
	var createdAt time.Time
	if err := rows.Scan(&id, &drillId, &content, &answear, &createdAt); err != nil {
		log.Fatal(err)
	}
	return Question{id, drillId, content, answear, createdAt}
}
