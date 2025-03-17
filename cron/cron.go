package cron

import (
	"fmt"
	cronjob "github.com/robfig/cron/v3"
	"topmusicstreaming/hub"
	"topmusicstreaming/utils"
)

func Start() {
	c := cronjob.New()
	c.AddFunc(setCronSpec(utils.EuropeLocation, 10, 00), func() { hub.Launch(utils.WW) })
	c.AddFunc(setCronSpec(utils.EuropeLocation, 11, 00), func() { hub.Launch(utils.US) })
	c.AddFunc(setCronSpec(utils.EuropeLocation, 12, 00), func() { hub.Launch(utils.FR) })
	c.AddFunc(setCronSpec(utils.EuropeLocation, 13, 00), func() { hub.Launch(utils.UK) })
	c.AddFunc(setCronSpec(utils.EuropeLocation, 14, 00), func() { hub.Launch(utils.JP) })
	c.AddFunc(setCronSpec(utils.EuropeLocation, 15, 00), func() { hub.Launch(utils.KR) })
	c.AddFunc(setCronSpec(utils.EuropeLocation, 16, 00), func() { hub.Launch(utils.TR) })
	c.AddFunc(setCronSpec(utils.EuropeLocation, 17, 00), func() { hub.Launch(utils.DE) })
	c.AddFunc(setCronSpec(utils.EuropeLocation, 18, 00), func() { hub.Launch(utils.ES) })
	c.AddFunc(setCronSpec(utils.EuropeLocation, 19, 00), func() { hub.Launch(utils.PT) })
	c.AddFunc(setCronSpec(utils.EuropeLocation, 20, 00), func() { hub.Launch(utils.IT) })	
	c.Start()
}

func setCronSpec(location string, hour, minute int) string {
	return fmt.Sprintf("CRON_TZ=%s %d %d * * *", location, minute, hour)
}
