package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	yaml "gopkg.in/yaml.v2"
)

type User struct {
	First_name string
	Last_name  string
	ID         string
	Grade      string
}

type Config struct {
	Data map[string]string
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

	r.GET("/:queryName", func(c *gin.Context) {
		queryName := c.Param("queryName")
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

		// this reads yml file

		var conf Config
		reader, err := os.Open("test.yml")

		if err != nil {
			log.Fatal(err)
		}

		buf, err := ioutil.ReadAll(reader)

		if err != nil {
			log.Fatal(err)
		}

		err = yaml.Unmarshal(buf, &conf)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", conf)
		if _, ok := conf.Data[queryName]; !ok {
			c.Data(404, "text/plain", []byte("Query not found: "+queryName))
			return
		}
		fmt.Println(conf.Data)

		query := conf.Data[queryName]

		rows, err := db.Query(query)

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
