package cron

import (
	"fmt"
	cronjob "github.com/robfig/cron/v3"
	"topmusicstreaming/hub"
	"topmusicstreaming/utils"
)

func Start() {
	c := cronjob.New()
	c.AddFunc(setCronSpec(utils.EuropeLocation, 15, 30), func() { hub.Launch(utils.US) })
	c.AddFunc(setCronSpec(utils.EuropeLocation, 16, 30), func() { hub.Launch(utils.FR) })
	c.AddFunc(setCronSpec(utils.EuropeLocation, 17, 30), func() { hub.Launch(utils.DE) })
	c.AddFunc(setCronSpec(utils.EuropeLocation, 18, 30), func() { hub.Launch(utils.ES) })
	c.AddFunc(setCronSpec(utils.EuropeLocation, 19, 30), func() { hub.Launch(utils.PT) })
	c.AddFunc(setCronSpec(utils.EuropeLocation, 20, 30), func() { hub.Launch(utils.IT) })
	c.Start()
}

func setCronSpec(location string, hour, minute int) string {
	return fmt.Sprintf("CRON_TZ=%s %d %d * * *", location, minute, hour)
}
