package main

import (
	"os"

	"github.com/toaweme/log"

	"github.com/toaweme/cli"

	"github.com/toaweme/sintax/cmd/docs"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Error("error getting working directory", "error", err)
	}

	// prettyHandler := prettylog.NewHandler(&slog.HandlerOptions{
	// 	Level:       log.LevelTrace,
	// 	AddSource:   false,
	// 	ReplaceAttr: nil,
	// })
	//
	// logger := slog.New(prettyHandler)
	// log.SetLogger(logger)
	log.SetLevel(log.LevelTrace)

	app := cli.NewApp(cli.Settings{}, cli.GlobalOptions{
		Cwd: cwd,
	})

	app.Add("docs", docs.NewDocsCommand())

	err = app.Run(os.Args[1:])
	if err != nil {
		log.Error("failed to run app", "error", err)
		os.Exit(1)
		// log.Fatal().Err(err).Msg("Failed to run app")
	}
}
