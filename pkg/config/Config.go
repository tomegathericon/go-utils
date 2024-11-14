package config

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/tomegathericon/go-utils/pkg/log"
	"os"
)

func Load(ctx context.Context) error {
	l := log.FromContext(ctx)
	env := os.Getenv("ENV")
	location := fmt.Sprintf("%s.env", env)
	if _, err := os.Open(location); err != nil {
		l.Error(err.Error())
		return err
	}
	return godotenv.Load(location)
}
