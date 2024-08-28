package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
)

func InitLogger() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
	log.SetReportCaller(true)

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf(" [%s:%d] ", f.Func.Name(), f.Line)
		},
	})
}
