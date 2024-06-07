package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment    string
	SMTPHost       string
	SMTPPort       int
	SMTPUsername   string
	SMTPPassword   string
	SMSAPIKey      string
	PushAPIKey     string
	JobQueueSize   int
	JobWorkerCount int
	ServerAddress  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	return &Config{
		Environment:    getEnv("ENVIRONMENT", "development"),
		SMTPHost:       getEnv("SMTP_HOST", "smtp.example.com"),
		SMTPPort:       getEnvAsInt("SMTP_PORT", 587),
		SMTPUsername:   getEnv("SMTP_USERNAME", "user@example.com"),
		SMTPPassword:   getEnv("SMTP_PASSWORD", "password"),
		SMSAPIKey:      getEnv("SMS_API_KEY", "your-sms-api-key"),
		PushAPIKey:     getEnv("PUSH_API_KEY", "your-push-api-key"),
		JobQueueSize:   getEnvAsInt("JOB_QUEUE_SIZE", 10),
		JobWorkerCount: getEnvAsInt("JOB_WORKER_COUNT", 5),
		ServerAddress:  getEnv("SERVER_ADDRESS", "0.0.0.0:8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return defaultValue
		}
		return value
	}
	return defaultValue
}
