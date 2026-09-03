package api

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/urbaniakmichal/data-generator/internal/config"
)

func NewServer(cs *config.ServerConfig) *http.Server {
	s := NewService()
	rh := NewRestHandler(s)

	mux := http.NewServeMux()
	mux.HandleFunc("GET "+ApiPathDataAsBatch, rh.GetDataAsBatch)
	mux.HandleFunc("GET "+ApiPathDataAsStream, rh.GetDataAsStream)
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	return &http.Server{
		Addr:              cs.Port,
		Handler:           mux,
		ReadHeaderTimeout: cs.ReadHeaderTimeout,
		ReadTimeout:       cs.ReadTimeout,
		WriteTimeout:      cs.WriteTimeout,
		IdleTimeout:       cs.IdleTimeout,
	}
}
