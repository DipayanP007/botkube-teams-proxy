package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"

	"net/http"

	"github.com/DipayanP007/botkube-teams-proxy/initializers"
	"github.com/DipayanP007/botkube-teams-proxy/routers"
	"github.com/DipayanP007/botkube-teams-proxy/service"
)

func main() {
	initializers.LoadENV()
	log.Info().Msg("Starting work queue consumer")
	initializers.Wg.Add(1)
	go service.Dequeue()
	log.Info().Msg(fmt.Sprintf("Listening on port :%v...", os.Getenv("PORT")))
	err := http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), routers.Router())
	log.Fatal().Err(err)
	initializers.Wg.Wait()
}
