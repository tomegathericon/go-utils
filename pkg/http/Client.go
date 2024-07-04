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
	headers    http.Header
	log        *log.Logger
}

func (c *Client) Headers() http.Header {
	return c.headers
}

func (c *Client) SetHeaders(headers http.Header) {
	c.headers = headers
}

func (c *Client) SetLog(log *log.Logger) {
	c.log = log
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
		log:        log.Must("json"),
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

func (c *Client) GET() (*http.Response, error) {
	req, err := http.NewRequest("GET", c.url.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.headers != nil {
		req.Header = c.headers
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
