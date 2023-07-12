package author

import (
	"golang/pkg/utils"
	authorDto "golang/src/author/dto"
	authorService "golang/src/author/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

)

type AuthorHandler struct {
	authorService authorService.AuthorService
}

func NewAuthorHandler(authorService authorService.AuthorService) *AuthorHandler {
	return &AuthorHandler{authorService}
}

func (authorHandler *AuthorHandler) Route(r *gin.RouterGroup) {
	authorRouter := r.Group("/api/author")

	authorRouter.GET("", authorHandler.FindAll)
	authorRouter.GET("/:id", authorHandler.FindById)
	authorRouter.POST("", authorHandler.Create)
	// authorRouter.PATCH("/:id", authorHandler.Update)
	// authorRouter.DELETE("/:id", authorHandler.Delete)
}

func (authorHandler *AuthorHandler) FindAll(ctx *gin.Context) {
	data := authorHandler.authorService.FindAll()

	// return data
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "success fetch data all author", data))
}

func (authorHandler *AuthorHandler) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := authorHandler.authorService.FindById(id)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, "author not found", nil))
			ctx.Abort()
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, err.Error(), nil))
			ctx.Abort()
			return
		}
	}

	// return data
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "success fetch data author with id "+ctx.Param("id"), data))
}

func (authorHandler *AuthorHandler) Create(ctx *gin.Context) {
	var input authorDto.CreateAuthorRequest

	requestErr := ctx.ShouldBindJSON(&input)

	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "Bad Request", requestErr.Error()))
		ctx.Abort()
		return
	}

	data, err := authorHandler.authorService.Create(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, err.Error(), nil))
		ctx.Abort()
		return
	}

	// return data
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "success create author", data))
}