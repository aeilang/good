package config

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Environment          string        `validate:"required" koanf:"ENVIRONMENT"`
	HTTPServerAddress    string        `validate:"required" koanf:"HTTP_SERVER_ADDRESS"`
	HTTPTimeOut          time.Duration `validate:"required" koanf:"HTTP_TIME_OUT"`
	LogLevel             string        `validate:"required" koanf:"LOG_LEVEL"`
	LogFileName          string        `koanf:"LOG_FILE_NAME"`
	LogMaxSize           int           `koanf:"LOG_MAX_SIZE"`
	LogMaxBackups        int           `koanf:"LOG_MAX_BACKUPS"`
	LogMaxAge            int           `koanf:"LOG_MAX_AGE"`
	LogCompress          bool          `koanf:"LOG_COMPRESS"`
	DBName               string        `validate:"required" koanf:"POSTGRES_DB"`
	DBUser               string        `validate:"required" koanf:"POSTGRES_USER"`
	DBPassword           string        `validate:"required" koanf:"POSTGRES_PASSWORD"`
	DBSource             string        `validate:"required" koanf:"DB_SOURCE"`
	SecretKey            string        `validate:"required" koanf:"SECRET_KEY"`
	AccessTokenDuration  time.Duration `validate:"required" koanf:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `validate:"required" koanf:"REFRESH_TOKEN_DURATION"`
	ESAddress            string        `validate:"required" koanf:"ES_ADDRESS"`
	ESIndex              string        `validate:"required" koanf:"ES_INDEX"`
	CorsOrigin           string        `validate:"required" koanf:"CORS_ORIGIN"`
	TOKEN                string        `validate:"required" koanf:"TOKEN"`
	MasterKey            string        `validate:"required" koanf:"MASTER_KEY"`
}

var configFile string = ".env.dev"

var (
	k        *koanf.Koanf
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {

	once.Do(func() {
		environment := os.Getenv("ENVIRONMENT")
		if environment == "production" {
			configFile = ".env"
		}

		k = koanf.New(".")
		validate := validator.New(validator.WithRequiredStructEnabled())

		log.Println("loading config...")

		fileProvider := file.Provider(configFile)
		envProvider := env.Provider("", ".", nil)

		if err := k.Load(fileProvider, dotenv.Parser()); err != nil {
			log.Printf("could not load config file: %s", err.Error())
		}

		if err := k.Load(envProvider, nil); err != nil {
			log.Printf("could not environment variables: %s", err.Error())
		}

		if err := k.Unmarshal("", &instance); err != nil {
			log.Panicf("error unmarshing config %v", err)
		}

		if err := validate.Struct(instance); err != nil {
			log.Panicf("correct configs were not loaded %v", err)
		}

	})

	return instance
}
