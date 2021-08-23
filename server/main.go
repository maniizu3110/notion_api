package main

import (
	"fmt"
	"net/url"
	"os"
	"server/app"
	"server/util"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func init(){
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		logrus.Fatal("cannot load config:", err)
	}

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBAddress, config.DBPort, config.DBName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Tokyo")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := gorm.Open(`mysql`, dsn)
	if err != nil {
		logrus.Fatal(err)
	}
	defer func() {
		err := dbConn.Close()
		if err != nil {
			logrus.Fatal(err)
		}
	}()

	server, err := app.NewServer(config, dbConn)
	if err != nil {
		logrus.Fatal("cannot create server:", err)
	}
	app.InitRouter(server)
}
