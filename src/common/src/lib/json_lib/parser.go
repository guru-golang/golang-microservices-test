package json_lib

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

func Encode(m interface{}) string {
	b, err := json.Marshal(m)
	if err != nil {
		log.Error().Err(err)
	}
	return string(b)
}

func Decode(b interface{}, m string) interface{} {
	err := json.Unmarshal([]byte(m), &b)
	if err != nil {
		log.Info().Msgf("%v", m)
		log.Error().Err(err)
	}

	return b
}
