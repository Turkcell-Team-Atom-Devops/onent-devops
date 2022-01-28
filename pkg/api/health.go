package api

import (
	"net/http"
)

var (
	//Port representd a server port
	Port = "3625"
	//ServerAddress represents a server addres
	ServerAddress = "0.0.0.0" + ":" + Port
)

// HealthProp ...
type HealthProp struct {
	StatusCode int   `json:"status_code"`
	Err        error `json:"error"`
}

// Services ...
type Services struct {
	API *HealthProp `json:"api"`
}

// HealthCheck ...
func HealthCheck() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var checkResult Services		
		APIStatus := getStatus(http.StatusOK, nil)

		checkResult = Services{
			API: APIStatus,
		}

		RespondWithJSON(w, checkResult.API.StatusCode, checkResult)
	}
}

func getStatus(state int, err error) *HealthProp {
	return &HealthProp{
		StatusCode: state,
		Err:        err,
	}
}
