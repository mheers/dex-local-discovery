package cmd

import (
	"github.com/mheers/dex-local-discovery/helpers"
	"github.com/mheers/dex-local-discovery/server"
	"github.com/spf13/cobra"
)

var (
	port     int // port to listen on
	upstream string

	issuer                        string
	authorization_endpoint        string
	token_endpoint                string
	jwks_uri                      string
	userinfo_endpoint             string
	device_authorization_endpoint string

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

			config := server.Config{
				Port:      port,
				Upstream:  upstream,
				Endpoints: endpoints,
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
	serveCmd.Flags().StringVarP(&issuer, "issuer", "s", "", "issuer to replace")
	serveCmd.Flags().StringVarP(&authorization_endpoint, "authorization_endpoint", "a", "", "authorization_endpoint to replace")
	serveCmd.Flags().StringVarP(&token_endpoint, "token_endpoint", "t", "", "token_endpoint to replace")
	serveCmd.Flags().StringVarP(&jwks_uri, "jwks_uri", "j", "", "jwks_uri to replace")
	serveCmd.Flags().StringVarP(&userinfo_endpoint, "userinfo_endpoint", "i", "", "userinfo_endpoint to replace")
	serveCmd.Flags().StringVarP(&device_authorization_endpoint, "device_authorization_endpoint", "d", "", "device_authorization_endpoint to replace")
	rootCmd.AddCommand(serveCmd)
}
