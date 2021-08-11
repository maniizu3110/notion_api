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

	//認証なしで利用できるエンドポイント
	{
		g := e.Group("/api/v1",middlewares.SetDB(db))
		handler.AssignUserHandlers(g.Group("/user"))
	}
	
	//認証が必要なエンドポイント
	{
		g := e.Group("/api/v1",middlewares.SetDB(db),middlewares.AuthMiddleware())
		handler.AssignNotionDatabaseHandlers(g.Group("/databases"))                                                  // auth ok
		handler.AssignBrockHandlers(g.Group("/block"))
	}
	e.Logger.Fatal(e.Start(":8080"))
}

// {
//     "access_token": "v2.local.GG4LIJPv7ZOW3oFQOz99E8VQj0JcQspHSp3M-VDzk4Iln7Q5diMYZrAZul2jED4_oupWwWOjhr9YZlkbwsld2gv1EiOh2H2PsE9bIVQKLAetJnHIuienj20vxEot3GhPyJUnvLnvtbFcn2iHnJhLOX0DqFUr2zzkCZrYdqmxQSqVRZoMTpSwCagL0VUR5v2kPqGRDthKAZ4bDwJr1M3VaVZ0CSICWyYRn4airo-_ktnqhzeZLmCva8EcA8OYcurBgq020pVjgO8.bnVsbA",
//     "user": {
//         "user": "leoleo",
//         "created_at": "2021-08-11T12:19:06+09:00"
//     }
// }