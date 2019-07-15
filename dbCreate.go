package main


import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db, err := sql.Open("mysql", "root:pranithkampelly@tcp(127.0.0.1:3306)/url")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()


	_, err = db.Query("CREATE TABLE links (id integer AUTO_INCREMENT ,short VARCHAR(255) UNIQUE , large VARCHAR(2000),count_id integer DEFAULT 0 ,PRIMARY KEY(id))")

	if err != nil {
		panic(err.Error())
	}

}