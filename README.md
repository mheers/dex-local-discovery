# dex-local-discovery

dex-local-discovery can be deployed next to [dex](https://github.com/dexidp/dex) to replace the external URLs with local URLs in the openid-configuration.

## Installation

### Install from source

```bash
git clone
cd dex-local-discovery
go install
```

### Install from docker

```bash
docker pull mheers/dex-local-discovery:latest
```

### Install from go

```bash
go install github.com/mheers/dex-local-discovery@latest
```

## Usage

```bash
dex-local-discovery serve \
    --port 5556 \
    --replace-old "https://dex.cluster.local/"
    --replace-new "http://dex.dex:5556/"
    --upstream "https://dex.cluster.local/.well-known/openid-configuration"
    --endpoints "token_endpoint,jwks_uri,userinfo_endpoint,device_authorization_endpoint"
```
