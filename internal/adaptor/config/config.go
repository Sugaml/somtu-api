package config

type Config struct {
	APP_NAME  string `json:"APP_NAME" default:"parking"`
	APP_ENV   string `json:"APP_ENV" default:"development"`
	APP_DEBUG string `json:"APP_DEBUG" default:"true"`
	APP_PORT  string `json:"APP_PORT" default:"8080"`

	JWT_SECRET string `json:"JWT_SECRET" default:"postgres"`

	DB_DRIVER       string `json:"DB_DRIVER" default:"postgres"`
	DB_SOURCE       string `json:"DB_SOURCE" default:"postgresql://postgres:postgres@localhost:5432/parking?sslmode=disable"`
	DB_DEBUG        string `json:"DB_DEBUG" default:"false"`
	DB_AUTO_MIGRATE string `json:"DB_AUTO_MIGRATE" default:"true"`

	HOST_PATH  string `json:"HOST_PATH" default:"localhost:8080"`
	SENTRY_DSN string `json:"SENTRY_DSN" default:"https://2215c53806bf33a997bc1c451b0fb482@o4508324654415872.ingest.us.sentry.io/4508324720082944"`

	REDIS_ADDR     string `json:"REDIS_ADDR" default:"localhost:6379"`
	REDIS_PASSWORD string `json:"REDIS_PASSWORD" default:""`

	NOTIFICATION_GRPC_URL string `json:"NOTIFICATION_GRPC_URL" default:"localhost:50050"`
}
