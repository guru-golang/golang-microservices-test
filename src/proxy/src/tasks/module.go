package tasks

import (
	"car-rent-platform/backend/common/src/lib/builtin_lib"
	"car-rent-platform/backend/proxy/src/tasks/proxy"
	"github.com/rs/zerolog/log"
	"sync"
)

func Init() {
	defer builtin_lib.Recovery()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer builtin_lib.Recovery()
		defer wg.Done()
		if err := proxy.New().Start(); err != nil {
			log.Error().Err(err)
		}
	}()
	wg.Wait()
}
