package settings

import (
	"github.com/markamdev/goutils/sysenvs"
	"github.com/markamdev/ramock/internal/logging"
)

const (
	defListenPort    = 8008
	defEndpointsFile = "endpoints.yaml"
)

type Configuration struct {
	ListenPort    int
	StrictPaths   bool
	EndpointsFile string
}

func ReadConfiguration(prefix string) Configuration {
	localLogger := logging.GetSubLogger("settings")

	localLogger.Info("reading configuration")
	nResolver := sysenvs.GetNewResolver(prefix)

	return Configuration{
		ListenPort:    sysenvs.GetIntEnvWithDefault(nResolver.GetVarName(varListenPort), defListenPort),
		StrictPaths:   sysenvs.GetBoolEnvWithDefault(nResolver.GetVarName(varStrictPaths), false),
		EndpointsFile: sysenvs.GetStringEnvWithDefault(nResolver.GetVarName(endpointsFile), defEndpointsFile),
	}
}

func ReadConfigurationFromEnv() Configuration {
	return ReadConfiguration("RAMOCK")
}
