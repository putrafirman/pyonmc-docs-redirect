package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define the HTTP port
	httpPort := ":80"

	// Define the target domain to redirect to
	targetDomain := "https://docs.pyonmc.my.id"

	// HTTP redirection handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logged IP [" + r.RemoteAddr + "] == User : " + r.UserAgent())
		http.Redirect(w, r, targetDomain+r.RequestURI, http.StatusTemporaryRedirect)
	})

	// Start the HTTP server
	fmt.Printf("HTTP server listening on %s\n", httpPort)
	err := http.ListenAndServe(httpPort, nil)
	if err != nil {
		fmt.Printf("HTTP server error: %v\n", err)
	}
}
