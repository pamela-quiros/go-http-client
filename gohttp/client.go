package gohttp

import (
	"net/http"
	"time"
)

type httpClient struct {
	client             *http.Client
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	Headers            http.Header
}

func New() HttpClient {

	httpClient := &httpClient{}
	return httpClient
}

type HttpClient interface {
	SetHeaders(headers http.Header)
	SetConnectionTimeout(timeout time.Duration)
	SetResponsetTimeout(timeout time.Duration)
	SerMaxIdleConnections(i int)

	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

func (c *httpClient) SetHeaders(headers http.Header) {
	c.Headers = headers
}

func (c *httpClient) SetConnectionTimeout(timeout time.Duration) {
	c.connectionTimeout = timeout
}

func (c *httpClient) SetResponsetTimeout(timeout time.Duration) {
	c.responseTimeout = timeout
}

func (c *httpClient) SerMaxIdleConnections(i int) {
	c.maxIdleConnections = i
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
