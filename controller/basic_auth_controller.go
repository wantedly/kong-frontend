package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wantedly/kong-frontend/kong"
	"github.com/wantedly/kong-frontend/model/basic_auth"
)

type BasicAuthController struct {
	Client *kong.Client
}

func NewBasicAuthController(client *kong.Client) *BasicAuthController {
	return &BasicAuthController{client}
}

func (self *BasicAuthController) Index(c *gin.Context) {
	consumerID := c.Param("consumerID")
	credentials, err := basic_auth.List(self.Client, consumerID)
	if err != nil {
		c.HTML(http.StatusBadRequest, "basic-auths.tmpl", gin.H{
			"error":      err,
			"consumerID": consumerID,
		})
		return
	}
	c.HTML(http.StatusOK, "basic-auths.tmpl", gin.H{
		"credentials": credentials.Data,
		"total":       credentials.Total,
		"consumerID":  consumerID,
	})
}

func (self *BasicAuthController) Get(c *gin.Context) {
	consumerID := c.Param("consumerID")
	credentialID := c.Param("credentialID")
	credential, err := basic_auth.Get(self.Client, consumerID, credentialID)
	if err != nil {
		c.HTML(http.StatusBadRequest, "basic_auth.tmpl", gin.H{
			"error":      err,
			"consumerID": consumerID,
		})
		return
	}
	c.HTML(http.StatusOK, "basic-auth.tmpl", gin.H{
		"credential": credential,
		"consumerID": consumerID,
	})
}

func (self *BasicAuthController) Create(c *gin.Context) {
	consumerID := c.Param("consumerID")
	form := new(kong.BasicAuthCredential)
	err := c.Bind(form)
	if err == nil {
		var credential *kong.BasicAuthCredential
		credential, err = basic_auth.Create(self.Client, consumerID, form)
		if err == nil {
			c.HTML(http.StatusOK, "basic-auth.tmpl", gin.H{
				"success":    "created",
				"credential": credential,
				"consumerID": consumerID,
			})
			return
		}
	}
	c.HTML(http.StatusBadRequest, "new-basic-auth.tmpl", gin.H{
		"error":      err,
		"credential": form,
		"consumerID": consumerID,
	})
}

func (self *BasicAuthController) Delete(c *gin.Context) {
	consumerID := c.Param("consumerID")
	credentialID := c.Param("credentialID")
	_, err := basic_auth.Delete(self.Client, consumerID, credentialID)
	if err != nil {
		c.HTML(http.StatusBadRequest, "basic-auth.tmpl", gin.H{
			"error":      err,
			"consumerID": consumerID,
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/consumers/"+consumerID+"/basic-auth/")
}

func (self *BasicAuthController) Update(c *gin.Context) {
	consumerID := c.Param("consumerID")
	credentialID := c.Param("credentialID")
	form := new(kong.BasicAuthCredential)
	err := c.Bind(form)
	if err == nil {
		form.ID = credentialID
		var res *kong.BasicAuthCredential
		res, err = basic_auth.Update(self.Client, consumerID, form)
		if err == nil {
			c.HTML(http.StatusOK, "basic-auth.tmpl", gin.H{
				"success":    "updated",
				"credential": res,
				"consumerID": consumerID,
			})
			return
		}
	}
	c.HTML(http.StatusBadRequest, "basic-auth.tmpl", gin.H{
		"error":      err,
		"credential": form,
		"consumerID": consumerID,
	})
}

func (self *BasicAuthController) New(c *gin.Context) {
	consumerID := c.Param("consumerID")
	c.HTML(http.StatusOK, "new-basic-auth.tmpl", gin.H{
		"consumerID": consumerID,
	})
}
