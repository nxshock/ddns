package main

import (
	"flag"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(0)

	initFlags()
	initHTTPClient()
}

func initFlags() {
	flag.StringVar(&flagWorkingDir, "d", defaultWorkingDir, "path to config directory")
	flag.DurationVar(&flagHTTPTimeout, "t", defaultHTTPTimeout, "http request timeout")
	flag.Parse()
}

func initHTTPClient() {
	httpClient = &http.Client{Timeout: flagHTTPTimeout}
}
