package api

import (
	"encoding/json"	
	"net/http"	
)

//ErrorResponseDTO represents error resposne
type ErrorResponseDTO struct {
	Code    int      `json:"code"`
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

//ErrorResponseDTO represents error resposne
type ResponseDTO struct {
	Code    int      `json:"code"`
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Payload interface{} `json:"payload"`
}


// RespondWithError ...
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, ErrorResponseDTO{Code: code, Status: "Error", Message: message})
}

// RespondWithErrors ...
func RespondWithErrors(w http.ResponseWriter, code int, message string, errors []string) {
	RespondWithJSON(w, code, ErrorResponseDTO{Code: code, Status: "Error", Message: message, Errors: errors})
}

// RespondWithJSON write json
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithSuccess(w http.ResponseWriter, code int, payload interface{}) {
	RespondWithJSON(w, code, ResponseDTO{Code: code, Status: "Success", Payload: payload})
}


