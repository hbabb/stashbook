// Package config loads Stashbook's settings from environment variables.
package config

import "github.com/caarlos0/env/v11"

// Config holds the settings Stashbook reads from environment variables.
// Doppler provides these when the app runs under `doppler run`.
type Config struct {
	DatabaseURL string `env:"DATABASE_URL,required"`
	Port        string `env:"PORT" envDefault:"8080"`
}

// Load reads the settings from the environment.
func Load() (Config, error) {
	return env.ParseAs[Config]()
}
