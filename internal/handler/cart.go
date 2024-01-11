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

// @Summary Add product to cart
// @Description add product to cart
// @Tags cart
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header (Bearer [token])"
// @Param request body model.CartRequest true "cart request details"
// @Security ApiKeyAuth
// @Success 200 {object} model.Meta "add product to cart success"
// @Failure 500 {object} model.Meta "Internal Server Error"
// @Failure 400 {object} model.Meta "Bad Request"
// @Router /cart [post]
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

// @Summary Get cart product
// @Description get product at cart
// @Tags cart
// @Produce json
// @Param Authorization header string true "Authorization header (Bearer [token])"
// @Security ApiKeyAuth
// @Success 200 {object} model.DataResponse "add product to cart success"
// @Failure 500 {object} model.Meta "Internal Server Error"
// @Router /cart [get]
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

	ctx.JSON(http.StatusOK, model.DataResponse{
		Meta: model.Meta{
			Message: "success",
			Status:  200,
		},
		Data: cartResponse,
	})
}

// @Summary Delete cart product
// @Description Delete product at cart
// @Tags cart
// @Produce json
// @Param id-product path string true "id product"
// @Param Authorization header string true "Authorization header (Bearer [token])"
// @Security ApiKeyAuth
// @Success 200 {object} model.Meta "success delete"
// @Failure 500 {object} model.Meta "Internal Server Error"
// @Router /cart/delete/{id-product} [delete]
func (h *cartHandler) DeleteCartProduct(ctx *gin.Context) {
	idUserString := ctx.GetString("id")
	idUser, _ := uuid.Parse(idUserString)
	idProductString := ctx.Param("id-product")
	idProduct, _ := uuid.Parse(idProductString)

	if err := h.cartService.DeleteCartProduct(idProduct, idUser); err != nil {
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
