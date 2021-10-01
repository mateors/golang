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

## Resource
* https://posener.github.io/http2
