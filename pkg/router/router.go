package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/Turkcell-Team-Atom-Devops/onent-devops/pkg/api"
)

// Router ...
type Router struct {
	router *mux.Router	
}

// New ...
func New() *Router {
	r := &Router{
		router: mux.NewRouter(),
	}
	r.initRoutes()
	return r
}

// ServeHTTP ...
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

func (r *Router) initRoutes() {
	// API Router Group
	apiRouter := mux.NewRouter().PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/getAll", api.GetAll()).Methods(http.MethodGet)
	apiRouter.HandleFunc("/recycle/{pool}", api.RecycleSingleAppPool()).Methods(http.MethodGet)
	apiRouter.HandleFunc("/iisreset", api.IISReset()).Methods(http.MethodGet)

	n := negroni.Classic()	
	r.router.PathPrefix("/api").Handler(n.With(	
		LimitHandler(),	
		negroni.Wrap(apiRouter),
	))

	healthRouter := mux.NewRouter().PathPrefix("/health").Subrouter()
	healthRouter.HandleFunc("/check", api.HealthCheck()).Methods(http.MethodGet)

	r.router.PathPrefix("/health").Handler(n.With(	
		LimitHandler(),	
		negroni.Wrap(healthRouter),
	))			
}
