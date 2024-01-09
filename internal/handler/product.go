package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"synapsis/internal/model"
	"synapsis/internal/service"
)

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *productHandler {
	return &productHandler{
		productService: productService,
	}
}

func (h *productHandler) GetAllProductByNameCategory(ctx *gin.Context) {
	categoryName := ctx.Param("category-name")

	products, err := h.productService.GetAllProductByCategoryName(categoryName)
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
		"data": products,
	})
}
