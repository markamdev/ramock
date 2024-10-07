package logging

import (
	"log/slog"
	"os"
	"path"

	"github.com/markamdev/goutils/sysenvs"
)

const (
	appTagName = "app"
)

var (
	appTagValue string
)

func init() {
	binaryName := path.Base(os.Args[0])
	appTagValue = sysenvs.GetStringEnvWithDefault("LOGGING_APP_TAG", binaryName)
	defLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With(appTagName, appTagValue)
	slog.SetDefault(defLogger)
}

func GetSubLogger(moduleName string) *slog.Logger {
	result := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With(appTagName, appTagValue).With("package", moduleName)
	return result
}
