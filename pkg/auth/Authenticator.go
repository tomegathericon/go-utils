package auth

import (
	"context"
	"errors"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
	*config
	ctx context.Context
}

func New(ctx context.Context) (*Authenticator, error) {
	c := newConfig()
	p, err := oidc.NewProvider(ctx, c.Domain())
	if err != nil {
		return nil, err
	}
	conf := oauth2.Config{
		ClientID:     c.ClientID(),
		ClientSecret: c.ClientSecret(),
		RedirectURL:  c.Callback(),
		Endpoint:     p.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	return &Authenticator{
		Provider: p,
		Config:   conf,
		ctx:      ctx,
	}, nil
}

func Must(ctx context.Context) *Authenticator {
	a, _ := New(ctx)
	return a
}

func (a *Authenticator) ExtractandVerifyIDToken(code string) (*oidc.IDToken, error) {
	token, err := a.Exchange(a.ctx, code)
	if err != nil {
		return nil, err
	}
	rawIdToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("id_token not found")
	}
	idToken, err := a.VerifierContext(a.ctx, &oidc.Config{ClientID: a.Config.ClientID}).Verify(a.ctx, rawIdToken)
	if err != nil {
		return nil, err
	}
	return idToken, nil
}
