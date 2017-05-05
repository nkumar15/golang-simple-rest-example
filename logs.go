package location

import (
	"fmt"

	logs "github.com/astaxie/beego/logs"
)

// LogConfig ...
type LogConfig struct {
	FileLocation string
}

var defaultLogger *logs.BeeLogger

// SetupLogger ...
func (env *Env) SetupLogger(logger *logs.BeeLogger, config LogConfig) error {
	if config.FileLocation == "" {
		config.FileLocation = "location.log"
	}

	beelogConfiguration := fmt.Sprintf(`{"filename":"%s","perm":"0664"}`, config.FileLocation)
	defaultLogger = logger
	err := defaultLogger.SetLogger(logs.AdapterFile, beelogConfiguration)

	return err
}

// GetLogger ...
func GetLogger() *logs.BeeLogger {
	return defaultLogger
}
