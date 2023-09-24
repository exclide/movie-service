package store

type Config struct {
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		DatabaseURL: "postgres://localhost:5436/movieserv?sslmode=disable&user=postgres&password=123",
	}
}
