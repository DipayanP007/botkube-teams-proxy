package models

import "strconv"

type Config struct {
	APIVersion      string   `json:"APIVersion,omitempty"`
	Action          string   `json:"Action,omitempty"`
	Actions         string   `json:"Actions,omitempty"`
	Cluster         string   `json:"Cluster,omitempty"`
	Code            string   `json:"Code,omitempty"`
	Count           int      `json:"Count,omitempty"`
	Error           string   `json:"Error,omitempty"`
	Kind            string   `json:"Kind,omitempty"`
	Level           string   `json:"Level,omitempty"`
	Messages        []string `json:"Messages,omitempty"`
	Name            string   `json:"Name,omitempty"`
	Namespace       string   `json:"Namespace,omitempty"`
	Reason          string   `json:"Reason,omitempty"`
	Recommendations []string `json:"Recommendations,omitempty"`
	Resource        string   `json:"Resource,omitempty"`
	TimeStamp       string   `json:"TimeStamp,omitempty"`
	Title           string   `json:"Title,omitempty"`
	Type            string   `json:"Type,omitempty"`
	Warnings        string   `json:"Warnings,omitempty"`
}

type Alert struct {
	Data      Config `json:"data"`
	Source    string `json:"source"`
	TimeStamp string `json:"timeStamp"`
}

type Payload struct {
	Action    string   `json:"Action,omitempty"`
	Actions   string   `json:"Actions,omitempty"`
	Cluster   string   `json:"Cluster,omitempty"`
	Error     string   `json:"Error,omitempty"`
	Level     string   `json:"Level,omitempty"`
	State     string   `json:"State"`
	Messages  []string `json:"Messages,omitempty"`
	TimeStamp string   `json:"TimeStamp,omitempty"`
	Warnings  string   `json:"Warnings,omitempty"`
	Reason    string   `json:"Reason,omitempty"`
}

func (p Payload) Get(k string) string {
	switch k {
	case "Action":
		return p.Action
	case "Actions":
		return p.Actions
	case "Cluster":
		return p.Cluster
	case "Error":
		return p.Error
	case "Level":
		return p.Level
	case "State":
		return p.State
	case "TimeStamp":
		return p.TimeStamp
	case "Messages":
		return strconv.Itoa(len(p.Messages))
	case "Warnings":
		return p.Warnings
	case "Reason":
		return p.Reason
	}
	return ""
}
