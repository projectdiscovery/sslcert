# sslcert

[![License](https://img.shields.io/github/license/projectdiscovery/sslcert)](LICENSE.md)
![Go version](https://img.shields.io/github/go-mod/go-version/projectdiscovery/sslcert?filename=go.mod)
[![Release](https://img.shields.io/github/release/projectdiscovery/sslcert)](https://github.com/projectdiscovery/sslcert/releases/)
[![Checks](https://github.com/projectdiscovery/sslcert/actions/workflows/build-test.yml/badge.svg)](https://github.com/projectdiscovery/sslcert/actions/workflows/build-test.yml)
[![GoDoc](https://pkg.go.dev/badge/projectdiscovery/sslcert)](https://pkg.go.dev/github.com/projectdiscovery/sslcert)



sslcert library generates a new tlsconfig usable within go standard library configured with a self-signed certificate generated on the fly. 


# Example

```go
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
		TLSConfig: tlsConfig,s
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	fmt.Println("Started HTTPS server on " + server.Addr)
	fmt.Println("Check it out at https://localhost:8000/")
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatal(err)
	}
}
```