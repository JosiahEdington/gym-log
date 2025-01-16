package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/JosiahEdington/gym-log/logs"
)

type Server struct {
	Config  *Config
	Logger  *logs.Logger
	Handler http.Handler
}

func NewServer(
	config *Config,
	logger *logs.Logger,
) *Server {
	mux := http.NewServeMux()
	addRoutes(
		mux,
		logger,
		config,
	)
	var handler http.Handler = mux
	return &Server{
		Config:  config,
		Logger:  logger,
		Handler: handler,
	}
}

func StartServer(ctx context.Context, srv *Server) error {
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(srv.Config.Host, srv.Config.Port),
		Handler: srv.Handler,
	}

	go func() {
		log.Printf("listening on %s\n", httpServer.Addr)

		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()

		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()

	wg.Wait()
	return nil

}

func encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func decoder[T any](r *http.Request) (T, error) {
	var v T

	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

func decodeValid[T Validator](r *http.Request) (T, map[string]string, error) {
	var v T

	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		return v, nil, fmt.Errorf("decode json: %w", err)
	}

	problems := v.Valid(r.Context())
	if len(problems) > 0 {
		return v, problems, fmt.Errorf("invalid %T: %d problems", v, len(problems))
	}
	return v, nil, nil
}

type Validator interface {
	Valid(ctx context.Context) (problems map[string]string)
}
