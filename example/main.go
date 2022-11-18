package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/projectdiscovery/sslcert"
)

func main() {
	tlsOptions := sslcert.DefaultOptions
	tlsOptions.Host = "example.com"

	// Create TLSConfig using options
	tlsConfig, err := sslcert.NewTLSConfig(tlsOptions)
	if err != nil {
		log.Fatal(err)
	}

	// using tlsconfig to host http server
	server := &http.Server{
		Addr:      ":8000",
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, er := w.Write([]byte("Hello World"))
		if er != nil {
			log.Print(er)
		}
	})

	fmt.Println("Started HTTPS server on " + server.Addr)
	fmt.Println("Check it out at https://localhost:8000/")
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatal(err)
	}
}
