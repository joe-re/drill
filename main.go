package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "drill"
	app.Usage = "make an explosive entrance"
	app.Commands = []cli.Command{
		{
			Name:  "create",
			Usage: "create a drill",
			Action: func(c *cli.Context) error {
				fmt.Println("Please enter drill name: ")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				if err := scanner.Err(); err != nil {
					return err
				}
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
					for scanner.Scan() {
						if scanner.Text() == "y" {
							fmt.Println("question" + strconv.Itoa(i) + "is registered")
							break
						} else if scanner.Text() == "n" {
							break
						}
					}
					if err := scanner.Err(); err != nil {
						return err
					}
					fmt.Println("continue? y/n")
					for scanner.Scan() {
						if scanner.Text() == "y" {
							i = i + 1
							break
						} else if scanner.Text() == "n" {
							return nil
						}
					}
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
