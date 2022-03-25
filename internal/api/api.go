// @title           Driver API
// @version         1.0
// @description     This is a Driver server.
// @termsOfService  https://www.selahattinceylan.com

// @contact.name   API Support
// @contact.url    https://www.selahattinceylan.com
// @contact.email  selahattinceylan9622@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8081
// @BasePath  /api/v1

package api

import (
	"net/http"

	_ "github.com/Selahattinn/bitaksi-driver/docs"
	"github.com/Selahattinn/bitaksi-driver/internal/api/response"
	"github.com/Selahattinn/bitaksi-driver/internal/service"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// API configuration
type Config struct {
	MatchServiceFlag string `yaml:"match_service_flag"`
}

// API represents the structure of the API
type API struct {
	Router  *mux.Router
	config  *Config
	Service service.Service
}

// New returns the api settings
func New(cfg *Config, router *mux.Router, svc service.Service) (*API, error) {
	api := &API{
		config:  cfg,
		Router:  router,
		Service: svc,
	}

	// Endpoint for browser preflight requests
	api.Router.Methods("OPTIONS").HandlerFunc(api.corsMiddleware(api.preflightHandler))

	// Endpoint for healtcheck
	api.Router.HandleFunc("/api/v1/health", api.corsMiddleware(api.logMiddleware(api.healthHandler))).Methods("GET")

	api.Router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8081/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	// Endpoint for Drivers
	api.Router.HandleFunc("/api/v1/driver", api.corsMiddleware(api.logMiddleware(api.AddDriver))).Methods("POST")
	api.Router.HandleFunc("/api/v1/driver", api.corsMiddleware(api.logMiddleware(api.UpdateDriver))).Methods("PUT")
	api.Router.HandleFunc("/api/v1/driver", api.corsMiddleware(api.logMiddleware(api.DeleteDriver))).Methods("DELETE")
	api.Router.HandleFunc("/api/v1/driver", api.corsMiddleware(api.logMiddleware(api.GetAllDrivers))).Methods("GET")
	api.Router.HandleFunc("/api/v1/driver/{id}", api.corsMiddleware(api.logMiddleware(api.FindDriver))).Methods("GET")

	// Endpoint for Search
	api.Router.HandleFunc("/api/v1/search", api.corsMiddleware(api.logMiddleware(api.authMiddleware(api.Search)))).Methods("POST")

	return api, nil
}

func (a *API) healthHandler(w http.ResponseWriter, r *http.Request) {
	response.Write(w, r, struct {
		Status string `json:"status"`
	}{
		"ok",
	})

}

func (a *API) preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

}
