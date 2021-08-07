package main

import (
	"fmt"
	"log"
	"net/url"
	"server/controller/handler"
	"server/middlewares"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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

	InitRouter(dbConn)
}

func InitRouter(db *gorm.DB){
	
	e := echo.New()
	middlewares.CORS(e)

	g := e.Group("/api/v1",middlewares.SetDB(db))
	handler.AssignNotionDatabaseHandlers(g.Group("/databases"))                                                  // auth ok
	handler.AssignBrockHandlers(g.Group("/block"))
	e.Logger.Fatal(e.Start(":8080"))
}