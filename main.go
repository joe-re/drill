package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joe-re/drill/database"

	"github.com/joe-re/drill/repository/drills"
	"github.com/joe-re/drill/repository/questions"
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
	i := len(questionsRepository.FindByDrillId(db, drill.ID)) + 1
	for {
		fmt.Println("Question " + strconv.Itoa(i) + ":")
		question := scanText()
		fmt.Println("Please enter answer:")
		answer := scanText()
		fmt.Println("Do you register it? y/n")
		fmt.Println("question:" + question)
		fmt.Println("answer  :" + answer)
		if yOrN() {
			questionsRepository.Create(db, drill.ID, question, answer)
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

func showQuestions(db *sql.DB, drillId int) {
	questions := questionsRepository.FindByDrillId(db, drillId)
	for _, question := range questions {
		fmt.Println("id:" + strconv.Itoa(question.ID) + ", question:" + question.Content + ", answear:" + question.Answear)
	}
}

func makeQuestions(db *sql.DB, drillId int) {
	questions := questionsRepository.FindByDrillId(db, drillId)
	first := true
	for _, question := range questions {
		if !first {
			fmt.Println("Next")
			scanText()
		}
		fmt.Println("Question" + strconv.Itoa(question.ID))
		fmt.Println(question.Content)
		fmt.Println("Enter your answear:")
		scanText()
		fmt.Println("Answear: " + question.Answear)
		first = false
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
		questionsRepository.CreateTable(db)
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
		{
			Name:  "show",
			Usage: "show questions in a drill",
			Action: func(c *cli.Context) error {
				fmt.Println(c.Args().Get(0))
				id, err := strconv.Atoi(c.Args().Get(0))
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				showQuestions(db, id)
				return nil
			},
		},
		{
			Name:  "question",
			Usage: "make questions",
			Action: func(c *cli.Context) error {
				fmt.Println(c.Args().Get(0))
				id, err := strconv.Atoi(c.Args().Get(0))
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				makeQuestions(db, id)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
