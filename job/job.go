package job

import (
	"gopkg.in/robfig/cron.v2"
	"fmt"
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
	Runtime      string
	Enable       bool
	Dependencies []Job
	Creator      User
	Users        []User
}

func ListJobs(c *cron.Cron) {
	jobs := c.Entries()
	for _, job := range jobs {
		fmt.Printf("%d: %s => %v, %v\n", job.ID, job.Schedule, job.Next, job.Prev)
	}
}
