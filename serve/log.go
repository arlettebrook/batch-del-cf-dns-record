package serve

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"

	nested "github.com/antonfisher/nested-logrus-formatter"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetFormatter(&nested.Formatter{
		TimestampFormat: time.DateTime,
		CallerFirst:     true,
		CustomCallerFormatter: func(f *runtime.Frame) string {
			return fmt.Sprintf(" %s:%d", filepath.Base(f.File), f.Line)
		},
	})
	Logger.SetReportCaller(true)

	switch GetConfig().LogLevel {
	case "debug":
		Logger.SetLevel(logrus.DebugLevel)
	}
}
