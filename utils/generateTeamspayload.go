package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/DipayanP007/botkube-teams-proxy/initializers"
	"github.com/DipayanP007/botkube-teams-proxy/models"
	"github.com/fatih/structs"
	"github.com/rs/zerolog/log"
)

func generateTeamsPayload(payload models.Payload) string {

	themeColour := initializers.SuccessColour
	if strings.ToLower(payload.Level) == "error" {
		themeColour = initializers.ErrorColour
	}

	i := 0
	facts := []map[string]string{}
	payload_struct_to_map := structs.Map(&payload)
	for k, v := range payload_struct_to_map {
		if v != "" && reflect.TypeOf(v).Kind() == reflect.String {
			if k == "Messages" {
				facts = append(facts, map[string]string{"name": k, "value": payload.Messages[0]})
				i++
				continue
			}
			if k == "State" || k == "Cluster" {
				continue
			}
			facts = append(facts, map[string]string{"name": k, "value": payload.Get(k)})
			i++
		} else if reflect.TypeOf(v).Kind() == reflect.Slice && payload.Get(k) != "0" {
			if k == "Messages" {
				facts = append(facts, map[string]string{"name": k, "value": payload.Messages[0]})
				i++
				continue
			}
			if k == "State" || k == "Cluster" {
				continue
			}
			facts = append(facts, map[string]string{"name": k, "value": payload.Get(k)})
			i++
		} else {
			i++
		}
	}
	facts_in_string, err := json.Marshal(facts)
	CheckNilErr(err)
	log.Debug().Any("Payload", facts_in_string).Msg("Facts of alerts")
	finalPayload := fmt.Sprintf(`
	{
		"@type": "MessageCard",
		"@context": "http://schema.org/extensions",
		"themeColor": "%v",
		"summary": "%v",
		"sections": [{
			"activityTitle": "%v",
			"activitySubtitle": "on %v",
			"activityImage": "%v",
			"facts": %v,
			"markdown": true
		}]
	}
	`, themeColour, payload.State, payload.State, payload.Cluster, initializers.Icon, string(facts_in_string))
	log.Debug().Any("Payload", finalPayload).Msg("Final teams Payload")
	return finalPayload
}

func PostToTeams(pay string) int {
	payload := strings.NewReader(pay)
	log.Debug().Msg(os.Getenv("WEBHOOK_URL"))
	log.Info().Msg("Submitting payload")
	response, err := http.Post(os.Getenv("WEBHOOK_URL"), "application/json", payload)
	CheckNilErr(err)
	log.Info().Msg("Sent alert successfully")
	log.Debug().Int("Response code:", response.StatusCode).Msg("Response code from upstream")
	defer response.Body.Close()
	return response.StatusCode
}
