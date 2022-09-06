package main

import (
	"errors"
	"fmt"
	"time"
	"urlshortener/shortener"

	"urlshortener/api"
	"urlshortener/serializer/json"

	"github.com/mateors/mcb"
)

type Nothing interface{}

type mcbClient struct {
	client *mcb.DB
}

func NewCouchbaseRepository() (shortener.RedirectRepository, error) {

	repo := &mcbClient{}
	db := mcb.Connect("localhost", "mateors", "Test123@", false)
	repo.client = db
	fmt.Println(repo.client.Ping())

	return repo, nil
}

func (t *mcbClient) Find(code string) (*shortener.Redirect, error) {

	redirect := &shortener.Redirect{}

	sql := fmt.Sprintf(`SELECT code,create_time,url FROM bagnbrand.app.test WHERE code="%s";`, code)
	//fmt.Println(sql)
	res := t.client.Query(sql)
	rows := res.GetRows()
	//fmt.Println(len(rows), rows)

	for _, row := range rows {
		dtfloat64 := row["create_time"].(float64)
		//fmt.Println("dtfloat64", dtfloat64)
		//fmt.Printf("%v %T\n", row["create_time"], row["create_time"])
		redirect.Code = fmt.Sprint(row["code"])
		redirect.URL = fmt.Sprint(row["url"])
		redirect.CreateTime = int64(dtfloat64)
	}
	return redirect, nil
}

// Store(redirect *Redirect) error
func (t *mcbClient) Store(redirect *shortener.Redirect) error {

	res := t.client.InsertIntoBucket(redirect.Code, "bagnbrand.app.test", &redirect)
	if res.Status != "success" {
		var msg string
		for _, err := range res.Errors {
			msg += fmt.Sprintf("%v %v\n", err.Code, err.Message)
		}
		return errors.New(msg)
	}
	return nil
}

func main() {

	//Step-1
	nrepo, err := NewCouchbaseRepository()
	fmt.Println(err)

	//Step-2
	//service := shortener.NewRedirectService(nrepo)
	aa := &api.Handler{}

	
	

	var red = &shortener.Redirect{}
	red.Code = "abcd"
	red.URL = "http://mateors.com"
	red.CreateTime = time.Now().Unix()

	bs, err := aa.Serializer("").Encode(red)
	fmt.Println(string(bs), err)

	ss := &json.Redirect{}
	bs, err = ss.Encode(red)
	fmt.Println(string(bs), err)

	//Step - 3
	//err = nrepo.Store(red)
	//fmt.Println(err)

	redObj, err := nrepo.Find("abcd")
	fmt.Println(err, redObj.Code, redObj.CreateTime, redObj.URL)

}
