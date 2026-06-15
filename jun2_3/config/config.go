package config

import (
	"fmt"
	"strings"
)

type ConfigValidationError struct {
	Field string
}

func (e ConfigValidationError) Error() string {
	return fmt.Sprintf("Invalid field: %s", e.Field)
}

type Config struct {
	host string
	port int
}

func New(host string, port int) (Config, error) {
	host = strings.TrimSpace(host)
	if host == "" {
		return Config{}, ConfigValidationError{
			Field: "host",
		}
	}
	if port <= 0 {
		return Config{}, ConfigValidationError{
			Field: "port",
		}
	}

	return Config{
		host: host,
		port: port,
	}, nil
}

func (c Config) Host() string {
	return c.host
}
func (c Config) Port() int {
	return c.port
}
