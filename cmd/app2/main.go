package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type App2Request struct {
	JobId *string `json:"jobId"`
}

type App2Response struct {
	CompleteStatus *string `json:"completeStatus"`
}

func App2Handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read request body: %v", err)
		return
	}
	var app2req *App2Request
	err = json.Unmarshal(body, &app2req)
	if err != nil {
		emsg := fmt.Sprintf("failed to unmarshall json: %v", err)
		log.Printf(emsg)
		http.Error(w, emsg, http.StatusBadRequest)
		return
	}

	if app2req.JobId == nil {
		emsg := fmt.Sprintf("jobId not found in request")
		log.Printf(emsg)
		http.Error(w, emsg, http.StatusBadRequest)
		return
	}

	// do the thing
	outstring := fmt.Sprintf("job success")

	// send it back
	app2resp := App2Response{CompleteStatus: &outstring}
	app2respjson, err := json.Marshal(app2resp)
	if err != nil {
		emsg := fmt.Sprintf("failed to marshall response: %v", err)
		log.Printf(emsg)
		http.Error(w, emsg, http.StatusInternalServerError)
		return
	}
	w.Write(app2respjson)
}

func main() {
	app2Port := os.Getenv("APP2_PORT")
	if app2Port == "" {
		app2Port = "8682"
	}
	http.HandleFunc("/", App2Handler)
	listenAddress := fmt.Sprintf(":%s", app2Port)
	log.Printf("Starting app2 at %s", listenAddress)
	http.ListenAndServe(listenAddress, nil)
}
