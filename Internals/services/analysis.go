package service

import (
    "net/http"
    "time"
    "sales-analytics/config"
    "sales-analytics/internals/loader"

    "github.com/gin-gonic/gin"
)

func GetTotalRevenue(c *gin.Context) {
    start := c.Query("start")
    end := c.Query("end")
    startDate, _ := time.Parse("2006-01-02", start)
    endDate, _ := time.Parse("2006-01-02", end)

    var total float64
    config.DB.Raw(`
        SELECT SUM(quantity_sold * unit_price * (1 - discount)) 
        FROM orders 
        WHERE date_of_sale BETWEEN ? AND ?
    `, startDate, endDate).Scan(&total)

    c.JSON(http.StatusOK, gin.H{"total_revenue": total})
}

func RefreshData(c *gin.Context) {
    err := loader.LoadCSV("data/sales.csv")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Refresh failed", "error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Refresh successful"})
}
