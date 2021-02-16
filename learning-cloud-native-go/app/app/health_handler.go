package app

import (
	"net/http"
)

func HandleLive(w http.ResponseWriter, _ *http.Request) {
	writeHealthy(w)
}

func (app *App) HandleReady(w http.ResponseWriter, r *http.Request) {
	if err := app.db.DB().Ping(); err != nil {
		app.Logger().Fatal().Err(err).Msg("")
		writeUnhealthy(w)
		return
	}
	writeHealthy(w)
}
func writeHealthy(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("."))
}
func writeUnhealthy(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("."))
}
