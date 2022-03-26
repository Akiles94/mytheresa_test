package adapters

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/Akiles94/mytheresa-test/domain/ports"
	"github.com/Akiles94/mytheresa-test/infrastructure/httputils"
	"github.com/gin-gonic/gin"
)

type API struct {
	Port         int
	ProductsImpl ports.IProducts
	handler      *gin.Engine
	routes       *gin.RouterGroup
}

func (api *API) Start() {
	handler := gin.New()
	handler.Use(gin.Recovery())
	handler.Use(httputils.Logger())

	routesV1 := handler.Group("/v1")

	api.handler = handler
	api.routes = routesV1

	routesV1.Use(httputils.ErrorHandler())

	api.initRoutes()
	api.ListenAndServe()
}

// ListenAndServe starts the API
func (api *API) ListenAndServe() {
	port := strconv.Itoa(api.Port)
	log.Print(port)
	host := fmt.Sprintf(":%s", port)

	srv := &http.Server{
		Addr:    host,
		Handler: api.handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	log.Printf("Server listening on address %s\n", host)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
