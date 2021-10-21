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

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/vladimircunichin/golang/calendar/internal/adapters/inmemory"
	"github.com/vladimircunichin/golang/calendar/internal/domain/usecases"
	"github.com/vladimircunichin/golang/calendar/internal/logger"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "config/config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()
	vi := viper.New()
	vi.SetConfigFile(configFile)
	vi.ReadInConfig()
	logger.Init(vi.GetString("log_level"), vi.GetString("log_file"))

	calendars := usecases.New(inmemory.New())
	logger.Info("calendar was created")

	calendars.SaveEvent(context.Background(), "owner", "title", "text", time.Date(2021, time.April, 10, 21, 34, 15, 0, time.UTC), time.Date(2021, time.April, 11, 21, 34, 15, 0, time.UTC))
	editEvent, _ := calendars.GetEvents(context.Background())
	calendars.Edit(context.Background(), editEvent[0].ID, "NewOwner", "NewTitle", "NewText", time.Date(2022, time.April, 10, 21, 34, 15, 0, time.UTC), time.Date(2023, time.April, 11, 21, 34, 15, 0, time.UTC))
	_, err := calendars.GetEvents(context.Background())
	if err != nil {
		logger.Error(err.Error())
	}
	InitServer(vi.GetString("http_listen.ip"), vi.GetString("http_listen.port"))
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
