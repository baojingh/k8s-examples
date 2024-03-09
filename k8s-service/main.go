package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"go.uber.org/automaxprocs/maxprocs"
)

var build = "develop"

func main() {

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
