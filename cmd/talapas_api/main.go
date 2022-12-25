package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lcrownover/talapas-api/internal/talapas_api"
)

func main() {
	apiPort := os.Getenv("TALAPAS_API_PORT")
	if apiPort == "" {
		apiPort = "8080"
	}
	listenAddress := fmt.Sprintf(":%s", apiPort)

	// app1 takes a user and a command and returns data:
	// '{"data": "{user} tried to run command: {command}"}'
	http.HandleFunc("/app1", talapas_api.App1APIHandler)
	http.HandleFunc("/app2", talapas_api.App2APIHandler)

	log.Printf("Starting API at %s", listenAddress)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("error starting server:", err)
	}
}
