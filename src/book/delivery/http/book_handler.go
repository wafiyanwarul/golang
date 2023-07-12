package book

import (
	"golang/pkg/utils"
	bookDto "golang/src/book/dto"
	bookService "golang/src/book/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService bookService.BookService
}

func NewBookHandler(bookService bookService.BookService) *BookHandler {
	return &BookHandler{bookService}
}

func (bookHandler *BookHandler) Route(r *gin.RouterGroup) {
	bookRouter := r.Group("/api/book")

	bookRouter.GET("", bookHandler.FindAll)
	bookRouter.GET("/:id", bookHandler.FindById)
	bookRouter.POST("", bookHandler.Create)
	// bookRouter.PATCH("/:id", bookHandler.Update)
	// bookRouter.DELETE("/:id", bookHandler.Delete)
}

func (bookHandler *BookHandler) FindAll(ctx *gin.Context) {
	data := bookHandler.bookService.FindAll()

	// return data
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "success fetch data all book", data))
}

func (bookHandler *BookHandler) FindById(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := bookHandler.bookService.FindById(id)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, "book not found", nil))
			ctx.Abort()
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, err.Error(), nil))
			ctx.Abort()
			return
		}
	}

	// return data
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "success fetch data book with id "+ctx.Param("id"), data))
}

func (bookHandler *BookHandler) Create(ctx *gin.Context) {
	var input bookDto.CreateBookRequest

	inputErr := ctx.ShouldBindJSON(&input)

	if inputErr != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, inputErr.Error(), nil))
		ctx.Abort()
		return
	}

	data, err := bookHandler.bookService.Create(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, err.Error(), nil))
		ctx.Abort()
		return
	}

	// return data

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "success create data book", data))
}