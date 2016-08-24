package server

import (
	"github.com/gin-gonic/gin"
	"github.com/wantedly/kong-oauth-token-generator/config"
	"github.com/wantedly/kong-oauth-token-generator/controller"
	"github.com/wantedly/kong-oauth-token-generator/kong"
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

	r.GET("/", apiController.Index)

	r.GET("/apis", apiController.Index)
	r.POST("/apis", apiController.Create)
	r.GET("/apis/:apiName", apiController.Get)
	r.POST("/apis/:apiName/delete", apiController.Delete)
	r.GET("/new-api", apiController.New)

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

	r.Run()
}
