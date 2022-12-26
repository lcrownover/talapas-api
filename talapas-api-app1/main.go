package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type App1Request struct {
	User    *string `json:"user"`
	Command *string `json:"command"`
}

type App1Response struct {
	Data *string `json:"data"`
}

func App1Handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read request body: %v", err)
		return
	}
	log.Printf("received request: %v", string(body))
	var app1req *App1Request
	err = json.Unmarshal(body, &app1req)
	if err != nil {
		emsg := fmt.Sprintf("failed to unmarshall json: %v", err)
		log.Printf(emsg)
		http.Error(w, emsg, http.StatusBadRequest)
		return
	}

	if app1req.User == nil {
		emsg := fmt.Sprintf("user not found in request")
		log.Printf(emsg)
		http.Error(w, emsg, http.StatusBadRequest)
		return
	}
	if app1req.Command == nil {
		emsg := fmt.Sprintf("command not found in request")
		log.Printf(emsg)
		http.Error(w, emsg, http.StatusBadRequest)
		return
	}

	// do the thing
	outstring := fmt.Sprintf("%s tried to run command: %s", *app1req.User, *app1req.Command)

	// send it back
	app1resp := App1Response{Data: &outstring}
	app1respjson, err := json.Marshal(app1resp)
	if err != nil {
		emsg := fmt.Sprintf("failed to marshall response: %v", err)
		log.Printf(emsg)
		http.Error(w, emsg, http.StatusInternalServerError)
		return
	}
	w.Write(app1respjson)
}

func main() {
	app1Port := os.Getenv("TALAPAS_API_APP1_PORT")
	if app1Port == "" {
		app1Port = "8681"
	}
	http.HandleFunc("/", App1Handler)
	listenAddress := fmt.Sprintf(":%s", app1Port)
	log.Printf("Starting app1 at %s", listenAddress)
	http.ListenAndServe(listenAddress, nil)
}
