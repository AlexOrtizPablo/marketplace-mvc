package controller

import (
	"github.com/gin-gonic/gin"
	"marketplace-mvc/model"
	"marketplace-mvc/service"
	"net/http"
	"strconv"
)

type ProductController struct {
	service *service.ProductService
}

func (c *ProductController) Create(ctx *gin.Context) {
	var product model.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = c.service.Create(product)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c *ProductController) GetById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, err := strconv.Atoi(idStr)
	id := uint(idInt)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	product, err := c.service.GetById(id)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) GetAll(ctx *gin.Context) {
	products, err := c.service.GetAll()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (c *ProductController) Update(ctx *gin.Context) {
	var product model.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = c.service.Update(product)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *ProductController) Delete(ctx *gin.Context) {
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
