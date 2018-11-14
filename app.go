package main

import (
	"database/sql"
	"encoding/json"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Message struct {
	Objects []Object `json:"objects"`
}

type Object struct {
	Id   int8
	Name string
}

func main() {
	e := echo.New()
	e.Static("/", "public")
	e.GET("/data", func(c echo.Context) error {
		// Connect to database
		connStr := "postgres://admin_postgres:password@mydbinstance.cl0v5hbertly.us-east-2.rds.amazonaws.com/db01"
		db, err := sql.Open("postgres", connStr)
		checkErr(err)

		// Fire Query
		rows, err := db.Query("SELECT did,name FROM distributors")
		checkErr(err)

		// Loop over returned rows and build response
		ob := Object{}
		mes := Message{}

		for rows.Next() {
			err = rows.Scan(&ob.Id, &ob.Name)
			checkErr(err)
			mes.Objects = append(mes.Objects, ob)
		}

		res, err := json.Marshal(mes)
		checkErr(err)
		return c.String(http.StatusOK, string(res))

	})

	e.Logger.Fatal(e.Start(":3000"))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
