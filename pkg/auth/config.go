package auth

import "os"

type config struct {
	clientID, clientSecret, domain, callback string
}

func (c *config) Callback() string {
	return c.callback
}

func (c *config) SetCallback(callback string) {
	c.callback = callback
}

func (c *config) ClientID() string {
	return c.clientID
}

func (c *config) SetClientID(clientID string) {
	c.clientID = clientID
}

func (c *config) ClientSecret() string {
	return c.clientSecret
}

func (c *config) SetClientSecret(clientSecret string) {
	c.clientSecret = clientSecret
}

func (c *config) Domain() string {
	return c.domain
}

func (c *config) SetDomain(domain string) {
	c.domain = domain
}

func newConfig() *config {
	return &config{
		clientID:     os.Getenv("CLIENT_ID"),
		clientSecret: os.Getenv("CLIENT_SECRET"),
		domain:       os.Getenv("AUTH_DOMAIN"),
		callback:     os.Getenv("AUTH_CALLBACK"),
	}
}
