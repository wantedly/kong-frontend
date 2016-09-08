package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wantedly/kong-frontend/kong"
	"github.com/wantedly/kong-frontend/model/consumer"
)

type ConsumerController struct {
	Client *kong.Client
}

func NewConsumerController(client *kong.Client) *ConsumerController {
	return &ConsumerController{client}
}

func (self *ConsumerController) Index(c *gin.Context) {
	consumers, err := consumer.List(self.Client)
	if err != nil {
		c.HTML(http.StatusBadRequest, "consumers.tmpl", gin.H{
			"error": err,
		})
		return
	}
	c.HTML(http.StatusOK, "consumers.tmpl", gin.H{
		"consumers": consumers.Consumer,
		"total":     consumers.Total,
	})
}

func (self *ConsumerController) Get(c *gin.Context) {
	id := c.Param("consumerID")
	res, err := consumer.Get(self.Client, id)
	if err != nil {
		c.HTML(http.StatusBadRequest, "consumer.tmpl", gin.H{
			"error": err,
		})
		return
	}
	c.HTML(http.StatusOK, "consumer.tmpl", gin.H{
		"consumer": res,
	})
}

func (self *ConsumerController) Create(c *gin.Context) {
	form := new(kong.Consumer)
	err := c.Bind(form)
	if err == nil {
		var res *kong.Consumer
		res, err = consumer.Create(self.Client, form)
		if err == nil {
			c.HTML(http.StatusOK, "consumer.tmpl", gin.H{
				"success":  "created",
				"consumer": res,
			})
			return
		}
	}
	c.HTML(http.StatusBadRequest, "new-consumer.tmpl", gin.H{
		"error":    err,
		"consumer": form,
	})
}

func (self *ConsumerController) Delete(c *gin.Context) {
	id := c.Param("consumerID")
	_, err := consumer.Delete(self.Client, id)
	if err != nil {
		c.HTML(http.StatusBadRequest, "consumer.tmpl", gin.H{
			"error": err,
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/consumers")
}

func (self *ConsumerController) Update(c *gin.Context) {
	id := c.Param("consumerID")
	form := new(kong.Consumer)
	err := c.Bind(form)
	if err == nil {
		form.ID = id
		var res *kong.Consumer
		res, err = consumer.Update(self.Client, form)
		if err == nil {
			c.HTML(http.StatusOK, "consumer.tmpl", gin.H{
				"success":  "updated",
				"consumer": res,
			})
			return
		}
	}
	c.HTML(http.StatusBadRequest, "consumer.tmpl", gin.H{
		"error":    err,
		"consumer": form,
	})
}

func (self *ConsumerController) New(c *gin.Context) {
	c.HTML(http.StatusOK, "new-consumer.tmpl", gin.H{})
}
