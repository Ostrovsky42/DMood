package config

import "os"

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type BotConfig struct {
	TgToken string
	PngPath string
}

type Config struct {
    DBConfig
	BotConfig
}

func NewConfig() *Config {
	return &Config{
		DBConfig: DBConfig{
			Host: getEnv("DB_HOST", ""),
			Port: getEnv("DB_PORT", ""),
			User: getEnv("DB_USER_NAME", ""),
			Password: getEnv("DB_PASSWORD", ""),
			DBName: getEnv("DB_NAME", ""),
		},
		BotConfig: BotConfig{
		TgToken: getEnv("TG_TOKEN",""),
		PngPath: getEnv("PNG_PATH",""),
		},
	}
}


func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}