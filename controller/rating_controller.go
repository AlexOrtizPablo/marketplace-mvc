package controller

import (
	"github.com/gin-gonic/gin"
	"marketplace-mvc/model"
	"marketplace-mvc/service"
	"net/http"
	"strconv"
)

type RatingController struct {
	service *service.RatingService
}

func (c *RatingController) Create(ctx *gin.Context) {
	var rating model.Rating
	err := ctx.ShouldBindJSON(&rating)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = c.service.Create(rating)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c *RatingController) GetById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, err := strconv.Atoi(idStr)
	id := uint(idInt)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	rating, err := c.service.GetById(id)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, rating)
}

func (c *RatingController) GetAll(ctx *gin.Context) {
	ratings, err := c.service.GetAll()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, ratings)
}

func (c *RatingController) Update(ctx *gin.Context) {
	var rating model.Rating
	err := ctx.ShouldBindJSON(&rating)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = c.service.Update(rating)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *RatingController) Delete(ctx *gin.Context) {
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
