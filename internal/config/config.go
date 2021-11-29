package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Environment string

	PGHost     string
	PGPort     int
	PGUser     string
	PGPassword string
	PGDatabase string
	PGSSL      string
}

func Set() Config {
	if err := godotenv.Load(); err != nil {
		panic(".env file not found")
	}

	conf := Config{}

	conf.Environment = cast.ToString(key("ENVIRONMENT"))

	conf.PGHost = cast.ToString(key("PG_HOST"))
	conf.PGPort = cast.ToInt(key("PG_PORT"))
	conf.PGUser = cast.ToString(key("PG_USER"))
	conf.PGPassword = cast.ToString(key("PG_PASSWORD"))
	conf.PGDatabase = cast.ToString(key("PG_DATABASE"))
	conf.PGSSL = cast.ToString(key("PG_SSL"))

	return conf
}

func (cfg *Config) PostgreSQLURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.PGUser,
		cfg.PGPassword,
		cfg.PGHost,
		cfg.PGPort,
		cfg.PGDatabase,
		cfg.PGSSL,
	)
}

func key(key string) interface{} {
	_, ok := os.LookupEnv(key)
	if !ok {
		log.Panicf("[%s] is key not found", key)
	}

	return os.Getenv(key)
}
