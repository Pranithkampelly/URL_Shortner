package bulk

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"url_shortner_copy/pakages"
)

type link struct {
	short,large,custom string
	id,count   int
}
func Upload(c *gin.Context) {

	file, err := c.FormFile("myFile")

	if err != nil {
		c.String(http.StatusOK, fmt.Sprint("'please upload file"))
	}
	log.Println(file.Filename)

	err = c.SaveUploadedFile(file, "saved/"+file.Filename)
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	file1, err := os.Open("saved/" + file.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	scanner := bufio.NewScanner(file1)

	

	

	for scanner.Scan() {
		fmt.Println("entered")
		link_large := scanner.Text()
		db, err := sql.Open("mysql", "root:pranithkampelly@tcp(127.0.0.1:3306)/url")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		v, _ := db.Query("SELECT short FROM links WHERE  large =?", link_large)

		if (!v.Next()) {
			p, _ := pakages.GenerateRandomString(6)
			shortlink := "http://0.0.0.0:8089/new/" + p

			_, err = db.Query("INSERT INTO links (large,short) VALUES (?,?)", link_large, shortlink)

			// if there is an error inserting, handle it
			if err != nil {
				panic(err.Error())
			}

		} else {
			var p link
			err = db.QueryRow("select id,large,short from links WHERE large =?", link_large).Scan(&p.id, &p.large, &p.short)

		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}