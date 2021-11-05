package main

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

// `Config` is the configuration of the inspector.
type Config struct {
	// Binance is the Binance configuration for accessing the API.
	Binance struct {
		// ApiURL is the API URL.
		ApiURL string

		// ApiKey is the key for accessing the user API.
		ApiKey string

		// SecretKey is the secret key used in conjunction with the ApiKey.
		SecretKey string
	}
}

// `getConfigurationEnv` gets the inspector configuration from a `.env` file.
// It returns the configurations `Config` and an error, if any.
func getConfigurationEnv() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}
	apiKey, apiFound := os.LookupEnv("API_KEY")
	secretKey, secretFound := os.LookupEnv("SECRET_KEY")
	apiURL, urlFound := os.LookupEnv("API_URL")
	if !(apiFound && secretFound && urlFound) {
		return Config{}, errors.New("cannot load configuration with empty values")
	}
	return Config{
		Binance: struct {
			ApiURL    string
			ApiKey    string
			SecretKey string
		}{
			ApiURL:    apiURL,
			ApiKey:    apiKey,
			SecretKey: secretKey,
		},
	}, nil
}
