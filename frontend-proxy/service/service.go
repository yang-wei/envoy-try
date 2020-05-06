package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	serviceName := os.Getenv("SERVICE_NAME")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("route: %s, serviceName: %s", r.URL.Path, serviceName))
	})

	http.ListenAndServe(":8001", nil)
}
