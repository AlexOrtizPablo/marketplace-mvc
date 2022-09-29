package controller

import (
	"github.com/gin-gonic/gin"
	"marketplace-mvc/model"
	"marketplace-mvc/service"
	"net/http"
	"strconv"
)

type AddressController struct {
	service *service.AddressService
}

func (c *AddressController) Create(ctx *gin.Context) {
	var address model.Address
	err := ctx.ShouldBindJSON(&address)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = c.service.Create(address)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c *AddressController) GetById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, err := strconv.Atoi(idStr)
	id := uint(idInt)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	address, err := c.service.GetById(id)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, address)
}

func (c *AddressController) GetAll(ctx *gin.Context) {
	addresses, err := c.service.GetAll()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, addresses)
}

func (c *AddressController) Update(ctx *gin.Context) {
	var address model.Address
	err := ctx.ShouldBindJSON(&address)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = c.service.Update(address)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *AddressController) Delete(ctx *gin.Context) {
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
