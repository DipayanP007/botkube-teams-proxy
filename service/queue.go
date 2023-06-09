package service

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"time"

	"github.com/rs/zerolog/log"

	"github.com/DipayanP007/botkube-teams-proxy/initializers"
	"github.com/DipayanP007/botkube-teams-proxy/models"
	"github.com/DipayanP007/botkube-teams-proxy/utils"
)

func Enqueue(alert models.Alert) {
	log.Debug().Msg("Checking if Basic Monitoring is enabled")
	if basic_monitoring := strings.ToLower(os.Getenv("BASIC_MONITORING")); basic_monitoring == "true" {
		log.Info().Msg("BASIC_MONITORING is enabled. Filtering only pod events")
		if strings.ToLower(alert.Data.Kind) != "pod" {
			log.Debug().Msg("Dropping event as kind is not pod")
			return
		}
	} else {
		log.Info().Msg("Using Verbose monitoring to capture all events.")
	}
	log.Info().Msg("Enqueue new alert")
	initializers.WorkQueue <- alert
	log.Debug().Any("Payload", alert).Msg("Alert payload added to Work Queue")
}

func Dequeue() {
	poll_durn, err := strconv.Atoi(os.Getenv("POLL_DURN"))
	utils.CheckNilErr(err)
	defer initializers.Wg.Done()
	for {
		select {
		case alert_payload, ok := <-initializers.WorkQueue:
			if ok {
				log.Info().Msg("Processing new alert")
				log.Debug().Any("Payload", alert_payload).Msg("Processing new alert")
				processAlert(alert_payload)
				log.Info().Msg(fmt.Sprintf("Successfully processed alert. Waiting for %v milliseconds before polling again", poll_durn))
				time.Sleep(time.Duration(poll_durn) * time.Millisecond)
			} else {
				log.Panic().Msg("Channel is closed")
			}
		default:
			log.Info().Msg(fmt.Sprintf("No alert payload found. Checking again after %vms...", poll_durn))
			time.Sleep(time.Duration(poll_durn) * time.Millisecond)
		}
	}
}
