package rest

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/padmoney/common/credentials"
)

type Client interface {
	AddHeader(key, value string)
	Delete(url string) Response
	Get(url string) Response
	Header(key string) string
	Post(url string, body []byte) Response
	Put(url string, body []byte) Response
	Send(method, url string, body io.Reader) Response
}

type client struct {
	credentials credentials.Credentials
	headers     map[string]string
}

func NewClient(credentials credentials.Credentials) Client {
	c := &client{
		credentials: credentials,
		headers:     make(map[string]string),
	}
	c.setAuthorization()
	return c
}

func (r *client) Delete(url string) Response {
	return r.Send(http.MethodDelete, url, nil)
}

func (r *client) Get(url string) Response {
	return r.Send(http.MethodGet, url, nil)
}

func (r *client) Post(url string, body []byte) Response {
	r.AddHeader("Content-Type", "application/json")
	return r.Send(http.MethodPost, url, bytes.NewReader(body))
}

func (r *client) Put(url string, body []byte) Response {
	r.AddHeader("Content-Type", "application/json")
	return r.Send(http.MethodPut, url, bytes.NewReader(body))
}

func (c *client) Send(method, url string, body io.Reader) Response {
	c.setAuthorization()
	url = c.credentials.URL() + url
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return Response{Error: err}
	}
	c.setHeaders(req)
	httpClient := &http.Client{Timeout: 5 * time.Second}
	res, err := httpClient.Do(req)
	if err != nil {
		return Response{Error: err}
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Response{Error: err}
	}
	if res.StatusCode != http.StatusOK {
		return Response{
			Code:  res.StatusCode,
			Error: errors.New(string(content)),
		}
	}
	return Response{
		Body:  content,
		Code:  res.StatusCode,
		Error: err}
}

func (c *client) AddHeader(key, value string) {
	if c.headers == nil {
		c.headers = make(map[string]string)
	}
	c.headers[key] = value
}

func (c client) Header(key string) string {
	return c.headers[key]
}

func (c *client) setAuthorization() {
	basicToken := "Basic " + c.credentials.Token()
	c.AddHeader("Authorization", basicToken)
}

func (c *client) setHeaders(req *http.Request) {
	for k, v := range c.headers {
		req.Header.Add(k, v)
	}
}
