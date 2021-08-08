package main

import (
	"embed"
	_ "embed"
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

//go:embed assets/*
var staticFS embed.FS

func init() {
	router = gin.Default()

	// de, err := tem.ReadDir("assets")
	// fmt.Println(">>", err)
	// for k, v := range de {

	// 	fs, _ := v.Info()
	// 	fmt.Println(k, v.Name(), fs.Name(), fs.Size())
	// }
	//gin.SetMode(gin.ReleaseMode) //

	//fmt.Println(path.Join("/assets/", "/resources"))

}

func main() {

	//gin.HandlerFunc
	router.LoadHTMLGlob("template/*")

	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(tem))))
	//fs := http.FileServer(tem)
	//router.StaticFS("/resources", http.FS(tem))

	router.GET("/:foldername/*filepath", func(c *gin.Context) {

		fmt.Println("folder:", c.Param("foldername"))

		for key, val := range c.Params {
			fmt.Println(key, "==", val)
		}

		c.FileFromFS(path.Join("/assets/", c.Request.URL.Path), http.FS(staticFS))
	})

	router.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.gohtml", gin.H{"title": "Hello world"})

	})

	router.GET("/about", Aboutpage)

	router.Run(":8080")

}

func Aboutpage(c *gin.Context) {

	data := struct {
		Title string
		Page  string
	}{

		Title: "About us",
		Page:  "About",
	}
	c.HTML(http.StatusOK, "about.gohtml", data)

}
