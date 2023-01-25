package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReplaceEndpoints(t *testing.T) {
	upstream := []byte(`{
		"issuer": "http://dex-local-discovery.dex:5556",
		"authorization_endpoint": "http://dex-local-discovery.dex:5556/auth",
		"token_endpoint": "http://dex-local-discovery.dex:5556/token",
		"jwks_uri": "http://dex-local-discovery.dex:5556/keys",
		"userinfo_endpoint": "http://dex-local-discovery.dex:5556/userinfo",
		"device_authorization_endpoint": "http://dex-local-discovery.dex:5556/device/code",
		"grant_types_supported": [
		  "authorization_code",
		  "refresh_token",
		  "urn:ietf:params:oauth:grant-type:device_code"
		],
		"response_types_supported": [
		  "code"
		],
		"subject_types_supported": [
		  "public"
		],
		"id_token_signing_alg_values_supported": [
		  "RS256"
		],
		"code_challenge_methods_supported": [
		  "S256",
		  "plain"
		],
		"scopes_supported": [
		  "openid",
		  "email",
		  "groups",
		  "profile",
		  "offline_access"
		],
		"token_endpoint_auth_methods_supported": [
		  "client_secret_basic",
		  "client_secret_post"
		],
		"claims_supported": [
		  "iss",
		  "sub",
		  "aud",
		  "iat",
		  "exp",
		  "email",
		  "email_verified",
		  "locale",
		  "name",
		  "preferred_username",
		  "at_hash"
		]
	  }`)

	endpoints := map[string]string{
		"issuer": "https://dex.cluster.local",
	}

	result, err := replaceEndpoints(upstream, endpoints)
	require.NoError(t, err)
	require.NotNil(t, result)
}
