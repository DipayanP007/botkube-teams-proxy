package utils

import (
	"fmt"

	"github.com/DipayanP007/botkube-teams-proxy/models"

	"github.com/rs/zerolog/log"
)

func ParsePayload(alert models.Alert) string {
	payload := models.Payload{
		Action:    alert.Data.Action,
		Actions:   alert.Data.Actions,
		Cluster:   alert.Data.Cluster,
		Error:     alert.Data.Error,
		State:     fmt.Sprintf("%v %v/%v is in state %v", alert.Data.Kind, alert.Data.Namespace, alert.Data.Name, alert.Data.Title),
		Level:     alert.Data.Level,
		Messages:  alert.Data.Messages,
		TimeStamp: alert.Data.TimeStamp,
		Warnings:  alert.Data.Warnings,
		Reason:    alert.Data.Reason,
	}
	log.Debug().Any("Payload", payload).Msg("Parsed alert payload")
	return generateTeamsPayload(payload)

}
