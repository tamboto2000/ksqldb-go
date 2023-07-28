/*
Copyright © 2021 Thomas Meitz

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Parts of this apiclient are borrowed from Zalando Skipper
https://github.com/zalando/skipper/blob/master/net/httpclient.go

Zalando licence: MIT
https://github.com/zalando/skipper/blob/master/LICENSE
*/

package net

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/tamboto2000/ksqldbx/internal/utils"
	log "github.com/tamboto2000/ksqldbx/logger"
	"go.uber.org/zap"
)

const (
	DefaultIdleConnTimeout = 10 * time.Second
	DefaultBaseUrl         = "http://localhost:8088"
)

type HTTPClientFactory interface {
	NewHTTPClient(options Options, logger *zap.Logger) (*Client, error)
}

type HTTPClient interface {
	GetUrl(endpoint string) string
	Do(*http.Request) (*http.Response, error)
	Get(ctx context.Context, url string) (*http.Response, error)
	Post(ctx context.Context, url, contentType string, body io.Reader) (*http.Response, error)
	BasicAuth() string
	Close()
}

// The ksqlDB client
type Client struct {
	options Options
	uri     *url.URL
	client  http.Client
	tr      *Transport
	logger  log.Logger
}

// Credentials holds the username and password
type Credentials struct {
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
}

// NewHTTPClient creates a new http client
func NewHTTPClient(options Options, logger log.Logger) (*Client, error) {
	var uri *url.URL
	var err error

	if options.BaseUrl == "" {
		options.BaseUrl = DefaultBaseUrl
	}

	if uri, err = utils.GetUrl(options.BaseUrl); err != nil {
		return nil, fmt.Errorf("%+w", err)
	}

	tr := NewTransport(options)

	if logger == nil {
		logger, err = log.NewDefaultLogger()
		if err != nil {
			return nil, err
		}
	}

	return &Client{
		logger: logger,
		client: http.Client{
			Transport: tr,
		},
		options: options,
		tr:      tr,
		uri:     uri,
	}, nil
}

func (c *Client) Close() {
	c.tr.Close()
}

// Do delegates the given http.Request to the underlying http.Client
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

// Get executes a http request and returns a response or error
func (c *Client) Get(ctx context.Context, endpoint string) (*http.Response, error) {
	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s%s", c.options.BaseUrl, endpoint), nil)
	if err != nil {
		c.logger.Error(err.Error())
		return nil, err
	}
	return c.Do(req)
}

// GetUrl returns the full url for the requested endpoint
func (c *Client) GetUrl(endpoint string) string {
	return c.options.BaseUrl + endpoint
}

// Post executes a post request and returns the response or error.
func (c *Client) Post(ctx context.Context, endpoint, contentType string, body io.Reader) (*http.Response, error) {
	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s%s", c.options.BaseUrl, endpoint), body)
	if err != nil {
		c.logger.Error(err.Error())
		return nil, err
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	return c.Do(req)
}

func (c *Client) BasicAuth() string {
	if c.options.Credentials.Username != "" && c.options.Credentials.Password != "" {
		// Add the Authorization header to the request
		auth := c.options.Credentials.Username + ":" + c.options.Credentials.Password
		return base64.StdEncoding.EncodeToString([]byte(auth))
	}
	return ""
}
