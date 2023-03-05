package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
}

func main() {
	app := Config{}

	log.Printf("Starting Brocker Server on port %s\n", webPort)

	// todo : define http server

	svr := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start server

	err := svr.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
