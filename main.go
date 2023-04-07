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

// {
// 	"data": {
// 	  "APIVersion": "v1",
// 	  "Action": "",
// 	  "Actions": null,
// 	  "Cluster": "test-cluster",
// 	  "Code": "",
// 	  "Count": 0,
// 	  "Error": "",
// 	  "Kind": "Pod",
// 	  "Level": "success",
// 	  "Messages": null,
// 	  "Name": "nginx-748c667d99-vjfwh",
// 	  "Namespace": "default",
// 	  "Reason": "",
// 	  "Recommendations": [
// 		"The 'latest' tag used in 'nginx' image of Pod 'default/nginx-748c667d99-vjfwh' container 'nginx' should be avoided."
// 	  ],
// 	  "Resource": "v1/pods",
// 	  "TimeStamp": "2023-04-06T09:45:56Z",
// 	  "Title": "v1/pods created",
// 	  "Type": "create",
// 	  "Warnings": null
// 	},
// 	"source": "k8s-recommendation-events",
// 	"timeStamp": "0001-01-01T00:00:00Z"
//   }

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
