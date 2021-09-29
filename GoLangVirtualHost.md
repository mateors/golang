# How to provide the virtual hosts functionality using golang

```golang

package main

import(
    "net/url"
    "net/http"
    "net/http/httputil"
)

func main() {
    vhost1, err := url.Parse("http://127.0.0.1:1980")
    if err != nil {
            panic(err)
    }
    proxy1 := httputil.NewSingleHostReverseProxy(vhost1)
    http.HandleFunc("publicdomain1.com/", handler(proxy1))

    vhost2, err := url.Parse("http://127.0.0.1:1981")
    if err != nil {
            panic(err)
    }
    proxy2 := httputil.NewSingleHostReverseProxy(vhost2)
    http.HandleFunc("publicdomain2.com/", handler(proxy2))

    err = http.ListenAndServe(":80", nil)
    if err != nil {
          panic(err)
    }
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
            p.ServeHTTP(w, r)
    }
}

```

## Resource
* https://stackoverflow.com/questions/14170799/how-to-get-virtualhost-functionality-in-go
