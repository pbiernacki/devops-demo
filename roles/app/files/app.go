package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type timeHandler struct{}

func (t timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's %v here\n", time.Now())
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		h, _ := os.Hostname()

		fmt.Fprintf(w, "Hi there, I'm served from %s!\n", h)
	})

	http.Handle("/time", timeHandler{})

	log.Fatal(http.ListenAndServe(":8484", nil))
}
