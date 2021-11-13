package config

type HTTPServer struct {
	Host string `env:"HTTP_SERVER_HOST" default:"0.0.0.0"`
	Port int    `env:"HTTP_SERVER_PORT" default:"8080"`
}
