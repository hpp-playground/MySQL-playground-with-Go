package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type kyokyu struct {
	DepartmentID string `json:"department_id"`
	PartsID      string `json:"parts_id"`
	MerchantID   string `json:"merchant_id"`
	Price        string `json:"price"`
	Quantity     string `json:"quantity"`
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", handle)
	e.POST("/", handlePost)

	e.Logger.Fatal(e.Start(":1323"))
}

type param struct {
	Query string `json:"query"`
}

func handlePost(c echo.Context) error {
	param := new(param)
	if err := c.Bind(param); err != nil {
		return err
	}
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/kadai")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println(param.Query)

	rows, err := db.Query(param.Query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []kyokyu
	for rows.Next() {
		kyokyu := kyokyu{}
		if err := rows.Scan(&kyokyu.DepartmentID, &kyokyu.PartsID, &kyokyu.MerchantID, &kyokyu.Price, &kyokyu.Quantity); err != nil {
			log.Fatal(err)
		}
		result = append(result, kyokyu)
	}

	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func handle(c echo.Context) error {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/kadai")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM 供給")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []kyokyu
	for rows.Next() {
		kyokyu := kyokyu{}
		if err := rows.Scan(&kyokyu.DepartmentID, &kyokyu.PartsID, &kyokyu.MerchantID, &kyokyu.Price, &kyokyu.Quantity); err != nil {
			log.Fatal(err)
		}
		result = append(result, kyokyu)
	}

	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
	return c.JSON(http.StatusAccepted, result)
}
