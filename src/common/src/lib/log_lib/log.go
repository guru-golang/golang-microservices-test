package log_lib

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"car-rent-platform/backend/common/src/lib/config_lib"
)

var levels = map[string]zerolog.Level{
	"info":     zerolog.InfoLevel,
	"debug":    zerolog.DebugLevel,
	"error":    zerolog.ErrorLevel,
	"warn":     zerolog.WarnLevel,
	"fatal":    zerolog.FatalLevel,
	"panic":    zerolog.PanicLevel,
	"no":       zerolog.NoLevel,
	"disabled": zerolog.Disabled,
	"trace":    zerolog.TraceLevel,
}

func Init() {
	setLogLevel()
}

type MessageHook zerolog.LevelHook

func (h MessageHook) Run(e *zerolog.Event, l zerolog.Level, msg string) {
	if l == zerolog.ErrorLevel {
	}
}

func setLogLevel() {

	var level zerolog.Level
	logLevel := config_lib.Config.Get("app_log").(string)
	level = levels[logLevel]
	wd, _ := os.Getwd()

	format := func(i interface{}) string { return fmt.Sprintf("[%s]", i) }
	emptyFormat := func(i interface{}) string { return "" }
	beforeFormat := func(i interface{}) string { return fmt.Sprintf("[%s: ", i) }
	afterFormat := func(i interface{}) string { return fmt.Sprintf("%s]", i) }
	callerFormat := func(i interface{}) string {
		return strings.Replace(fmt.Sprintf("[%s]", i), strings.Replace(wd, "\\", "/", -1), "", 1)
	}

	output := zerolog.NewConsoleWriter()
	output.TimeFormat = time.Stamp
	output.FormatTimestamp = emptyFormat
	output.FormatLevel = format
	output.FormatLevel = format
	output.FormatCaller = format
	output.FormatMessage = format
	output.FormatFieldName = beforeFormat
	output.FormatFieldValue = afterFormat
	output.FormatErrFieldName = beforeFormat
	output.FormatErrFieldValue = afterFormat
	output.FormatCaller = callerFormat
	output.NoColor = true

	zerolog.SetGlobalLevel(level)
	zerolog.DisableSampling(true)

	var messageHook = MessageHook(zerolog.NewLevelHook())

	log.Logger = zerolog.New(output).With().Caller().Logger().Hook(messageHook)
}
