package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"synapsis/internal/model"
	"synapsis/internal/service"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) Register(ctx *gin.Context) {
	var userRegister model.Register

	if err := ctx.ShouldBindJSON(&userRegister); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": model.Meta{
				Message: err.Error(),
				Status:  400,
			},
		})
		return
	}

	if err := h.userService.Register(userRegister); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"meta": model.Meta{
				Message: err.Error(),
				Status:  500,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"meta": model.Meta{
			Message: "success",
			Status:  200,
		},
	})
}

func (h *userHandler) Login(ctx *gin.Context) {
	var login model.Login

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": model.Meta{
				Message: err.Error(),
				Status:  400,
			},
		})
		return
	}

	token, err := h.userService.Login(login)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"meta": model.Meta{
				Message: err.Error(),
				Status:  500,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"meta": model.Meta{
			Message: "success",
			Status:  200,
		},
		"token": token,
	})
}
