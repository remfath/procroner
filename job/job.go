package job

import (
	"fmt"
	"gopkg.in/robfig/cron.v2"
	"os"
	"time"
	"log"
)

const (
	CATEGORY_NORMAL = iota
	CATEGORY_EMAIL
)

const (
	STATUS_ALL      = iota
	STATUS_ENABLED
	STATUS_DISABLED
	STATUS_RUNNING
	STATUS_HANGING
	STATUS_CANCELD
	STATUS_FINISHED
	STATUS_SUCCESS
	STATUS_FAILED
)

type Phone int
type Email string

type User struct {
	Id    int
	Name  string
	Phone Phone
	Email Email
}

type Job struct {
	Id           int
	Name         string
	Desc         string
	Command      string
	Runtime      string
	Enable       bool
	Dependencies []Job
	Creator      User
	Users        []User
	OutputText   string
	ErrorText    string
	Category     int
	EnableEmail  bool
	EnableSMS    bool
	Status       int
}

func AddTestJobs(c *cron.Cron) {
	c.AddFunc("* * * * * *", func() {
		f, err := os.Create(fmt.Sprintf("%s/%s.txt", os.TempDir()+"/test", time.Now().Format(time.RFC3339)))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	})
}

func ListJobs(c *cron.Cron, listType int) {
	fmt.Printf(">>> %d \n", listType)
	jobs := c.Entries()
	for _, job := range jobs {
		fmt.Printf("%d: %s => %v, %v\n", job.ID, job.Schedule, job.Next, job.Prev)
	}
}
