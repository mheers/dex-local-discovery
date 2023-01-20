package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mdaverde/jsonpath"
	"github.com/mheers/dex-local-discovery/helpers"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Port       int      `json:"port"`
	Upstream   string   `json:"upstream"`
	ReplaceOld string   `json:"replace_old"`
	ReplaceNew string   `json:"replace_new"`
	Endpoints  []string `json:"endpoints"`
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
		// get upstream
		upstream, err := helpers.GetURL(s.Config.Upstream)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// replace
		for _, endpointName := range s.Config.Endpoints {
			// get endpoint
			endpointValue, err := getEndpoint(upstream, endpointName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// replace
			endpointValue = strings.Replace(endpointValue, s.Config.ReplaceOld, s.Config.ReplaceNew, 1)

			// write
			upstream, err = setEndpoint(upstream, endpointName, endpointValue)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// write response
		w.Header().Set("Content-Type", "application/json")
		w.Write(upstream)
	})

	// http handler for config
	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
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

func getEndpoint(input []byte, name string) (string, error) {
	var payload interface{}

	err := json.Unmarshal([]byte(input), &payload)
	if err != nil {
		return "", err
	}

	endpoint, err := jsonpath.Get(payload, name)
	if err != nil {
		return "", err
	}

	return endpoint.(string), nil
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
