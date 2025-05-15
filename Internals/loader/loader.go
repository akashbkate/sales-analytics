package loader

import (
	"encoding/csv"

	"os"
	"strconv"
	"time"

	"sales-analytics/config"
	"sales-analytics/internals/models"
)

func LoadCSV(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for i, row := range records {
		if i == 0 {
			continue // skip header
		}

		quantity, _ := strconv.Atoi(row[7])
		unitPrice, _ := strconv.ParseFloat(row[8], 64)
		discount, _ := strconv.ParseFloat(row[9], 64)
		shippingCost, _ := strconv.ParseFloat(row[10], 64)
		dateOfSale, _ := time.Parse("2006-01-02", row[6])

		config.DB.Create(&models.Customer{
			ID:      row[2],
			Name:    row[12],
			Email:   row[13],
			Address: row[14],
		})

		config.DB.Create(&models.Product{
			ID:       row[1],
			Name:     row[3],
			Category: row[4],
		})

		config.DB.Create(&models.Order{
			ID:            row[0],
			ProductID:     row[1],
			CustomerID:    row[2],
			Region:        row[5],
			DateOfSale:    dateOfSale,
			QuantitySold:  quantity,
			UnitPrice:     unitPrice,
			Discount:      discount,
			ShippingCost:  shippingCost,
			PaymentMethod: row[11],
		})
	}
	return nil
}
