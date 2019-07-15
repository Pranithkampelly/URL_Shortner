package taken_url

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)
type link struct {
	short,large,custom string
	id,count   int
}


func Token(c *gin.Context) {
	token := c.Param("token")
	shortlink := "http://0.0.0.0:8089/new/" + token
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