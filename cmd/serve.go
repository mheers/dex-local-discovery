package cmd

import (
	"github.com/mheers/dex-local-discovery/helpers"
	"github.com/mheers/dex-local-discovery/server"
	"github.com/spf13/cobra"
)

var (
	port          int // port to listen on
	upstream      string
	upstream_host string

	issuer                        string
	authorization_endpoint        string
	token_endpoint                string
	jwks_uri                      string
	userinfo_endpoint             string
	device_authorization_endpoint string
	end_session_endpoint          string
	introspection_endpoint        string
	revocation_endpoint           string

	serveCmd = &cobra.Command{
		Use: "serve",
		RunE: func(cmd *cobra.Command, args []string) error {
			helpers.SetLogLevel(LogLevelFlag)
			endpoints := map[string]string{}

			if token_endpoint != "" {
				endpoints["token_endpoint"] = token_endpoint
			}
			if jwks_uri != "" {
				endpoints["jwks_uri"] = jwks_uri
			}
			if userinfo_endpoint != "" {
				endpoints["userinfo_endpoint"] = userinfo_endpoint
			}
			if device_authorization_endpoint != "" {
				endpoints["device_authorization_endpoint"] = device_authorization_endpoint
			}
			if authorization_endpoint != "" {
				endpoints["authorization_endpoint"] = authorization_endpoint
			}
			if issuer != "" {
				endpoints["issuer"] = issuer
			}
			if end_session_endpoint != "" {
				endpoints["end_session_endpoint"] = end_session_endpoint
			}
			if introspection_endpoint != "" {
				endpoints["introspection_endpoint"] = introspection_endpoint
			}
			if revocation_endpoint != "" {
				endpoints["revocation_endpoint"] = revocation_endpoint
			}

			config := server.Config{
				Port:         port,
				Upstream:     upstream,
				UpstreamHost: upstream_host,
				Endpoints:    endpoints,
			}

			s := server.NewServer(config)
			err := s.Run()
			if err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	serveCmd.Flags().IntVarP(&port, "port", "p", 5555, "port to listen on")
	serveCmd.Flags().StringVarP(&upstream, "upstream", "u", "https://dex.cluster.local/.well-known/openid-configuration", "upstream to proxy to")
	serveCmd.Flags().StringVarP(&upstream_host, "upstream_host", "U", "dex.cluster.local", "upstream host to request the upstream from (header manipulation)")
	serveCmd.Flags().StringVarP(&issuer, "issuer", "s", "", "issuer to replace")
	serveCmd.Flags().StringVarP(&authorization_endpoint, "authorization_endpoint", "a", "", "authorization_endpoint to replace")
	serveCmd.Flags().StringVarP(&token_endpoint, "token_endpoint", "t", "", "token_endpoint to replace")
	serveCmd.Flags().StringVarP(&jwks_uri, "jwks_uri", "j", "", "jwks_uri to replace")
	serveCmd.Flags().StringVarP(&userinfo_endpoint, "userinfo_endpoint", "i", "", "userinfo_endpoint to replace")
	serveCmd.Flags().StringVarP(&device_authorization_endpoint, "device_authorization_endpoint", "d", "", "device_authorization_endpoint to replace")
	serveCmd.Flags().StringVarP(&end_session_endpoint, "end_session_endpoint", "e", "", "end_session_endpoint to replace")
	serveCmd.Flags().StringVarP(&introspection_endpoint, "introspection_endpoint", "I", "", "introspection_endpoint to replace")
	serveCmd.Flags().StringVarP(&revocation_endpoint, "revocation_endpoint", "r", "", "revocation_endpoint to replace")
	rootCmd.AddCommand(serveCmd)
}
