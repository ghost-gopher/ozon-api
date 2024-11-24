package seller

import (
	"errors"
	"github.com/ghost-gopher/ozon-api/internal"
	"github.com/ghost-gopher/ozon-api/internal/seller/product"
	"strings"
	"time"
)

const (
	DefaultSellerAPIBaseURI    = "https://api-seller.ozon.ru"
	DefaultSellerUserAgent     = "seller-pro/v1.0"
	DefaultSellerClientTimeout = 25
)

type (
	Seller struct {
		client  *client
		options *options

		product *product.Product
	}

	options struct {
		APIKey    string
		ClientId  string
		BaseURI   string
		UserAgent string
		Timeout   time.Duration
	}
)

func NewDefault() *Seller {
	opts := &options{
		BaseURI:   DefaultSellerAPIBaseURI,
		UserAgent: DefaultSellerUserAgent,
		Timeout:   DefaultSellerClientTimeout,
	}

	return &Seller{
		options: opts,
	}
}

func New(key string, id string) (*Seller, error) {
	slr := NewDefault()

	if err := slr.ChangeAPIKey(key); err != nil {
		return nil, err
	}

	if err := slr.ChangeClientId(id); err != nil {
		return nil, err
	}

	return slr, nil
}

func (s *Seller) ChangeAPIBaseURI(uri string) *Seller {
	if uri = strings.TrimSpace(uri); len(uri) == 0 {
		s.options.BaseURI = uri
	}

	return s
}

func (s *Seller) ChangeClientUserAgent(agent string) *Seller {
	if agent = strings.TrimSpace(agent); len(agent) == 0 {
		s.options.UserAgent = agent
	}

	return s
}

func (s *Seller) ChangeClientTimeout(timeout time.Duration) *Seller {
	if timeout > 0 {
		s.options.Timeout = timeout
	}

	return s
}

func (s *Seller) ChangeAPIKey(key string) error {
	if key = strings.TrimSpace(key); len(key) == 0 {
		return errors.New("parameter API Key is required")
	}

	s.options.APIKey = key

	return nil
}

func (s *Seller) ChangeClientId(id string) error {
	if id = strings.TrimSpace(id); len(id) == 0 {
		return errors.New("parameter Client Id is required")
	}

	s.options.ClientId = id

	return nil
}

func (s *Seller) httpclient() *client {
	if s.client == nil {
		opts := ClientOptions{
			internal.HeaderKeyContentType: internal.TypeApplicationJson,
			internal.HeaderKeyAccept:      internal.TypeApplicationJson,
			internal.HeaderKeyUserAgent:   s.options.UserAgent,
			internal.HeaderKeyClientId:    s.options.ClientId,
			internal.HeaderKeyApiKey:      s.options.APIKey,
		}
		s.client = NewClient(s.options.BaseURI, opts, s.options.Timeout)
	}

	return s.client
}

func (s *Seller) Product() *product.Product {
	if s.product == nil {
		s.product = product.New(s.httpclient())
	}

	return s.product
}
