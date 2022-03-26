package adapters

import (
	"net/http"

	"github.com/Akiles94/mytheresa-test/infrastructure/httputils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (a *API) ListProducts(c *gin.Context) {
	queryParams, err := httputils.GetProductsParams(c)
	if err != nil {
		log.Error("Error parsing limit and offset from string to int")
		c.JSON(http.StatusInternalServerError, err)
	}
	result, exc := a.ProductsImpl.GetProducts(*queryParams)
	if exc != nil {
		log.Error("Error getting response from GetProducts use case: ", exc.Message)
		c.JSON(http.StatusInternalServerError, exc)
		return
	}
	c.JSON(http.StatusOK, result)
	return
}
