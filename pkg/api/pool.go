package api

import (
	"net/http"

	"github.com/Turkcell-Team-Atom-Devops/onent-devops/pkg/app"
	"github.com/Turkcell-Team-Atom-Devops/onent-devops/pkg/logger"
	"github.com/gorilla/mux"
)

// RecycleSingleAppPool ...
func RecycleSingleAppPool() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		poolName, ok := vars["pool"]
		logger.Log.Debug(poolName)
		if !ok {
			RespondWithError(w, http.StatusBadRequest, "poolName can not be empty.")
			return
		}

		poolService := app.Pool{Name: poolName}
		result, err := poolService.RecycleSingleAppPool()		
		if err != nil {
			RespondWithErrors(w, http.StatusServiceUnavailable, err.Error(), nil)
			return
		}
		
		RespondWithSuccess(w, http.StatusOK, result)
	}
}

// IISReset ...
func IISReset() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		poolService := app.Pool{}
		result, err := poolService.IISReset()
		if err != nil {
			RespondWithErrors(w, http.StatusServiceUnavailable, err.Error(), nil)
			return
		}
		RespondWithSuccess(w, http.StatusOK, result)
	}
}

// GetAll ...
func GetAll() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		poolService := app.Pool{}
		result, err := poolService.GetAll()
		if err != nil {
			RespondWithErrors(w, http.StatusServiceUnavailable, err.Error(), nil)
			return
		}
		RespondWithSuccess(w, http.StatusOK, result)
	}
}
