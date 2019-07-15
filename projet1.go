package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	. "url_shortner_copy/bulk"
	. "url_shortner_copy/custom_new"
	. "url_shortner_copy/html"
	. "url_shortner_copy/post"
	. "url_shortner_copy/taken_url"
)


func main() {

		router := gin.Default() //Initiating gin framework
		router.LoadHTMLGlob("template/*") //this is direct html request to templete folder
		router.GET("/", Home)
		router.GET("/short", Getting) // Displays the form to generate short url  and a link to custom url generation
		router.GET("/custom", Custom)
		router.POST("/new", Posting)
		router.POST("/custom/new", Custom_new)
	    router.GET("/new/:token", Token)
		router.GET("/bulk", Bulkupload)
		router.POST("/bulk/new", Upload)
	    router.Run(":8089")

	}