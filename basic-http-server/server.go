package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func apiHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World!")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "ok")
}

func setupHandlers(mux *http.ServeMux){
	mux.HandleFunc("/api", apiHandler)
	mux.HandleFunc("/healthz", healthCheckHandler)
}

func main(){
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}

	mux := http.NewServeMux()
	setupHandlers(mux)

	log.Fatal(http.ListenAndServe(listenAddr, mux))  // The ListenAndServe() function in the net/http package starts an HTTP server at a given network address. It is a good idea to make this address configurable.
}

/*
	In the main() function, the following
	lines check if the LISTEN_ADDR environment variable has been
	specified, and if not, it defaults to ":8080"


	The Getenv() function defined in the os package looks for the value of
	an environment variable. If an environment variable LISTEN_ADDR is
	found, its value is returned as a string. If no such environment
	variable exists, an empty string is returned.


	The default value of ":8080" means that the
	server will listen on all network interfaces on the port 8080. If you
	wanted the server only to be reachable on the computer where your
	application is running, you would set the environment variable
	LISTEN_ADDR to "127.0.0.1:8080" and then start the application.


	Next, we call the ListenAndServe() function specifying the address to
	listen on ( listenAddr ) and the handler for the server. We will
	specify nil as the value for the handler and thus our function call to
	ListenAndServe is as follows:
	log.Fatal(http.ListenAndServe(listenAddr, nil))


*/