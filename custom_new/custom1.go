package custom_new

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
   ."url_shortner_copy/Models"
)

func Custom_new(c *gin.Context) {
	largelink:= c.PostForm("largelink")
	db, err := sql.Open("mysql", "root:pranithkampelly@tcp(127.0.0.1:3306)/url")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	v, _ := db.Query("SELECT short FROM links WHERE  large =?",largelink )
	if(!v.Next()) {
		//p,_ := GenerateRandomString(6)
		shortlink := "http://0.0.0.0:8089/new/"+c.PostForm("customlink")
		//customlink := c.PostForm("customlink")

		_, err = db.Query("INSERT INTO links (large,short) VALUES (?,?)", largelink, shortlink)
		//_, err = db.Query("INSERT INTO custom_links (large,short,custom) VALUES (?,?,?)", largelink, shortlink,customlink)


		// if there is an error inserting, handle it
		if err != nil {
			c.HTML(http.StatusOK, "error1.html", gin.H{})

			//panic(err.Error())

		} else {
			c.HTML(http.StatusOK, "url_display.html", gin.H{"shortlink": shortlink})
		}

	} else {
		var p Link
		err = db.QueryRow("select id,large,short from links WHERE large =?",largelink).Scan(&p.Id, &p.Large, &p.Short)

		c.HTML(http.StatusOK, "display.html", gin.H{"url":p.Short})
	}
}
