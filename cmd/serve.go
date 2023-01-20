package cmd

import (
	"github.com/mheers/dex-local-discovery/server"
	"github.com/spf13/cobra"
)

var (
	port       int // port to listen on
	upstream   string
	replaceOld string
	replaceNew string
	endpoints  []string
	serveCmd   = &cobra.Command{
		Use: "serve",
		RunE: func(cmd *cobra.Command, args []string) error {

			config := server.Config{
				Port:       port,
				Upstream:   upstream,
				ReplaceOld: replaceOld,
				ReplaceNew: replaceNew,
				Endpoints:  endpoints,
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
	serveCmd.Flags().StringVarP(&replaceOld, "replace-old", "o", "https://dex.cluster.local/", "string to replace in upstream")
	serveCmd.Flags().StringVarP(&replaceNew, "replace-new", "n", "http://dex.dex:5556/", "string to replace with")
	serveCmd.Flags().StringSliceVarP(&endpoints, "endpoints", "e", []string{"token_endpoint", "jwks_uri", "userinfo_endpoint", "device_authorization_endpoint"}, "endpoints to rewrite")
	rootCmd.AddCommand(serveCmd)
}
