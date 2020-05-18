package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	dest := os.Getenv("DEST")
	if dest == "" {
		// original destination, here for easy deploy to original app
		// that handles golangbridge.org
		dest = "https://github.com/gobridge/about-us/blob/master/README.md"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println(
		http.ListenAndServe(
			":"+port,
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					http.Redirect(w, r, dest, http.StatusMovedPermanently)
				},
			),
		),
	)
}
