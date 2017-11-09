package main

import (
	"flag"
	"os"

	"github.com/go-kit/kit/log"
)

const (
	defaultPort = "80"
)

func main() {
	var (
		port     = envString("PORT", defaultPort)
		httpAddr = flag.String("http.addr", ":"+port, "HTTP listen Ports")
	)
	flag.Parse()
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	}

}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
