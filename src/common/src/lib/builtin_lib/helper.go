package builtin_lib

import (
	"github.com/rs/zerolog/log"
)

func Recovery() {
	if err := recover(); err != nil {
		log.Error().Msgf("%v", err)
	}
}
