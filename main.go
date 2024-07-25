package main

import (
    "github.com/gin-gonic/gin"
	"app/config/router"
	"app/dao"
)

func main() {
	router := gin.Default()

    //connect db
	dao.ConnectDB()

	//create router
	routerConfig.CreateRouter(router)

	//run server
	router.Run("localhost:8080")
}