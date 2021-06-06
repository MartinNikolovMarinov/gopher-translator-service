package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/gopher-translator-service/internal"
	"github.com/gopher-translator-service/pkg/logger"
)

const (
	AppName = "Gopher Translation Service"
	Version = "0.0.0"
)

type ExitCode int

const (
	SucessEC          ExitCode = 0
	InteruptedEC      ExitCode = 1
	KilledEC          ExitCode = 2
	UnknownSignalEC   ExitCode = 3
	AppInitFailed     ExitCode = 4
	AppRunFailed      ExitCode = 5
	AppShutdownFailed ExitCode = 6
)

var portCMD = flag.Int("port", 8080, "Port to start the http server on")

func main() {
	flag.Parse()

	var app internal.App
	log := logger.NewStdLogger(logger.DebugLevel)
	appCfg := &internal.AppInitCfg{
		Port: *portCMD,
		Log: log,
	}
	if err := app.Init(appCfg); err != nil {
		log.Errorln(err)
		os.Exit(int(AppInitFailed))
	}

	sigCh := make(chan os.Signal, 0)
	signal.Notify(sigCh, os.Interrupt, os.Kill)
	go tryToGracefullyShutdown(sigCh, &app, log)

	log.Infof("Starting %s (v%s) on port: %d\n", AppName, Version, *portCMD)
	if err := app.Run(); err != nil {
		log.Errorln(err)
		os.Exit(int(AppRunFailed))
	}
}

// Some signals allows us some time for cleanup, we might want to gracefully shutdown the server in that case.
// If we have async logging it's a good idea to flush it before the process is terminated!
func tryToGracefullyShutdown(sigCh <-chan os.Signal, app *internal.App, log logger.Logger) {
	sig := <-sigCh
	log.Errorf("Signal received: %+v\n", sig)
	if err := app.Shutdown(); err != nil {
		log.Errorln(err)
		os.Exit(int(AppShutdownFailed))
	}

	switch sig {
	case os.Interrupt:
		os.Exit(int(InteruptedEC))
	case os.Kill:
		os.Exit(int(KilledEC))
	default:
		os.Exit(int(UnknownSignalEC))
	}
}
