package plex

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type App struct {
	c *httpClient
	o *Options
}

func (a *App) get(url, authToken string) ([]byte, error) {
	return a.c.get(url, authToken)
}

func (a *App) post(url string, body io.Reader) ([]byte, error) {
	return a.c.post(url, body)
}

func (a *App) command(url, authToken string) error {
	return a.c.command(url, authToken)
}

type Options struct {
	client           http.Client
	clientIdentifier string
	logger           *log.Logger
	product          string
}

type Option = func(*Options)

func Product(product string) Option {
	return func(a *Options) {
		a.product = product
	}
}

func Logger(logger *log.Logger) Option {
	return func(a *Options) {
		a.logger = logger
	}
}

func Client(client http.Client) Option {
	return func(a *Options) {
		a.client = client
	}
}

var DefaultAppOptions = [...]Option{
	Logger(log.Default()),
	Product("go-plex"),
	Client(http.Client{Timeout: 10 * time.Second}),
}

func New(clientIdentifier string, opts ...Option) *App {
	o := &Options{clientIdentifier: clientIdentifier}

	// Default options
	for _, i := range DefaultAppOptions {
		i(o)
	}
	for _, i := range opts {
		i(o)
	}
	a := &App{
		c: &httpClient{},
		o: o,
	}
	a.c.a = a
	return a
}

func (a *App) User(authToken string) (*User, error) {
	d, err := a.get(APIBaseURLv2+"/user", authToken)
	if err != nil {
		return nil, err
	}
	var u User
	err = json.Unmarshal(d, &u)
	if err != nil {
		return nil, err
	}
	u.app = a
	return &u, err
}
