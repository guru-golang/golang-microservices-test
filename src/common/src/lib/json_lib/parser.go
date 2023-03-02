package json_lib

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

func Encode(m interface{}) string {
	b, err := json.Marshal(m)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	return string(b)
}

func Decode[T any](b T, m string) T {
	err := json.Unmarshal([]byte(m), &b)
	if err != nil {
		log.Info().Msgf("%v", m)
		log.Error().Msg(err.Error())
	}

	return b
}

func DecodeWithTimeout(decoder *json.Decoder, timeout time.Duration, value any) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	errch := make(chan error, 1)

	go func() { errch <- decoder.Decode(value) }()

	select {
	case err := <-errch:
		return err
	case <-ctx.Done():
		fmt.Println("timeout")
		return errors.New("decode timeout") // or any other error you prefer
	}
}
