package log

import (
	"encoding/json"
	"go.uber.org/zap"
)

var Sugar *zap.SugaredLogger

func init() {
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "pmserver.log"],
	  "errorOutputPaths": ["stderr", "errors.log"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	Sugar = logger.Sugar()
	Sugar.Info("logger construction succeeded")

}
