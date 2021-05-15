package utils

import "os"

func GetStringEnv(key, defaultValue string) string {
	env := os.Getenv(key)
	if env != "" {
		return env
	}
	return defaultValue
}
