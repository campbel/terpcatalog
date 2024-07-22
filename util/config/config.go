package config

import (
	"os"
	"strings"

	"github.com/campbel/terpcatalog/shared/types"
)

func CatalogPort() string {
	return getEnv("CATALOG_PORT", "8080")
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

func MongoConnectionString() string {
	return getEnv("MONGO_CONNECTION_STRING", "mongodb://localhost:27017")
}

func EmailToken() string {
	return getEnv("EMAIL_TOKEN", "")
}

func EmailSender() string {
	return getEnv("EMAIL_SENDER", "auto@terpcatalog.com")
}

func EmailOrderList() []types.EmailAddress {
	var list []types.EmailAddress
	emails := strings.Split(getEnv("EMAIL_ORDER_LIST", "Chris Campbell:campbel@hey.com,Scott Campbell:scott@terpscout.com"), ",")
	for _, email := range emails {
		parts := strings.Split(email, ":")
		list = append(list, types.EmailAddress{Name: parts[0], Email: parts[1]})
	}
	return list
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
