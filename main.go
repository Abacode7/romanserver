package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Abacode7/romannumerals"
)

func main() {

	// Define paths and functions
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		paths := strings.Split(r.URL.Path, "/")

		if paths[1] == "roman_numerals" {
			argVal, err := strconv.Atoi(paths[2])
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("400 - Bad request. Pls check your url path."))
			} 
			

			if argVal < 1 || argVal > 10{
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not found. Resource for the given argument not found."))
			}else{
				fmt.Fprintf(w, "%q", html.EscapeString(romannumerals.Numerals[argVal]))
			}

		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad request. Pls check your url path."))
		}
	})

	// Create server object
	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Start server
	fmt.Println("Starting server...")
	server.ListenAndServe()
}
