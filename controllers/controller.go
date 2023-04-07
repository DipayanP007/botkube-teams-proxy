package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DipayanP007/botkube-teams-proxy/models"
	"github.com/DipayanP007/botkube-teams-proxy/service"
	"github.com/DipayanP007/botkube-teams-proxy/utils"

	"github.com/rs/zerolog/log"
)

func DisplayPayload(w http.ResponseWriter, r *http.Request) {

	/*
		var json_out interface{}
		json.NewDecoder(r.Body).Decode(&json_out)
		jsonformap, _ := json.Marshal(json_out) //, "", "  ")
		// log.Println("This is request")
		// log.Println(string(jsonformap))
		log.Println("Going to send alert")
		response_from_upstream := Post(string(jsonformap))
		log.Println("Response from upstream", response_from_upstream)
		json.NewEncoder(w).Encode(map[string]string{
			"response": response_from_upstream,
		})
		// json.NewEncoder(w).Encode(string(jsonformap))
	*/

	var alert models.Alert
	json.NewDecoder(r.Body).Decode(&alert)
	jsonformap, err := json.Marshal(alert) //, "", "  ")
	utils.CheckNilErr(err)
	log.Info().Msg("Payload received")
	log.Debug().Any("Payload", string(jsonformap))
	// log.Println(alert)
	// payload := models.Payload{
	// 	Action:    alert.Data.Action,
	// 	Actions:   alert.Data.Actions,
	// 	Cluster:   alert.Data.Cluster,
	// 	Error:     alert.Data.Error,
	// 	State:     fmt.Sprintf("%v %v/%v is in state %v", alert.Data.Kind, alert.Data.Namespace, alert.Data.Name, alert.Data.Title),
	// 	Level:     alert.Data.Level,
	// 	Messages:  alert.Data.Messages,
	// 	TimeStamp: alert.Data.TimeStamp,
	// 	Warnings:  alert.Data.Warnings,
	// 	Reason:    alert.Data.Reason,
	// }
	service.Enqueue(alert)
	w.Write([]byte("Request enqueued"))
}

func Health(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Health Check success")
	json.NewEncoder(w).Encode(map[string]string{
		"healthy": "true",
	})
}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Home Page")
	w.WriteHeader(201)
	w.Write([]byte("Webhook app"))
}
