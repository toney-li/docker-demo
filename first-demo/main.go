package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/labstack/echo"
	cli "github.com/urfave/cli/v2"
)

func main() {
	cliCmd(os.Args)
}

func handlerindex(c echo.Context) error {
	fmt.Println("Hello world handlerindex")
	return c.JSON(http.StatusOK, `{"hello":"world"}`)
}

func cliCmd(args []string) {
	app := &cli.App{
		Name:     "Sub_Cmd",
		Version:  "v1.0.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Toney",
				Email: "lyxmq@126.com",
			},
		},
		Copyright: "(c) 2019 Toney",
		Usage:     "Cli sub commands test",
		UsageText: "sub_cmd [global options] command [command options] [arguments...]",
		ArgsUsage: "[args and such]",

		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					fmt.Println("added task: ", c.Args().First())
					return nil
				},
			},
			{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "http listen port",
				Action: func(c *cli.Context) error {
					// fmt.Println("added task: ", c.Args().First())
					port := c.Args().First()
					pattern := `^(\d+)$`
					result, _ := regexp.MatchString(pattern, port)
					if !result {
						fmt.Printf("port must is number %s", port)
					} else {
						httpBoostrap(port)
					}
					return nil
				},
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action: func(c *cli.Context) error {
					fmt.Println("completed task: ", c.Args().First())
					return nil
				},
			},
			{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "options for task templates",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a new template",
						Action: func(c *cli.Context) error {
							fmt.Println("new task template: ", c.Args().First())
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(c *cli.Context) error {
							fmt.Println("removed task template: ", c.Args().First())
							return nil
						},
					},
				},
			},
		},
	}
	err := app.Run(args)
	if err != nil {
		log.Fatal(err)
	}
}

func httpBoostrap(port string) {
	e := echo.New()
	e.GET("/", handlerindex)
	fmt.Println("starting echo")
	err := e.Start(":" + port)
	if err != nil {
		fmt.Println("echo", err)
	}
}
