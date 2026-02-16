package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	Env             string
	HTTPAddr        string
	DBDSN           string
	TokenPepper     string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

func MustLoad() Config {
	return Config{
		Env:             get("ENV", "dev"),
		HTTPAddr:        get("HTTP_ADDR", ":8080"),
		DBDSN:           mustGet("DB_DSN"),
		TokenPepper:     mustGet("TOKEN_PEPPER"),
		ReadTimeout:     mustDuration(get("HTTP_READ_TIMEOUT", "5s")),
		WriteTimeout:    mustDuration(get("HTTP_WRITE_TIMEOUT", "10s")),
		ShutdownTimeout: mustDuration(get("SHUTDOWN_TIMEOUT", "10s")),
	}
}

func get(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func mustGet(k string) string {
	v := os.Getenv(k)
	if v == "" {
		panic(fmt.Sprintf("missing required env: %s", k))
	}
	return v
}

func mustDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic("invalid duration: " + s)
	}
	return d
}
