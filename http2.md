# Golang HTTP2
Go’s standard library HTTP server supports HTTP/2 by default \
Most users will use it indirectly through the automatic use by the net/http package (from Go 1.6 and later)

[HTTP2 Demo code](https://github.com/golang/net/blob/master/http2/h2demo/h2demo.go)

Everything is automatically configured for us, we don’t even need to import Go’s standard library http2 package \

HTTP/2 enforces TLS. In order to achieve this we first need a private key and a certificate. On Linux, the following command does the job. Run it and follow the prompted questions.

```
openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt
```
The command will generate two files: server.key and server.crt

* server.key: Contains our server private key
* server.crt: The server certificate - represents the server’s identity and contains the server’s public key

Now, for the server code, in its simplest form, we will just use Go’s standard library HTTP server and enable TLS with the generated SSL files.
```golang
package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a server on port 8000
	// Exactly how you would run an HTTP/1.1 server
	srv := &http.Server{Addr: ":8000", Handler: http.HandlerFunc(handle)}

	// Start the server with TLS, since we are running HTTP/2 it must be
	// run with TLS.
	// Exactly how you would run an HTTP/1.1 server with TLS connection.
	log.Printf("Serving on https://0.0.0.0:8000")
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}

func handle(w http.ResponseWriter, r *http.Request) {
	// Log the request protocol
	log.Printf("Got connection: %s", r.Proto)
	// Send a message back to the client
	w.Write([]byte("Hello"))
}
```

## Resource
* https://posener.github.io/http2
* https://github.com/posener/h2conn
