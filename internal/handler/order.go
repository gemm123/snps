package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"synapsis/internal/model"
	"synapsis/internal/service"
)

type orderHandler struct {
	orderService service.OrderService
	cartService  service.CartService
}

func NewOrderHandler(orderService service.OrderService, cartService service.CartService) *orderHandler {
	return &orderHandler{
		orderService: orderService,
		cartService:  cartService,
	}
}

// @Summary Add order products
// @Description Add products to order
// @Tags order
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization header (Bearer [token])"
// @Param request body []model.OrderProductRequest true "order product details"
// @Security ApiKeyAuth
// @Success 200 {object} model.Meta "add product to cart success"
// @Failure 500 {object} model.Meta "Internal Server Error"
// @Failure 400 {object} model.Meta "Bad Request"
// @Router /order [post]
func (h *orderHandler) AddOrderProduct(ctx *gin.Context) {
	var orderProductRequest []model.OrderProductRequest

	idUserString := ctx.GetString("id")
	idUser, _ := uuid.Parse(idUserString)

	if err := ctx.ShouldBindJSON(&orderProductRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": model.Meta{
				Message: err.Error(),
				Status:  400,
			},
		})
		return
	}

	if err := h.orderService.AddOrderProduct(orderProductRequest, idUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"meta": model.Meta{
				Message: err.Error(),
				Status:  500,
			},
		})
		return
	}

	for _, opr := range orderProductRequest {
		if err := h.cartService.DeleteCartProduct(opr.IdProduct, idUser); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"meta": model.Meta{
					Message: err.Error(),
					Status:  500,
				},
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"meta": model.Meta{
			Message: "success",
			Status:  200,
		},
	})
}

// @Summary Get order product
// @Description Get all order product
// @Tags order
// @Produce json
// @Param Authorization header string true "Authorization header (Bearer [token])"
// @Security ApiKeyAuth
// @Success 200 {object} model.DataResponse "add product to cart success"
// @Failure 500 {object} model.Meta "Internal Server Error"
// @Router /order [get]
func (h *orderHandler) GetOrderProduct(ctx *gin.Context) {
	idUserString := ctx.GetString("id")
	idUser, _ := uuid.Parse(idUserString)

	orderResponses, err := h.orderService.GetAllOrderProduct(idUser)
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
		Data: orderResponses,
	})
}

// @Summary Pay order
// @Description Pay order to change status order
// @Tags order
// @Produce json
// @Param Authorization header string true "Authorization header (Bearer [token])"
// @Param id-order path string true "id order"
// @Security ApiKeyAuth
// @Success 200 {object} model.Meta "add product to cart success"
// @Failure 500 {object} model.Meta "Internal Server Error"
// @Router /order/pay/{id-order} [put]
func (h *orderHandler) PayOrder(ctx *gin.Context) {
	idOrderString := ctx.Param("id-order")
	idOrder, _ := uuid.Parse(idOrderString)

	if err := h.orderService.PayOrder(idOrder); err != nil {
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
