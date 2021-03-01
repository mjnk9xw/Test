package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Create new cron")
	c := cron.New()
	c.AddFunc("@every 1h0m1s", func() { fmt.Println("[Job 1]Every minute job") })

	c.Start()
	time.Sleep(10 * time.Hour)

	c.Stop() // Stop the scheduler (does not stop any jobs already running).

}
