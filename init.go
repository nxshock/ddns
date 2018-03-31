package main

import (
	"flag"
	"net/http"
)

func init() {
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
