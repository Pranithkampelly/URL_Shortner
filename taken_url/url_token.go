package taken_url

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	. "url_shortner_copy/Models"
)



func Token(c *gin.Context) {
	token := c.Param("token")
	shortlink := "http://0.0.0.0:8089/new/" + token
	db, err := sql.Open("mysql", "root:pranithkampelly@tcp(127.0.0.1:3306)/url")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var p Link

	err = db.QueryRow("select id,large,short,count_id from links WHERE short =?",shortlink).Scan(&p.Id, &p.Large, &p.Short,&p.Count)

	db.Query( "UPDATE links SET count_id=?  WHERE id=?",p.Count+1,p.Id)

	if(err!=nil) {
		c.HTML(http.StatusOK, "error.html", gin.H{})
	} else {

		c.Redirect(301, p.Large)
	}

}