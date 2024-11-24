package internal

import "context"

const (
	HeaderKeyContentType string = "Content-Type"
	HeaderKeyAccept      string = "Accept"
	HeaderKeyUserAgent   string = "User-Agent"
	HeaderKeyClientId    string = "Client-Id"
	HeaderKeyApiKey      string = "Api-Key"

	TypeApplicationJson string = "application/json"
)

type Client interface {
	Request(ctx context.Context, method string, path string, src interface{}, tgt interface{}) (Response, error)
}

type Response struct {
	Status int         `json:"-"`
	Result interface{} `json:"result"`
}
