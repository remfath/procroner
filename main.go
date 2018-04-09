package main

import (
	"os"
	"fmt"
	"time"
	"log"
	"gopkg.in/robfig/cron.v2"
)

func main() {
	Croner := cron.New()
	Croner.AddFunc("* * * * * *", func() {
		f, err := os.Create(fmt.Sprintf("%s/%s.txt", os.TempDir()+"/test", time.Now().Format(time.RFC3339)))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	})
	Croner.Start()
}
