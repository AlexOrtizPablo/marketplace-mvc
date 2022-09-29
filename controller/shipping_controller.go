package controller

import (
	"github.com/gin-gonic/gin"
	"marketplace-mvc/model"
	"marketplace-mvc/service"
	"net/http"
	"strconv"
)

type ShippingController struct {
	service *service.ShippingService
}

func (c *ShippingController) Create(ctx *gin.Context) {
	var shipping model.Shipping
	err := ctx.ShouldBindJSON(&shipping)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = c.service.Create(shipping)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c *ShippingController) GetById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, err := strconv.Atoi(idStr)
	id := uint(idInt)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	shipping, err := c.service.GetById(id)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, shipping)
}

func (c *ShippingController) GetAll(ctx *gin.Context) {
	shippingList, err := c.service.GetAll()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, shippingList)
}

func (c *ShippingController) Update(ctx *gin.Context) {
	var shipping model.Shipping
	err := ctx.ShouldBindJSON(&shipping)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = c.service.Update(shipping)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *ShippingController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, err := strconv.Atoi(idStr)
	id := uint(idInt)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = c.service.Delete(id)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
