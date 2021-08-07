package main

import (
	_ "embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed students.txt
//var f string
var router *gin.Engine

func init() {
	router = gin.Default()
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

	router.Run(":8080")

}
