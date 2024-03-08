// /infrastructure/api/coingecko_api.go
package api

import (
	"crypto/tls"
	"net/http"
	"time"
)

var Client *http.Client

func InitAPIClient() {
	tlsConfig := &tls.Config{}
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	Client = &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}
}
