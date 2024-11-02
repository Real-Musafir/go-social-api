package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request){

	fmt.Print("Come here to run")
	w.Write([]byte("OKKK OKKKK"))
}