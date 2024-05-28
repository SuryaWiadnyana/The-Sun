package config

import "os"

func GetToken() string {
	return os.Getenv("API_TOKEN")
}
