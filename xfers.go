package xfers

import (
	"net/http"
	"time"
)

// Client is xfers client.
type Client struct {
	apiKey    string
	secretKey string
	baseURL   string
	env       EnvironmentType
	requester Requester
	logger    Logger
}

// Option is config for xfers client.
type Option struct {
	APIKey    string
	SecretKey string
	BaseURL   string
	Env       EnvironmentType
	Requester Requester
	Logger    Logger
}

// New to create new xfers client with config.
func New(option Option) *Client {
	if option.Logger == nil {
		option.Logger = defaultLogger(LogError)
	}

	if option.Requester == nil {
		option.Requester = defaultRequester(&http.Client{
			Timeout: 10 * time.Second,
		}, option.Logger)
	}

	return &Client{
		apiKey:    option.APIKey,
		secretKey: option.SecretKey,
		baseURL:   option.BaseURL,
		requester: option.Requester,
		logger:    option.Logger,
		env:       option.Env,
	}
}

// NewDefault to create new xfers client with default config.
func NewDefault(apiKey string, secretKey string, env EnvironmentType) *Client {
	return New(Option{
		APIKey:    apiKey,
		SecretKey: secretKey,
		BaseURL:   envURL[env],
		Requester: defaultRequester(&http.Client{
			Timeout: 10 * time.Second,
		}, defaultLogger(envLog[env])),
		Logger: defaultLogger(envLog[env]),
		Env:    env,
	})
}
