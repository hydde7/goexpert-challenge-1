package main

import (
	"os"
	"os/signal"

	"github.com/hydde7/goexpert-challenge-1/cmd"
	"github.com/hydde7/goexpert-challenge-1/internal/cfg"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "goexpert-challenge-1"
	app.Usage = "Challenge 1 for GoExpert"
	app.Flags = cfg.Flags
	app.Action = cli.ActionFunc(run)
	go app.Run(os.Args)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logrus.Info("Shutting down...")
}

func run(c *cli.Context) error {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	level, err := logrus.ParseLevel(cfg.App.LogLevel)
	if err != nil {
		logrus.WithError(err).Fatal("failed to parse log level")
	}
	logrus.SetLevel(level)

	router := cmd.SetupRouter()

	err = router.Run(cfg.App.Address)
	if err != nil {
		logrus.WithError(err).Fatal("failed to start server")
	}

	return err
}
