package controller

import (
	"github.com/gin-gonic/gin"
	"marketplace-mvc/model"
	"marketplace-mvc/service"
	"net/http"
	"strconv"
)

type UserController struct {
	service *service.UserService
}

func (c *UserController) Create(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = c.service.Create(user)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c *UserController) GetById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, err := strconv.Atoi(idStr)
	id := uint(idInt)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	user, err := c.service.GetById(id)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) GetAll(ctx *gin.Context) {
	users, err := c.service.GetAll()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) Update(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = c.service.Update(user)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *UserController) Delete(ctx *gin.Context) {
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

func (c *UserController) Login(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = c.service.Login(user)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusOK)
}
