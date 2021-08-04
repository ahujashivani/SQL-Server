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

// create structs that will be used
type User struct {
	First_name string
	Last_name  string
	ID         string
	Grade      string
}

type Config struct {
	Data map[string]string
}

// helper function to keep code dry
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Establish local variables from the environment variables found in docker-compose
func establishConnection() string {
	var user = os.Getenv("DB_USER")
	var pass = os.Getenv("DB_PASSWORD")
	var host = os.Getenv("DB_HOST")
	var db = os.Getenv("DB_DATABASE")

	// concatenate into form -> user:password@tcp(host)/database_being_accessed
	connection_string := user + ":" + pass + "@tcp(" + host + ")/" + db

	return connection_string
}

// Currently Deprecated/unused
func readYaml(queryName string, conf Config) string {
	reader, err := os.Open("test.yml")
	checkError(err)

	// reads yaml file
	buf, err := ioutil.ReadAll(reader)
	checkError(err)

	// unmarshals yaml file from raw to unicode
	err = yaml.Unmarshal(buf, &conf)
	checkError(err)

	query := conf.Data[queryName]
	return query
}

func connectDB(connection_string string, query string) *sql.Rows {
	db, err := sql.Open("mysql", connection_string)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection Established")
	}

	rows, err := db.Query(query)
	checkError(err)

	defer db.Close() // To be clean and tidy, close database when done.

	return rows
}

func displayData(rows *sql.Rows, first_name string, last_name string,
	id string, grade string, users []User) []User {

	for rows.Next() {
		err := rows.Scan(&first_name, &last_name, &id, &grade)
		checkError(err)
		users = append(users, User{First_name: first_name, Last_name: last_name, ID: id, Grade: grade})
	}

	return users
}

// Sets up connection, reads config file, executes queries found in config file on connected database
func executeQuery(c *gin.Context) {
	var first_name string
	var last_name string
	var id string
	var grade string
	var users []User
	var conf Config

	connection_string := establishConnection()
	queryName := c.Param("queryName")

	// =============================
	//TODO: put this in function (that works as expected)
	// locates and opens yaml file
	reader, err := os.Open("test.yml")
	checkError(err)

	// reads yaml file
	buf, err := ioutil.ReadAll(reader)
	checkError(err)

	// unmarshals yaml file from raw to unicode
	err = yaml.Unmarshal(buf, &conf)
	checkError(err)

	query := conf.Data[queryName]
	// ============================

	// Error message alerts user to a failure without closing the app
	// Allows for user to retry their request
	if _, ok := conf.Data[queryName]; !ok {
		c.Data(404, "text/plain", []byte("Query not found: "+queryName))
		return
	}

	rows := connectDB(connection_string, query)
	users = displayData(rows, first_name, last_name, id, grade, users)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
		"query": users,
	})
}

func main() {
	var port = ":" + os.Getenv("DB_PORT")

	r := gin.Default()
	r.LoadHTMLGlob("*.html")
	r.Handle("GET", "/:queryName", executeQuery)
	r.Run(port)
}
