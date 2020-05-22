package configs

import "os"

type serverPort struct{}
type serverPortSelectorInterface interface {
	PortSelector(string) string
}

var ServerPort serverPortSelectorInterface = &serverPort{}

func (p *serverPort) PortSelector(port string) string {
	portEnv := os.Getenv("PORT")
	if portEnv == "" {
		portEnv = port
	}
	return portEnv
}
