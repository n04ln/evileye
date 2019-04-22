package log

import "go.uber.org/zap"

var logger *zap.Logger

func init() {
	var err error
	// NOTE: if zap.NewDevelopment() is failed, panic it
	logger, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	logger.Info("initialize logger is successful")
}

// L returns *zap.Logger
func L() *zap.Logger {
	return logger
}
