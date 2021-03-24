package constants

import "os"

// Environment ...
type Environment struct {
	JwtString  string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSslmode  string
	DBHost     string
}

// GetEnvironment ...
func GetEnvironment() Environment {
	return Environment{
		JwtString:  os.Getenv("JWT_STRING"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		DBSslmode:  os.Getenv("DB_SSL_MODE"),
		DBHost:     os.Getenv("DB_HOST"),
	}
}

// Non Environment variables
