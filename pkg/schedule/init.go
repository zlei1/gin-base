package schedule

import (
	"log"

	"github.com/robfig/cron/v3"
)

func Init() {
	c := cron.New()

	c.AddFunc("30 0 0 * * *", func() {
		log.Println("timed task start")
	})

	c.Start()
}
