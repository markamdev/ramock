package settings

import (
	"github.com/markamdev/goutils/sysenvs"
	"github.com/markamdev/ramock/internal/logging"
)

const (
	defListenPort = 8008
)

type Configuration struct {
	ListenPort  int
	StrictPaths bool
}

func ReadConfiguration(prefix string) Configuration {
	localLogger := logging.GetSubLogger("settings")

	localLogger.Info("reading configuration")
	nResolver := sysenvs.GetNewResolver(prefix)

	return Configuration{
		ListenPort:  sysenvs.GetIntEnvWithDefault(nResolver.GetVarName(varListenPort), defListenPort),
		StrictPaths: sysenvs.GetBoolEnvWithDefault(nResolver.GetVarName(varStrictPaths), false),
	}
}
