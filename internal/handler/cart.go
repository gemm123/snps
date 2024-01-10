package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"synapsis/internal/model"
	"synapsis/internal/service"
)

type cartHandler struct {
	cartService service.CartService
}

func NewCartHandler(cartService service.CartService) *cartHandler {
	return &cartHandler{
		cartService: cartService,
	}
}

func (h *cartHandler) AddProductToCart(ctx *gin.Context) {
	var cartRequest model.CartRequest

	idUserString := ctx.GetString("id")
	idUser, _ := uuid.Parse(idUserString)

	if err := ctx.ShouldBindJSON(&cartRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": model.Meta{
				Message: err.Error(),
				Status:  400,
			},
		})
		return
	}

	if err := h.cartService.AddProductToCart(cartRequest, idUser); err != nil {
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

func (h *cartHandler) GetCartProduct(ctx *gin.Context) {
	idUserString := ctx.GetString("id")
	idUser, _ := uuid.Parse(idUserString)

	cartResponse, err := h.cartService.GetAllCartProduct(idUser)
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
		"data": cartResponse,
	})
}
