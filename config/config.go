package config

import (
	"fmt"
	"os"
	"strconv"
)

type (
	Config struct {
		App        App
		MasterDB   DB
	}

	App struct {
		Name         string
		Env          string
		Port         int
		DefaultLang  string
		ReadTimeout  int
		WriteTimeout int
	}

	DB struct {
		Host     string
		Port     int
		User     string
		Password string
		DB       string
	}
)

func New() *Config {
	return &Config{
		App: App{
			Name:         getEnv("APP_NAME", "go-boilerplate"),
			Env:          getEnv("APP_ENV", "development"),
			Port:         getEnvAsInt("APP_PORT", 3000),
			DefaultLang:  getEnv("APP_DEFAULT_LANG", "en"),
			ReadTimeout:  getEnvAsInt("APP_READ_TIMEOUT", 10),
			WriteTimeout: getEnvAsInt("APP_WRITE_TIMEOUT", 10),
		},
		MasterDB: DB{
			Host:     getEnv("POSTGRES_HOST_MASTER", "localhost"),
			Port:     getEnvAsInt("POSTGRES_PORT_MASTER", 5432),
			User:     getEnv("POSTGRES_USER_MASTER", "postgres"),
			Password: getEnv("POSTGRES_PASSWORD_MASTER", "postgres"),
			DB:       getEnv("POSTGRES_DB_MASTER", "booking_to_go_v1"),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}

	if nextValue := os.Getenv(key); nextValue != "" {
		return nextValue
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func (c *Config) DatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.MasterDB.User,
		c.MasterDB.Password,
		c.MasterDB.Host,
		c.MasterDB.Port,
		c.MasterDB.DB,
	)
}