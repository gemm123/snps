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

// @Summary Get all products by category name
// @Description Retrieve a list of products based on the specified category name
// @Tags product
// @Produce json
// @Param category-name path string true "Category name (baju/celana/topi)"
// @Param Authorization header string true "Authorization header (Bearer [token])"
// @Security ApiKeyAuth
// @Success 200 {object} model.DataResponse "List of products"
// @Failure 500 {object} model.Meta "Internal Server Error"
// @Router /category/{category-name}/product [get]
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

	ctx.JSON(http.StatusOK, model.DataResponse{
		Meta: model.Meta{
			Message: "success",
			Status:  200,
		},
		Data: products,
	})
}
