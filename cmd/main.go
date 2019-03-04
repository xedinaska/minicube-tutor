package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/xedinaska/minicube-tutor/server"
	"net/http"
)

const timeout = 5*time.Second

func main() {
	log.SetLevel(log.DebugLevel)

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		sig := <-gracefulStop
		log.Infof("received control signal: %+v", sig)
		os.Exit(0)
	}()

	s := server.Server{
		Router: http.NewServeMux(),
	}
	s.Routes()

	srv := http.Server{
		Handler: s.Router,
		IdleTimeout: timeout,
		ReadTimeout: timeout,
		WriteTimeout: timeout,
		Addr:":8888",
	}

	log.Infof("starting app at :8888...")

	srv.ListenAndServe()
}
