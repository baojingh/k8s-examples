package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var build = "develop"

func main() {
	log, err := initLogger("SALES-API")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()
	if err := run(log); err != nil {
		log.ErrorW("startup", "ERROR", err)
		os.Exit(1)
	}

	if _, err := maxprocs.Set(); err != nil {
		fmt.Errorf("maxprocs %w", err)
	}

	g := runtime.GOMAXPROCS(0)

	log.Printf("starting service, CPU [%d]", g)
	defer log.Println("ended service")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Println("service is ending")

}

func run(log *zap.SugaredLogger) error {
	return nil
}

func initLogger(service string) (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	config.InitialFields = map[string]interface{}{
		"service": service,
	}
	log, err := config.Build()
	if err != nil {
		return nil, err
	}
	return log.Sugar(), nil

}
