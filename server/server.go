package server

import (
	"github.com/gin-gonic/gin"
	"github.com/koudaiii/kong-oauth-token-generator/config"
	"github.com/koudaiii/kong-oauth-token-generator/controller"
	"github.com/koudaiii/kong-oauth-token-generator/kong"
)

func Run(config *config.KongConfiguration) {
	r := gin.Default()
	r.Static("/assets", "assets")
	r.LoadHTMLGlob("templates/*")

	r.Use(gin.Logger())
	client := kong.NewClient(nil, config)
	rootController := controller.NewRootController(client)
	apiController := controller.NewAPIController(client)
	// consumerController := controller.NewConsumerController(config)
	// oauth2Controller := controller.NewOAuth2Controller(config)

	r.GET("/", rootController.Index)

	r.GET("/apis", apiController.Index)
	r.POST("/apis", apiController.Create)
	r.GET("/apis/:apiName", apiController.Get)
	r.POST("/apis/:apiName/delete", apiController.Delete)
	r.GET("/new-api", apiController.New)

	r.Run()
}
