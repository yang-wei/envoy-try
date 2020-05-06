package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	serviceName := os.Getenv("SERVICE_NAME")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log := fmt.Sprintf("route: %s, serviceName: %s", r.URL.Path, serviceName)
		fmt.Println(log)
		fmt.Fprint(w, log)
	})
	fmt.Println("Listen on port 9001")
	http.ListenAndServe(":9001", nil)
}
