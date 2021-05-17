package main

import (
    "github.com/gin-gonic/gin"
    "github.com/mngeow/golang-ticketing-module/routers/tables"
)

func main() {
	router := gin.Default()

    v1 := router.Group("/api")
    tables.TablesRegisterRoutes(v1.Group("/tables"))
    router.Run()

}