package main

import (
	"fmt"
	"os"
    "log"
    "net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, os.Getenv("HOSTNAME"))
	})

	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("envvar")
		log.Print(param)
		if param != "" {
			fmt.Fprintf(w, os.Getenv(param))
		}
	})

	if err := http.ListenAndServe("-1", nil); err != nil {
		log.Fatal(err)
	}
}
