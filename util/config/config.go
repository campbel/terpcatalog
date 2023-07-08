package config

import "os"

func Port() string {
	return getEnv("PORT", "8080")
}

func AdminPort() string {
	return getEnv("ADMIN_PORT", "8081")
}

func IsDevelopment() bool {
	return Environment() == "development"
}

func Environment() string {
	return getEnv("ENVIRONMENT", "production")
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
