package server

import (
	"github.com/gin-gonic/gin"
	"github.com/wantedly/kong-frontend/config"
	"github.com/wantedly/kong-frontend/controller"
	"github.com/wantedly/kong-frontend/kong"
)

func Run(config *config.KongConfiguration) {
	r := gin.Default()
	r.Static("/assets", "assets")
	r.LoadHTMLGlob("templates/*")

	r.Use(gin.Logger())
	client := kong.NewClient(nil, config)
	apiController := controller.NewAPIController(client)
	oauth2Controller := controller.NewOAuth2Controller(client)
	pluginController := controller.NewPluginController(client)
	consumerController := controller.NewConsumerController(client)
	basicAuthController := controller.NewBasicAuthController(client)

	r.GET("/", apiController.Index)

	r.GET("/apis", apiController.Index)
	r.POST("/apis", apiController.Create)
	r.GET("/apis/:apiName", apiController.Get)
	r.POST("/apis/:apiName/delete", apiController.Delete)
	r.GET("/new-api", apiController.New)

	r.GET("/consumers", consumerController.Index)
	r.POST("/consumers", consumerController.Create)
	r.GET("/consumers/:consumerID", consumerController.Get)
	r.POST("/consumers/:consumerID", consumerController.Update)
	r.POST("/consumers/:consumerID/delete", consumerController.Delete)
	r.GET("/new-consumer", consumerController.New)

	r.GET("/oauth2s", oauth2Controller.Index)
	r.POST("/oauth2s", oauth2Controller.Create)
	r.GET("/oauth2s/:consumerName", oauth2Controller.Get)
	r.POST("/oauth2s/:consumerName/delete", oauth2Controller.Delete)
	r.GET("/new-oauth2", oauth2Controller.New)

	r.GET("/apis/:apiName/plugins", pluginController.Index)
	r.POST("/apis/:apiName/plugins", pluginController.Create)
	r.GET("/apis/:apiName/plugins/:pluginID", pluginController.Get)
	r.POST("/apis/:apiName/plugins/:pluginID/delete", pluginController.Delete)
	r.GET("/apis/:apiName/new-plugin", pluginController.New)
	r.POST("/apis/:apiName/new-plugin", pluginController.SetConfig)

	r.GET("/consumers/consumerID/basic-auth", basicAuthController.Index)
	r.POST("/consumers/:consumerID/basic-auth", basicAuthController.Create)
	r.GET("/consumers/:consumerID/basic-auth/:credentialID", basicAuthController.Get)
	r.POST("/consumers/:consumerID/basic-auth/:credentialID", basicAuthController.Update)
	r.POST("/consumers/:consumerID/basic-auth/:credentialID/delete", basicAuthController.Delete)
	r.GET("/consumers/:consumerID/new-basic-auth", basicAuthController.New)

	r.Run()
}
