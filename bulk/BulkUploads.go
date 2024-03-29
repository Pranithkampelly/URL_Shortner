package bulk

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"sync"
	. "url_shortner_copy/Models"
	"url_shortner_copy/pakages"
)
var wg sync.WaitGroup


  func scanning(scaning string ){

	  link_large := scaning
	  fmt.Println(link_large)
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
		  var p Link
		  err = db.QueryRow("select id,large,short from links WHERE large =?", link_large).Scan(&p.Id, &p.Large, &p.Short)

	  }
	  wg.Done()
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
		  wg.Add(1)
	  	go scanning(scanner.Text())

	  }
	  if err := scanner.Err(); err != nil {
		  log.Fatal(err)
	  }
	 wg.Wait()
  }