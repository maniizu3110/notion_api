package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"server/middlewares"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/dstotijn/go-notion"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbAddress := viper.GetString(`database.address`)
	dbName := viper.GetString(`database.dbname`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbAddress, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Tokyo")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := gorm.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middlewares.CORS(e)

	testKey := viper.GetString(`notion_test.key`)
	fmt.Println(testKey)
	testDB := viper.GetString(`notion_test.db_id`)
	fmt.Println(testDB)
	client := notion.NewClient(testKey)
	e.GET("/", func(c echo.Context) error {
		page, err := client.FindDatabaseByID(context.Background(), testDB)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, page)
	})
	e.GET("/brock", func(c echo.Context) error {
		query := new(notion.PaginationQuery)
		block, err := client.FindBlockChildrenByID(context.Background(), "5edbea1047824fa38a42f815489f0206", query)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, block)
	})
	log.Fatal(e.Start(viper.GetString("server.address")))
}
