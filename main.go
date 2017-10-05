package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli"
)

func DbExec(db *sql.DB, q string) {
	var _, err = db.Exec(q)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func connectToDataBase() *sql.DB {
	first := !Exists("./data.db")
	if first {
		fmt.Printf("first")
	} else {
		fmt.Printf("false")
	}
	if first {
		os.Create("./data.db")
	}
	var db *sql.DB
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if first {
		createProjectsTable(db)
	}
	return db
}

func createProjectsTable(db *sql.DB) {
	var q = ""
	q = "CREATE TABLE projects ("
	q += " id INTEGER PRIMARY KEY AUTOINCREMENT"
	q += ", name VARCHAR(255) NOT NULL"
	q += ", created_at TIMESTAMP DEFAULT (DATETIME('now','localtime'))"
	q += ")"
	DbExec(db, q)
}

func YorN() bool {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "y" {
			return true
		} else if scanner.Text() == "n" {
			return false
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return false
}

func AddQuestion(db *sql.DB, drillId int) error {
	scanner := bufio.NewScanner(os.Stdin)
	i := 1
	for {
		fmt.Println("Question " + strconv.Itoa(i) + ":")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			return err
		}
		question := scanner.Text()
		fmt.Println("Please enter answer:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			return err
		}
		answer := scanner.Text()
		fmt.Println("Do you register it?")
		fmt.Println("question:" + question)
		fmt.Println("answer  :" + answer)
		fmt.Println("y/n")
		if YorN() {
			fmt.Println("question" + strconv.Itoa(i) + "is registered")
		}
		fmt.Println("continue? y/n")
		if YorN() {
			i = i + 1
		} else {
			return nil

		}
	}
}

func createProject(db *sql.DB, projectName string) {
	q := "INSERT INTO projects "
	q += " (name)"
	q += " VALUES"
	q += " (?)"
	db.Exec(q, projectName)
}

func main() {
	app := cli.NewApp()
	app.Name = "drill"
	app.Usage = "make an explosive entrance"
	app.Commands = []cli.Command{
		{
			Name:  "create",
			Usage: "create a drill",
			Action: func(c *cli.Context) error {
				db := connectToDataBase()
				fmt.Println("Please enter drill name: ")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				if err := scanner.Err(); err != nil {
					return err
				}
				createProject(db, scanner.Text())
				return nil
			},
		},
	}

	app.Run(os.Args)
}
