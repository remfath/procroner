package server

import (
	"log"
	"os"
	"github.com/takama/daemon"
)

const (
	name        = "Pro Croner"
	description = "Pro Croner"
)

type Service struct {
	daemon.Daemon
}

var stdLog, errLog *log.Logger
var service *Service

func init() {
	stdLog = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	errLog = log.New(os.Stderr, "", log.Ldate|log.Ltime)
	ser, err := daemon.New(name, description)
	if err != nil {
		errLog.Println("Server Error: ", err)
		os.Exit(1)
	}
	service = &Service{ser}
}

func Install() (string, error) {
	return service.Install()
}

func Start() (string, error) {
	return service.Start()
}

func Stop() (string, error) {
	return service.Stop()
}

func Remove() (string, error) {
	return service.Remove()
}

func Status() (string, error) {
	return service.Status()
}

func Restart() (string, error) {
	status, err := Stop()
	if err != nil {
		return status, err
	}
	return Start()
}

func Reload() (string, error) {
	return Restart()
}