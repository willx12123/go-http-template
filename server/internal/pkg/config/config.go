package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DB       string
}

type HTTPConfig struct {
	Port int
}

type AdminConfig struct {
	Token string
}

type AppConfig struct {
	Admin    *AdminConfig
	HTTP     *HTTPConfig
	Database *DatabaseConfig
}

const (
	goEnvKey = "GO_ENV"
	envDev   = "development"
	envProd  = "production"
)

var (
	goEnv string

	defaultConf *AppConfig
)

func Default() *AppConfig {
	return defaultConf
}

func IsProd() bool {
	return goEnv == envProd
}

func Init() {
	loadEnv()

	var databaseConfig DatabaseConfig
	envconfig.MustProcess("database", &databaseConfig)

	var httpConfig HTTPConfig
	envconfig.MustProcess("http", &httpConfig)

	var adminConfig AdminConfig
	envconfig.MustProcess("admin", &adminConfig)

	defaultConf = &AppConfig{
		HTTP:     &httpConfig,
		Database: &databaseConfig,
		Admin:    &adminConfig,
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	goEnv = os.Getenv(goEnvKey)
	if goEnv == "" {
		goEnv = envDev
	}
	err = godotenv.Load(fmt.Sprintf(".env.%s", goEnv))
	if err != nil {
		panic(err)
	}
}
