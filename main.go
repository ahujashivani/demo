package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		db, err := sql.Open("mysql", "root:rootpassword@tcp(127.0.0.1:3306)/test-database")
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Connection Established")
		}
		var (
			id   int
			name string
		)

		type Message struct {
			Num  int
			Name string
		}

		rows, err := db.Query("SELECT id,first_name FROM authors")
		if err != nil {
			log.Fatal(err)
		}

		var users []Message
		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, Message{Num: id, Name: name})
		}
		c.JSON(http.StatusOK, users)
	})

	r.Run(":8384")
}
