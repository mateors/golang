package main

import (
	"embed"
	_ "embed"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

//go:embed template
var tem embed.FS

func init() {
	router = gin.Default()
	fmt.Println(">>", tem)
	//gin.SetMode(gin.ReleaseMode) //
}

func main() {

	//var f embed.FS
	//f.ReadDir()
	//fmt.Println(f)

	//gin.HandlerFunc
	router.LoadHTMLGlob("template/*")

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
