package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cyingfan/gvmgo/cache"
	"github.com/cyingfan/gvmgo/parser"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "GVMGO"
	app.Usage = "Alternative client for sdkman (http://sdkman.io)"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name:  "broadcast",
			Usage: "Show broadcast message",
			Action: func(c *cli.Context) error {
				fmt.Println(cache.GetBroadcast())
				return nil
			},
		},
		{
			Name:  "list",
			Usage: "Lists all candidates",
			Action: func(c *cli.Context) error {
				fmt.Printf("%-20s%-30s%s\n", "Package Name", "Candidate Name", "Latest Version")
				fmt.Printf("----------------------------------------------------------------\n")
				for _, candidate := range parser.ParseCandidateList(cache.GetCandidates()) {
					fmt.Printf("%-20s%-30s%s\n", candidate.ShortName, candidate.Name, candidate.Version)
				}
				return nil
			},
		},
		{
			Name:  "status",
			Usage: "Show status",
			Action: func(c *cli.Context) error {
				lastUpdate := cache.GetLastUpdateISO()
				fmt.Printf("Last update: %s\n", string(lastUpdate))
				return nil
			},
		},
		{
			Name:  "update",
			Usage: "Update",
			Action: func(c *cli.Context) error {
				cache.Update()
				cache.SetLastUpdate(time.Now())
				fmt.Println("Update success")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
