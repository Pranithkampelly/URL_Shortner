package main


import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
	"time"
)
type Welcome struct {
	Name string
	Time string
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

func main() {

	welcome := Welcome{"PranithKampelly", time.Now().Format(time.Stamp)}


	templates := template.Must(template.ParseFiles("template/welcome-template.html"))

	http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {

		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/new" , func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		largelink := r.Form.Get("largelink")


		db, err := sql.Open("mysql", "root:pranithkampelly@tcp(127.0.0.1:3306)/url")

		// if there is an error opening the connection, handle it
		if err != nil {
			panic(err.Error())
		}

		// defer the close till after the main function has finished
		// executing


		defer db.Close()



		v, _ := db.Query("SELECT * FROM links WHERE  large =?",largelink )

		if(!v.Next()) {
			p,_ := GenerateRandomString(6)
			shortlink := "http://0.0.0.0:8080/"+p

			_, err = db.Query("INSERT INTO links (large,short) VALUES (?,?)", largelink, shortlink)

			// if there is an error inserting, handle it
			if err != nil {
				panic(err.Error())
			}

			template2 := template.Must(template.ParseFiles("template/url_display.html"))
			data := struct {
				AppVersion string
			}{
				AppVersion: shortlink ,
			}
			if err := template2.ExecuteTemplate(w, "url_display.html",data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else {
			template1 := template.Must(template.ParseFiles("template/display.html"))
			if err := template1.ExecuteTemplate(w, "display.html", welcome); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		// be careful deferring Queries if you are using transactions


	})
	fmt.Println("Listening");
	fmt.Println(http.ListenAndServe(":8080", nil));
}