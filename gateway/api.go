package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/sliaptsou/backend/proto"
)

type Api struct {
	cl proto.BackendClient
}

func NewApi(cl proto.BackendClient) *Api {
	return &Api{
		cl: cl,
	}
}

// TODO: add tests
func (api *Api) GetQueryCount(c *gin.Context) {
	log.Print("Received GetQueryCount request")
	r := new(proto.Empty)
	resp, err := api.cl.GetQueryCount(c.Request.Context(), r)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, resp)
	return
}

func (api *Api) HealthCheck(c *gin.Context) {
	log.Print("Received HealthCheck request")
	r := new(proto.Empty)
	resp, err := api.cl.HealthCheck(c.Request.Context(), r)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, resp)
	return
}
