package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

const baseURL = "https://pokeapi.co/api/v2"

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
