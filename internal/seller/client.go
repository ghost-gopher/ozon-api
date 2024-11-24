package seller

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ghost-gopher/ozon-api/internal"
	"net/http"
	"time"
)

type ClientOptions map[string]string

type client struct {
	internal.Client
	host    string
	options ClientOptions
	timeout time.Duration
}

func NewClient(host string, options ClientOptions, timeout time.Duration) *client {
	cit := &client{
		host:    host,
		options: options,
		timeout: timeout,
	}

	return cit
}

func (c *client) Request(ctx context.Context, method string, path string, src interface{}, tgt interface{}) (internal.Response, error) {
	if len(c.options) == 0 {
		return internal.Response{}, errors.New("no options for headers of client provided")
	}

	clt := http.Client{
		Timeout:   time.Second * c.timeout,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{}},
	}

	domain := fmt.Sprintf("%s/%s", c.host, path)

	payload := &bytes.Buffer{}

	if src != nil {
		if err := json.NewEncoder(payload).Encode(src); err != nil {
			return internal.Response{}, errors.New(fmt.Sprintf("unable to marshal target: %s", err.Error()))
		}
	}

	rqt, err := http.NewRequestWithContext(ctx, method, domain, payload)
	if err != nil {
		return internal.Response{}, err
	}

	for key, value := range c.options {
		rqt.Header.Add(key, value)
	}

	answer, err := clt.Do(rqt)
	if err != nil {
		return internal.Response{}, errors.New(fmt.Sprintf("unable to make request: %s", err.Error()))
	}

	if answer == nil {
		return internal.Response{}, errors.New("unknown error")
	}

	rps := internal.Response{
		Status: answer.StatusCode,
		Result: tgt,
	}

	switch answer.StatusCode {
	case http.StatusOK:
		{
			if err := json.NewDecoder(answer.Body).Decode(&rps); err != nil {
				return rps, errors.New(fmt.Sprintf("unable to decode response: %s", err.Error()))
			}

			if err := answer.Body.Close(); err != nil {
				return rps, errors.New(fmt.Sprintf("unable to close response: %s", err.Error()))
			}
		}
	default:
		{
			return rps, errors.New(fmt.Sprintf("unexplored http status %s for client", answer.Status))
		}
	}
	return rps, nil
}
