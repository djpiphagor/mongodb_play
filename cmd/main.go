package main

import (
	"context"
	"flag"
	"log/slog"
	"mongodb_play/internal/app"
	"mongodb_play/internal/config"
	"os"
	"os/signal"

	"github.com/pkg/errors"
)

func main() {
	path, err := cfgpather()
	if err != nil {
		slog.With(slog.Any("err", err)).Error("error config file path")
		os.Exit(1)
	}

	cfg, err := config.LoadConfig(path)
	if err != nil {
		slog.With(slog.Any("err", err)).Error("can't read a config file")
		os.Exit(1)
	}

	if err := applogger(cfg.Env); err != nil {
		slog.With(slog.Any("err", err)).Error("app logger error")
		os.Exit(1)
	}

	os.Exit(run(cfg))
}

func run(cfg *config.Config) (exitCode int) {

	defer func() {
		if panicErr := recover(); panicErr != nil {
			slog.With(slog.Any("err", panicErr)).Error("recover after panic")
			exitCode = 1
		}
	}()

	errCh := make(chan error, 1)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := app.New(cfg).Run(ctx); err != nil {
		errCh <- err
	}

	select {
	case err := <-errCh:
		slog.With(slog.Any("err", err)).Error("fatal error, the app shutdown")
		exitCode = 1
	case <-ctx.Done():
		slog.Info("the app shutdown")
	}

	return exitCode
}

func cfgpather() (string, error) {
	var path string
	flag.StringVar(&path, "config", "", "path to the config file")
	flag.Parse()
	if flag.NArg() < 1 {
		return path, errors.New("no args were detected")
	}
	if path == "" {
		return path, errors.New("empty config path")
	}

	if _, e := os.Stat(path); os.IsNotExist(e) {
		return path, errors.New("config file is not exist")
	}

	return path, nil
}

func applogger(level string) error {
	var l slog.Level
	switch level {
	case "local":
		l = slog.LevelDebug
	case "dev":
		l = slog.LevelInfo
	case "debug":
		l = slog.LevelDebug
	case "prod":
		l = slog.LevelError
	default:
		return errors.New("incorrect log level")
	}
	opts := &slog.HandlerOptions{
		Level: l,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	return nil
}
