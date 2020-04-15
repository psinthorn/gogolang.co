package utils

import "os"

type server struct{}

var (
	Server server
)

func (s *server) RunningPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	return port
}
