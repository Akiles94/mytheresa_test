package adapters

import "github.com/gin-gonic/gin"

func (a *API) initRoutes() {
	a.routes.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "MyTheresa API Service OK!"})
	})
	a.routes.GET("/products", a.ListProducts)
}
