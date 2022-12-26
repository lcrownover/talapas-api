package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func App1APIHandler(w http.ResponseWriter, r *http.Request) {
	// app1Handler brokers data between the API server and app1
	// validation for input data happens in app1
	log.Print("forwarding request for app1")
	app1Host := os.Getenv("TALAPAS_API_DIRECTOR_APP1_HOST")
	if app1Host == "" {
		app1Host = "localhost"
	}
	app1Port := os.Getenv("TALAPAS_API_DIRECTOR_APP1_PORT")
	if app1Port == "" {
		app1Port = "8681"
	}
	app1URL := fmt.Sprintf("http://%s:%s", app1Host, app1Port)
	res, err := http.Post(app1URL, "application/json", r.Body)
	if err != nil {
		log.Printf("error posting to app1: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("error reading response body: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(body)
}

func App2APIHandler(w http.ResponseWriter, r *http.Request) {
	// app1Handler brokers data between the API server and app1
	// validation for input data happens in app1
	log.Print("forwarding request for app2")
	app2Host := os.Getenv("TALAPAS_API_DIRECTOR_APP2_HOST")
	if app2Host == "" {
		app2Host = "localhost"
	}
	app2Port := os.Getenv("TALAPAS_API_DIRECTOR_APP2_PORT")
	if app2Port == "" {
		app2Port = "8682"
	}
	app2URL := fmt.Sprintf("http://%s:%s", app2Host, app2Port)
	res, err := http.Post(app2URL, "application/json", r.Body)
	if err != nil {
		log.Printf("error posting to app2: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("error reading response body: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(body)
}
