package server

import (
	"context"
	"fmt"
	"github.com/cc2k19/go-tin/web"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// Settings type to be loaded from the environment
type Settings struct {
	Port int `mapstructure:"port" description:"port of the server"`
}

// DefaultSettings returns the default values for configuring the server
func DefaultSettings() *Settings {
	return &Settings{
		Port: 8080,
	}
}

// Validate validates the server settings
func (s *Settings) Validate() error {
	if s.Port == 0 {
		return fmt.Errorf("validate Settings: Port missing")
	}
	return nil
}

// Server is the server to process incoming HTTP requests
type Server struct {
	*mux.Router

	Config *Settings
}

// New creates a new server with the provided REST api configuration and server configuration
func New(config *Settings, api *web.API) *Server {
	router := mux.NewRouter().StrictSlash(true)

	for _, controller := range api.Controllers {
		routes := controller.Routes()
		for _, route := range routes {
			router.HandleFunc(route.Endpoint.Path, route.Handler).Methods(route.Endpoint.Method)
		}
	}

	return &Server{
		Router: router,
		Config: config,
	}
}

// Run starts the server awaiting for incoming requests
func (s *Server) Run(ctx context.Context, wg *sync.WaitGroup) {
	if err := s.Config.Validate(); err != nil {
		log.Panicf("invalid server config: %s\n", err)
	}

	handler := &http.Server{
		Handler: s.Router,
		Addr:    ":" + strconv.Itoa(s.Config.Port),
	}

	startServer(ctx, handler, wg)
}

func startServer(ctx context.Context, server *http.Server, wg *sync.WaitGroup) {
	wg.Add(1)
	go gracefulShutdown(ctx, server, wg)

	log.Printf("Server listening on %s...\n", server.Addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func gracefulShutdown(ctx context.Context, server *http.Server, wg *sync.WaitGroup) {
	<-ctx.Done()
	defer wg.Done()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Error: ", err)
		if err := server.Close(); err != nil {
			log.Println("Error: ", err)
		}
	} else {
		log.Println("Server stopped")
	}
}
