package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bitbucket.org/VladimirCunichin/golang/src/master/calendar/internal/adapters/inmemory"
	"bitbucket.org/VladimirCunichin/golang/src/master/calendar/internal/domain/config"
	"bitbucket.org/VladimirCunichin/golang/src/master/calendar/internal/domain/usecases"
	"bitbucket.org/VladimirCunichin/golang/src/master/calendar/internal/logger"
	"github.com/gorilla/mux"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "config/config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()
	conf := config.GetConfigFromFile(configFile)
	logger.Init(conf.Log.LogLevel, conf.Log.LogFile)
	_ = usecases.New(inmemory.New())
	logger.Info("calendar was created")
	InitServer(conf.HttpListen.Ip, conf.HttpListen.Port)
}

func InitServer(listenIP, listenPort string) {
	router := mux.NewRouter()
	router.HandleFunc("/hello", hello).Methods("GET")
	server := &http.Server{
		Addr:         listenIP + ":" + listenPort,
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal("Error while starting HTTP server", "error", err)
		}
	}()
	logger.Info("HTTP server started on host: " + listenIP + ", port: " + listenPort)
	<-done
	logger.Info("HTTP server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server shutdown error", "error", err)
	}
	logger.Info("server shut down properly")
}

func hello(w http.ResponseWriter, r *http.Request) {
	logger.Info("Incoming msg", "host", r.Host, "url", r.URL.Path)
	message := "Hello World"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", message)
}
