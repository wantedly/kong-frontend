package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/wantedly/kong-frontend/model/oauth2"

	"github.com/gin-gonic/gin"
	"github.com/wantedly/kong-frontend/kong"
)

type generateOAuth2Params struct {
	Username string `form:"userName" json:"userName" binding:"required"`
}

type OAuth2Controller struct {
	Client *kong.Client
}

func NewOAuth2Controller(client *kong.Client) *OAuth2Controller {
	return &OAuth2Controller{client}
}

func (self *OAuth2Controller) Index(c *gin.Context) {
	consumers, err := oauth2.List(self.Client)
	c.HTML(http.StatusOK, "oauth2s.tmpl", gin.H{
		"alert":     false,
		"error":     false,
		"message":   "",
		"consumers": consumers.Consumer,
		"total":     consumers.Total,
		"err":       err,
	})
	return
}

func (self *OAuth2Controller) Get(c *gin.Context) {
	consumerName := c.Param("consumerName")

	if !oauth2.Exists(self.Client, consumerName) {
		c.HTML(http.StatusNotFound, "oauth2.tmpl", gin.H{
			"error":   true,
			"message": fmt.Sprintf("%s does not exist.", consumerName),
		})
		return
	}

	consumerDetail, assigneesOAuth2Detail, err := oauth2.Get(self.Client, consumerName)
	fmt.Fprintf(os.Stdout, "consumerDetail %+v\n assigneesOAuth2Detail %+v\n err %+v\n", consumerDetail, assigneesOAuth2Detail, err)
	if consumerDetail == nil {
		fmt.Fprintf(os.Stderr, "Err: %+v\nTarget user name: %+v\n", err, consumerName)

		c.HTML(http.StatusInternalServerError, "oauth2.tmpl", gin.H{
			"error":   true,
			"message": "Failed to list app URLs.",
		})

		return
	}

	c.HTML(http.StatusOK, "oauth2.tmpl", gin.H{
		"error":                 false,
		"consumerDetail":        consumerDetail,
		"assigneesOAuth2Detail": assigneesOAuth2Detail,
	})
	return
}

func (self *OAuth2Controller) New(c *gin.Context) {
	c.HTML(http.StatusOK, "new-oauth2.tmpl", gin.H{
		"alert":   false,
		"error":   false,
		"message": "",
	})
	return
}

func (self *OAuth2Controller) Delete(c *gin.Context) {
	consumerName := c.Param("consumerName")
	message, err := oauth2.Delete(self.Client, consumerName)
	if err != nil {
		c.HTML(http.StatusBadRequest, "oauth2s.tmpl", gin.H{
			"error":   true,
			"err":     fmt.Sprintf("%s", err),
			"message": fmt.Sprintf("%s", message),
		})
	}

	c.Redirect(http.StatusMovedPermanently, "/oauth2s")
	return
}

func (self *OAuth2Controller) Create(c *gin.Context) {
	var form generateOAuth2Params
	if c.Bind(&form) == nil {
		fmt.Fprintf(os.Stdout, "User Name %+v\n", form.Username)
	} else {
		c.HTML(http.StatusBadRequest, "new-api.tmpl", gin.H{
			"error":   true,
			"message": fmt.Sprintf("Please fix params"),
		})
		return
	}

	if oauth2.Exists(self.Client, form.Username) {
		c.HTML(http.StatusConflict, "new-oauth2.tmpl", gin.H{
			"error":   true,
			"message": fmt.Sprintf("User Name %s already exist.", form.Username),
		})
		return
	}

	generateConsumer := &kong.Consumer{
		Username: form.Username,
	}

	createdConsumer, createdAssigneesOAuth2, err := oauth2.Create(self.Client, generateConsumer)
	if err != nil {
		c.HTML(http.StatusServiceUnavailable, "new-oauth2.tmpl", gin.H{
			"error":   true,
			"message": fmt.Sprintf("Please Check kong: %s", err),
		})
		return
	}

	c.HTML(http.StatusOK, "oauth2.tmpl", gin.H{
		"error":                 false,
		"consumerDetail":        createdConsumer,
		"assigneesOAuth2Detail": createdAssigneesOAuth2,
		"message":               fmt.Sprintf("Success"),
	})
	return
}
