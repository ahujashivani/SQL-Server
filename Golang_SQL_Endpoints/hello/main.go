package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	First_name string
	Last_name  string
	ID         string
	Grade      string
}

func main() {

	//var connection_string = os.Getenv("MYSQL_CONNECTION")
	var port = ":" + os.Getenv("DB_PORT")
	var user = os.Getenv("DB_USER")
	var pass = os.Getenv("DB_PASSWORD")
	var host = os.Getenv("DB_HOST")
	var db = os.Getenv("DB_DATABASE")
	connection_string := user + ":" + pass + "@tcp(" + host + ")/" + db

	r := gin.Default()
	r.LoadHTMLGlob("*.html")

	r.GET("/", func(c *gin.Context) {
		db, err := sql.Open("mysql", connection_string)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Connection Established")
		}

		var first_name string
		var last_name string
		var id string
		var grade string
		var users []User

		rows, err := db.Query("SELECT * FROM authors")

		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			err := rows.Scan(&first_name, &last_name, &id, &grade)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, User{First_name: first_name, Last_name: last_name, ID: id, Grade: grade})
		}

		fmt.Println(user)
		fmt.Println(pass)
		fmt.Println(host)
		fmt.Println(db)
		fmt.Println(port)
		fmt.Println(connection_string)

		fmt.Println(users)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
			"query": users,
		})

		//c.JSON(http.StatusOK, users)

		defer db.Close()

	})

	r.Run(port)
}
