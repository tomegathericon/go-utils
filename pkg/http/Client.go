package http

import (
	"fmt"
	"github.com/tomegathericon/go-utils/pkg/log"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	url        *url.URL
	body       io.Reader
	httpClient *http.Client
	headers    map[string]string
	log        *log.Logger
}

func (c *Client) SetLog(log *log.Logger) {
	c.log = log
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) Url() *url.URL {
	return c.url
}

func (c *Client) SetUrl(url *url.URL) {
	c.url = url
}

func (c *Client) Body() io.Reader {
	return c.body
}

func (c *Client) SetBody(body io.Reader) {
	c.body = body
}

func (c *Client) Headers() map[string]string {
	return c.headers
}

func (c *Client) SetHeaders(headers map[string]string) {
	c.headers = headers
}

func (c *Client) GET() (*http.Response, error) {
	req, err := http.NewRequest("GET", c.url.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.headers != nil {
		for k, v := range c.headers {
			req.Header.Set(k, v)
		}
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		err := fmt.Errorf("http status code %d", res.StatusCode)
		c.log.Error(err.Error())
		return nil, err
	}
	return res, nil
}
