package httputils

import (
	"strconv"

	"github.com/Akiles94/mytheresa-test/domain/dto"
	"github.com/Akiles94/mytheresa-test/domain/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetProductsFiltered(data []models.Product, filters dto.QueryParams) []models.Product {
	response := []models.Product{}
	for _, item := range data {
		catFlag := false
		priceFlag := false
		if filters.Category != nil {
			if item.Category == *filters.Category {
				catFlag = true
			}
		} else {
			catFlag = true
		}
		if filters.PriceLessThan != nil {
			if item.Price < *filters.PriceLessThan {
				priceFlag = true
			}
		} else {
			priceFlag = true
		}
		if catFlag && priceFlag {
			response = append(response, item)
		}
	}
	return response
}

func GetProductsParams(c *gin.Context) (*dto.QueryParams, error) {
	var result dto.QueryParams
	var category *string
	var priceLessThan *int
	tmpCategory := c.DefaultQuery("category", "")
	if tmpCategory != "" {
		category = &tmpCategory
	}
	tmpPricesLessThan, err := strconv.Atoi(c.DefaultQuery("priceLessThan", "0"))
	if err != nil {
		log.Error("Error converting from string to int for offset")
		return nil, err
	}
	if tmpPricesLessThan != 0 {
		priceLessThan = &tmpPricesLessThan
	}
	result = dto.QueryParams{
		Category:      category,
		PriceLessThan: priceLessThan,
	}

	return &result, nil
}
