package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)
type link struct {
	short,large,custom string
	id,count   int
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func getting(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome-template.html", gin.H{})
}
func posting(c *gin.Context) {
	largelink:= c.PostForm("largelink")
	db, err := sql.Open("mysql", "root:pranithkampelly@tcp(127.0.0.1:3306)/url")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	v, _ := db.Query("SELECT short FROM links WHERE  large =?",largelink )


	if(!v.Next()) {
		p,_ := GenerateRandomString(6)
		shortlink := "http://0.0.0.0:8080/new/"+p

		_, err = db.Query("INSERT INTO links (large,short) VALUES (?,?)", largelink, shortlink)

		// if there is an error inserting, handle it
		if err != nil {
			panic(err.Error())
		}
		c.HTML(http.StatusOK, "url_display.html", gin.H{"shortlink":shortlink})

	} else {
		var p link
		err = db.QueryRow("select id,large,short from links WHERE large =?",largelink).Scan(&p.id, &p.large, &p.short)

		c.HTML(http.StatusOK, "display.html", gin.H{"url":p.short})
	}
}

func token(c *gin.Context) {
	token := c.Param("token")
	shortlink := "http://0.0.0.0:8080/new/" + token
	db, err := sql.Open("mysql", "root:pranithkampelly@tcp(127.0.0.1:3306)/url")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var p link
	err = db.QueryRow("select id,large,short,count_id from links WHERE short =?",shortlink).Scan(&p.id, &p.large, &p.short,&p.count)

	db.Query( "UPDATE links SET count_id=?  WHERE id=?",p.count+1,p.id)

	if(err!=nil) {
	 	c.HTML(http.StatusOK, "error.html", gin.H{})
	 } else {

		 c.Redirect(301, p.large)
	 }

}

func custom(c *gin.Context) {
	c.HTML(http.StatusOK, "custom.html", gin.H{})
}
func custom_new(c *gin.Context) {
	largelink:= c.PostForm("largelink")
	db, err := sql.Open("mysql", "root:pranithkampelly@tcp(127.0.0.1:3306)/url")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	v, _ := db.Query("SELECT short FROM links WHERE  large =?",largelink )
	if(!v.Next()) {
		//p,_ := GenerateRandomString(6)
		shortlink := "http://0.0.0.0:8080/new/"+c.PostForm("customlink")
		//customlink := c.PostForm("customlink")

		_, err = db.Query("INSERT INTO links (large,short) VALUES (?,?)", largelink, shortlink)
		//_, err = db.Query("INSERT INTO custom_links (large,short,custom) VALUES (?,?,?)", largelink, shortlink,customlink)


		// if there is an error inserting, handle it
		if err != nil {
			panic(err.Error())
		}
		c.HTML(http.StatusOK, "url_display.html", gin.H{"shortlink":shortlink})

	} else {
		var p link
		err = db.QueryRow("select id,large,short from links WHERE large =?",largelink).Scan(&p.id, &p.large, &p.short)

		c.HTML(http.StatusOK, "display.html", gin.H{"url":p.short})
	}
}
func main() {

		router := gin.Default() //Initiating gin framework
		router.LoadHTMLGlob("template/*") //this is direct html request to templete folder
		router.GET("/", getting)// Displays the
		router.GET("/custom", custom)
		router.POST("/new",posting)
		router.POST("/custom/new", custom_new)
	    router.GET("/new/:token",token )
	    router.Run(":8080")


	}