package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joe-re/drill/database"

	"github.com/joe-re/drill/repository/drills"
	"github.com/urfave/cli"
)

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func yOrN() bool {
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

func scanText() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return scanner.Text()
}

func addQuestion(db *sql.DB, drill drillsRepository.Drill) error {
	i := 1
	for {
		fmt.Println("Question " + strconv.Itoa(i) + ":")
		question := scanText()
		fmt.Println("Please enter answer:")
		answer := scanText()
		fmt.Println("Do you register it? y/n")
		fmt.Println("question:" + question)
		fmt.Println("answer  :" + answer)
		if yOrN() {
			fmt.Println("question" + strconv.Itoa(i) + "is registered")
		}
		fmt.Println("continue? y/n")
		if yOrN() {
			i = i + 1
		} else {
			return nil

		}
	}
}

func showDrills(db *sql.DB) {
	drills := drillsRepository.All(db)
	for _, drill := range drills {
		fmt.Println("id:" + strconv.Itoa(drill.ID) + ", name:" + drill.Name)
	}
}

func main() {
	app := cli.NewApp()
	first := !exists("./data.db")
	if first {
		os.Create("./data.db")
	}
	db := database.Connect()
	if first {
		drillsRepository.CreateTable(db)
	}
	app.Name = "drill"
	app.Usage = "make an explosive entrance"
	app.Commands = []cli.Command{
		{
			Name:  "create",
			Usage: "create a drill",
			Action: func(c *cli.Context) error {
				fmt.Println("Please enter drill name: ")
				drillsRepository.Create(db, scanText())
				return nil
			},
		},
		{
			Name:  "list",
			Usage: "show drills",
			Action: func(c *cli.Context) error {
				showDrills(db)
				return nil
			},
		},
		{
			Name:  "add",
			Usage: "add qusstion to a drill",
			Action: func(c *cli.Context) error {
				fmt.Println(c.Args().Get(0))
				id, err := strconv.Atoi(c.Args().Get(0))
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				drill := drillsRepository.Find(db, id)
				if drill.ID == 0 {
					fmt.Println("can't find drill")
					os.Exit(0)
				}
				addQuestion(db, drill)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
