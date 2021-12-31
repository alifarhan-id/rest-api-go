package main

import (
	"rest-api-go/config"

	"rest-api-go/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/honorarirum/:id", inDB.GetHonorarium)

	router.Run(":3000")
}
