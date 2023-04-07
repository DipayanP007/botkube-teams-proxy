package utils

import "github.com/rs/zerolog/log"

func CheckNilErr(err error) {
	if err != nil {
		log.Panic().Err(err)
	}
}
