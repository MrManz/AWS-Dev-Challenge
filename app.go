package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
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
		dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("RDS_HOSTNAME"),
			os.Getenv("RDS_USERNAME"),
			os.Getenv("RDS_PASSWORD"),
			os.Getenv("RDS_DB_NAME"))

		db, err := sql.Open("postgres", dbinfo)
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
