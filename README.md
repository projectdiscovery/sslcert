# sslcert
This library generate a new tlsconfig usable within go standard library configured with a self-signed certificate generated on the fly. Example:

```go
package main

import (
	"log"
	"net/http"

	"github.com/projectdiscovery/sslcert"
)

func main() {
	tlsOptions := sslcert.DefaultOptions
	tlsOptions.Host = "test"
	tlsConfig, err := sslcert.NewTLSConfig(tlsOptions)
	if err != nil {
		log.Fatal(err)
	}
	server := &http.Server{
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatal(err)
	}
}
```