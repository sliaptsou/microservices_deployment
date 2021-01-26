package main

import (
	"github.com/gin-gonic/gin"
	proto "github.com/sliaptsou/backend/proto"
	"log"
	"net/http"
	"strconv"
)

func (api Api) GetList(c *gin.Context) {
	log.Print("Received GetList request")
	c.JSON(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
	return
}

func (api Api) GetOne(c *gin.Context) {
	log.Print("Received GetOne request")
	req := &proto.GetOneItemRequest{}

	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	req.Id = int32(id)

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	res, err := api.cl.GetOne(c.Request.Context(), req)
	log.Printf("gateway: %+v", err)
	if err != nil {
		c.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (api Api) Create(c *gin.Context) {
	log.Print("Received Create request")
	req := &proto.CreateRequest{}

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	res, err := api.cl.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (api Api) Update(c *gin.Context) {
	log.Print("Received Update request")
	c.JSON(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
	return
}

func (api Api) Delete(c *gin.Context) {
	log.Print("Received Delete request")
	c.JSON(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
	return
}
