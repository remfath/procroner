package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gopkg.in/robfig/cron.v2"
	"github.com/takama/daemon"
	"github.com/remfath/procroner/job"
)

const (
	name        = "Pro Croner"
	description = "it is a cron job service"
)

var stdLog, errLog *log.Logger

type Service struct {
	daemon.Daemon
}

func makeFile() {
	f, err := os.Create(fmt.Sprintf("%s/%s.txt", os.TempDir(), time.Now().Format(time.RFC3339)))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}

func (service *Service) Manage() (string, error) {
	usage := "Usage: procroner install | remove | start | stop | status"
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return usage, nil
		}
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	c := cron.New()
	c.AddFunc("* * * * * *", makeFile)
	c.Start()
	job.ListJobs(c)

	killSignal := <-interrupt
	stdLog.Println("Got signal:", killSignal)
	return "Service exited", nil
}

func init() {
	stdLog = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	errLog = log.New(os.Stderr, "", log.Ldate|log.Ltime)
}

func main() {
	srv, err := daemon.New(name, description)
	if err != nil {
		errLog.Println("Error: ", err)
		os.Exit(1)
	}
	service := &Service{srv}
	status, err := service.Manage()
	if err != nil {
		errLog.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(status)
}
