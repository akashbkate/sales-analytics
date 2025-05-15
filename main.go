package main

import (
    "github.com/gin-gonic/gin"
    "sales-analytics/config"
    "sales-analytics/Internals/handler"
)

func main() {
    config.InitDB()
    r := gin.Default()
    handler.RegisterRoutes(r)
    r.Run(":8080")
}
