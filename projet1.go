package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	. "url_shortner_copy/custom_new"
	. "url_shortner_copy/post"
	. "url_shortner_copy/taken_url"
)

func getting(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome-template.html", gin.H{})
}
func custom(c *gin.Context) {
	c.HTML(http.StatusOK, "custom.html", gin.H{})
}
func home(c *gin.Context) {
	c.HTML(http.StatusOK, "root.html", gin.H{})
}

func main() {

		router := gin.Default() //Initiating gin framework
		router.LoadHTMLGlob("template/*") //this is direct html request to templete folder
		router.GET("/", home)
		router.GET("/short", getting)// Displays the form to generate short url  and a link to custom url generation
		router.GET("/custom", custom)
		router.POST("/new", Posting)
		router.POST("/custom/new", Custom_new)
	    router.GET("/new/:token", Token)
	    router.Run(":8089")

	}