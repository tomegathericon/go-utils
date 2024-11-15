package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/tomegathericon/go-utils/pkg/log"
	"os"
)

func Load(ctx context.Context) error {
	l := log.FromContext(ctx)
	env, ok := os.LookupEnv("ENV")
	if !ok {
		l.Error("ENV variable not set")
		return errors.New("ENV variable not set")
	}
	location := fmt.Sprintf("%s.env", env)
	configPath, ok := os.LookupEnv("CONFIG_PATH")
	if ok {
		location = fmt.Sprintf("%s/%s", configPath, location)
	}
	if _, err := os.Open(location); err != nil {
		l.Error(err.Error())
		return err
	}
	return godotenv.Load(location)
}
