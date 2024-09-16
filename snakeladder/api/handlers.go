package api

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/snakeladder/engine"
)

var gameEngine *engine.Engine

func SetGameEngine(e *engine.Engine) {
	gameEngine = e
}

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	gameEngine.Play()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Game Played"))
}


func MetricsHandler() http.Handler {
	return promhttp.Handler()
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

