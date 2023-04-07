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

	var alert models.Alert
	json.NewDecoder(r.Body).Decode(&alert)
	jsonformap, err := json.Marshal(alert)
	utils.CheckNilErr(err)
	log.Info().Msg("Payload received")
	log.Debug().Any("Payload", string(jsonformap))

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
