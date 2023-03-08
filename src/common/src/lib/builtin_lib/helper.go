package builtin_lib

import (
	"github.com/rs/zerolog/log"
	"runtime"
	"strings"
)

func Recovery() {
	if err := recover(); err != nil {
		log.Error().Msgf("%v", err)
	}
}

func GetLocalPkgName() string {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	funcNameSlice := strings.Split(funcName, "/")
	funcNameSlice = strings.Split(funcNameSlice[len(funcNameSlice)-1:][0], ".")
	return funcNameSlice[0]
}
