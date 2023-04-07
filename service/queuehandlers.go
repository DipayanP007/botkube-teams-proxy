package service

import (
	"fmt"

	"github.com/rs/zerolog/log"
	// "log"
	"time"

	"github.com/DipayanP007/botkube-teams-proxy/models"
	"github.com/DipayanP007/botkube-teams-proxy/utils"
)

func processAlert(alert models.Alert) int {
	for i := 1; i <= 5; i++ {
		teams_payload := utils.ParsePayload(alert)
		// response_from_teams :=
		log.Debug().Any("Payload", teams_payload).Msg("Acquired final payload")
		if response_from_teams := utils.PostToTeams(teams_payload); response_from_teams == 200 {
			return response_from_teams
		}
		log.Warn().Msg(fmt.Sprintf("Failed to send alert to teams. Retries left %v. Retrying in 5000 milliseconds...", (5 - i)))
		time.Sleep(5 * time.Second)
	}
	log.Error().Msg("Not able to send alert to Teams. Server might be down or webhook URL is not accessible")
	log.Fatal().Msg("Exiting")
	return 0
}
