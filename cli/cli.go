package cli

import (
	"fmt"
	"os"
	"log"
	"github.com/urfave/cli"
	"github.com/remfath/procroner/server"
	"github.com/remfath/procroner/job"
)

func printSign() {
	fmt.Println()
	fmt.Println(`
	______          _____                           
	| ___ \        /  __ \                          
	| |_/ / __ ___ | /  \/_ __ ___  _ __   ___ _ __ 
	|  __/ '__/ _ \| |   | '__/ _ \| '_ \ / _ \ '__|
	| |  | | | (_) | \__/\ | | (_) | | | |  __/ |   
	\_|  |_|  \___/ \____/_|  \___/|_| |_|\___|_|`)
	fmt.Println()
}

func Show() {
	app := cli.NewApp()
	app.Name = "ProCroner"
	app.Usage = "manage cron jobs"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:  "install",
			Usage: "Install the ProCroner service",
			Action: func(c *cli.Context) error {
				msg, err := server.Install()
				fmt.Println(msg)
				return err
			},
		},
		{
			Name:  "start",
			Usage: "Start the ProCroner service",
			Action: func(c *cli.Context) error {
				msg, err := server.Start()
				fmt.Println(msg)
				return err
			},
		},
		{
			Name:  "stop",
			Usage: "Stop the ProCroner service",
			Action: func(c *cli.Context) error {
				msg, err := server.Stop()
				fmt.Println(msg)
				return err
			},
		},
		{
			Name:  "status",
			Usage: "Check the ProCroner status",
			Action: func(c *cli.Context) error {
				msg, err := server.Status()
				fmt.Println(msg)
				return err
			},
		},
		{
			Name:  "reload",
			Usage: "Reload the ProCroner configure",
			Action: func(c *cli.Context) error {
				msg, err := server.Reload()
				fmt.Println(msg)
				return err
			},
		},
		{
			Name:  "restart",
			Usage: "Restart the ProCroner",
			Action: func(c *cli.Context) error {
				msg, err := server.Restart()
				fmt.Println(msg)
				return err
			},
		},
		{
			Name:  "remove",
			Usage: "Remove the ProCroner",
			Action: func(c *cli.Context) error {
				msg, err := server.Remove()
				fmt.Println(msg)
				return err
			},
		},
		{
			Name:  "job",
			Usage: "Manage jobs",
			Subcommands: cli.Commands{
				{
					Name:  "test",
					Usage: "Add a test cron job",
					Action: func(c *cli.Context) error {
						job.AddTestJobs()
						return nil
					},
				},
				{
					Name:  "list",
					Usage: "List jobs",
					Action: func(c *cli.Context) error {
						var listType int
						if c.Bool("all") {
							listType = job.STATUS_ALL
						}
						if c.Bool("enable") {
							listType = job.STATUS_ENABLED
						}
						if c.Bool("disable") {
							listType = job.STATUS_DISABLED
						}
						if c.Bool("running") {
							listType = job.STATUS_RUNNING
						}
						if c.Bool("hanging") {
							listType = job.STATUS_HANGING
						}

						job.ListJobs(listType)
						return nil
					},
					Flags: []cli.Flag{
						cli.BoolFlag{Name: "all", Usage: "List all jobs"},
						cli.BoolFlag{Name: "enable", Usage: "List all enabled jobs"},
						cli.BoolFlag{Name: "disable", Usage: "List all disabled jobs"},
						cli.BoolFlag{Name: "running", Usage: "List all running jobs"},
						cli.BoolFlag{Name: "hanging", Usage: "List all hanging jobs"},
					},
				},
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		action := c.Command.Name
		if action == "" {
			printSign()
			cli.ShowAppHelp(c)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
