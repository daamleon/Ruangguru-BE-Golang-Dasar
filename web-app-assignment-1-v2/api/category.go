package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryAPI struct {
	categoryService service.CategoryService
}

func NewCategoryAPI(categoryService service.CategoryService) CategoryAPI {
	return CategoryAPI{categoryService: categoryService}
}

func (c *CategoryAPI) AddCategory(ctx *gin.Context) {
	var category model.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	err := c.categoryService.Store(&category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{Message: "category add success"})
}

func (c *CategoryAPI) GetCategoryByID(ctx *gin.Context) {
	categoryID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid category ID"})
		return
	}

	category, err := c.categoryService.GetByID(categoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

func (c *CategoryAPI) UpdateCategory(ctx *gin.Context) {
	categoryID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Category ID"})
		return
	}

	var category model.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	category.ID = categoryID
	err = c.categoryService.Update(categoryID, category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{Message: "category update success"})
}

func (c *CategoryAPI) DeleteCategory(ctx *gin.Context) {
	categoryID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid category ID"})
		return
	}

	err = c.categoryService.Delete(categoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{Message: "category delete success"})
}

func (c *CategoryAPI) GetCategoryList(ctx *gin.Context) {
	categories, err := c.categoryService.GetList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}