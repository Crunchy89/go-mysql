package main

import (
	"fmt"
	"os"

	_userRoute "github.com/Crunchy89/go-mysql/app/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	println("Initialize server")
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	println("env loaded")
	gin.SetMode(gin.ReleaseMode)
	var db *gorm.DB
	if os.Getenv("DB_DRIVER") == "pgsql" {
		var url string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
		db, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	} else {
		var url string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
		db, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	}

	if err != nil {
		panic(err.Error())
	}

	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "PUT", "GET", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowHeaders:     []string{"*"},
	}))

	_userRoute.UserRoute(server, db)

	port := ":" + os.Getenv("PORT")
	if port == ":" || port == "" {
		port = ":" + "8080"
	}
	if err := server.Run(port); err != nil {
		panic(err)
	}
}
