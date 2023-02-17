package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mdaverde/jsonpath"
	"github.com/mheers/dex-local-discovery/helpers"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Port         int               `json:"port"`
	Upstream     string            `json:"upstream"`
	UpstreamHost string            `json:"upstream_host"`
	Endpoints    map[string]string `json:"endpoints"`
}

type Server struct {
	Config Config
}

func NewServer(config Config) *Server {
	return &Server{
		Config: config,
	}
}

func (s *Server) Run() error {
	configJSON, err := json.MarshalIndent(s.Config, "", "  ")
	if err != nil {
		return err
	}

	logrus.Info("Starting server")
	logrus.Infof("\nConfig: %s", string(configJSON))

	// http handler for openid-configuration
	http.HandleFunc("/.well-known/openid-configuration", func(w http.ResponseWriter, r *http.Request) {
		helpers.LogHTTPRequest(r)

		// get upstream
		upstream, err := helpers.GetURL(s.Config.Upstream, s.Config.UpstreamHost)
		if err != nil {
			logrus.Errorf("Error getting upstream: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// replace
		upstream, err = replaceEndpoints(upstream, s.Config.Endpoints)
		if err != nil {
			logrus.Errorf("Error replacing endpoints: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// write response
		w.Header().Set("Content-Type", "application/json")
		w.Write(upstream)
	})

	// http handler for config
	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		helpers.LogHTTPRequest(r)

		w.Header().Set("Content-Type", "application/json")
		w.Write(configJSON)
	})

	// start server
	err = http.ListenAndServe(fmt.Sprintf(":%d", s.Config.Port), nil)
	if err != nil {
		return err
	}

	return nil
}

func replaceEndpoints(upstream []byte, endpoints map[string]string) ([]byte, error) {
	for endpointName, endpointValue := range endpoints {
		var err error
		// write
		upstream, err = setEndpoint(upstream, endpointName, endpointValue)
		if err != nil {
			return nil, err
		}
	}
	return upstream, nil
}

func setEndpoint(input []byte, name, value string) ([]byte, error) {
	var payload interface{}

	err := json.Unmarshal(input, &payload)
	if err != nil {
		return nil, err
	}

	err = jsonpath.Set(&payload, name, value)
	if err != nil {
		return nil, err
	}

	result, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return nil, err
	}

	return result, nil
}
