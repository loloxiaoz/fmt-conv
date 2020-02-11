package main

import (
	"github.com/loloxiaoz/fmt-conv/chrome"
	"github.com/loloxiaoz/fmt-conv/conf"
	"github.com/loloxiaoz/fmt-conv/xlog"
)

func main() {
	const op string = "main"
	config, err := conf.FromEnv()
	systemLogger := xlog.New(config.LogLevel(), "system")
	if err != nil {
		systemLogger.FatalOp(op, err)
	}
	// start Google Chrome headless.
	if err := chrome.Start(systemLogger); err != nil {
		systemLogger.FatalOp(op, err)
	}
}
