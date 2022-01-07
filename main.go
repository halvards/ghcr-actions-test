package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/halvards/ghcr-actions-test/pkg/version"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println(version.Version)
		return
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, version.Version)
	})
	fmt.Printf("Listening on port %v\n", port)
	http.ListenAndServe(":"+port, nil)
}
