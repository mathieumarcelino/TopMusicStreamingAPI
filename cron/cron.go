package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"topmusicstreaming/hub"
	"topmusicstreaming/utils"
)

func Start() {
	cUS := cron.New()
	cUS.AddFunc(setCronSpec(utils.EuropeLocation, 15, 30), func() { hub.Hub(utils.US) })
	cUS.Start()

	cFR := cron.New()
	cFR.AddFunc(setCronSpec(utils.EuropeLocation, 16, 30), func() { hub.Hub(utils.FR)})
	cFR.Start()

	cDE := cron.New()
	cDE.AddFunc(setCronSpec(utils.EuropeLocation, 17, 30), func() { hub.Hub(utils.DE) })
	cDE.Start()

	cES := cron.New()
	cES.AddFunc(setCronSpec(utils.EuropeLocation, 18, 30), func() { hub.Hub(utils.ES)})
	cES.Start()

	cPT := cron.New()
	cPT.AddFunc(setCronSpec(utils.EuropeLocation, 19, 30), func() { hub.Hub(utils.PT) })
	cPT.Start()

	cIT := cron.New()
	cIT.AddFunc(setCronSpec(utils.EuropeLocation,20, 30), func() { hub.Hub(utils.IT) })
	cIT.Start()
}

func setCronSpec(location string, hour, minute int) string {
	return fmt.Sprintf("CRON_TZ=%s %d %d * * *", location, minute, hour)
}

