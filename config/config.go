package config

type Config struct {
	HTTPPort string

	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     string
	PostgresMaxConn  int32
}

func Load() Config {

	cfg := Config{}

	cfg.HTTPPort = ":8080"

	cfg.PostgresHost = "locahost"
	cfg.PostgresMaxConn = 30

	return cfg
}
