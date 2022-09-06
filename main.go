package main

import (
	"fmt"
	"urlshortener/database/couchbase"
	"urlshortener/shortener"
)

type Nothing interface{}

func main() {

	//Step-1
	nrepo, err := couchbase.NewCouchbaseRepository()
	//fmt.Println(err)

	//Step-2
	service := shortener.NewRedirectService(nrepo)
	redr, err := service.Find("abcd")

	fmt.Println(redr, err)

	// var red = &shortener.Redirect{}
	// red.Code = "abcd"
	// red.URL = "http://mateors.com"
	// red.CreateTime = time.Now().Unix()

	//aa := &api.Handler{}
	// bs, err := aa.Serializer("").Encode(red)
	// fmt.Println(string(bs), err)
	//redObj1, err1 := aa.RedirectService.Find("abcd")
	//fmt.Println(redObj1, err1)

	// ss := &json.Redirect{}
	// bs, err := ss.Encode(red)
	// fmt.Println(string(bs), err)

	//Step - 3
	//err = nrepo.Store(red)
	//fmt.Println(err)

}
