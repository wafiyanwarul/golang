package publisher

import (
	"golang/pkg/utils"
	publisherDto "golang/src/publisher/dto"
	publisherService "golang/src/publisher/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

)

type PublisherHandler struct {
	publisherService publisherService.PublisherService
}

func NewPublisherHandler(publisherService publisherService.PublisherService) *PublisherHandler {
	return &PublisherHandler{publisherService}
}

func (publisherHandler *PublisherHandler) Route(r *gin.RouterGroup) {
	publisherRouter := r.Group("/api/publisher")

	publisherRouter.GET("", publisherHandler.FindAll)
	publisherRouter.GET("/:id", publisherHandler.FindById)
	publisherRouter.POST("", publisherHandler.Create)
	// publisherRouter.PATCH("/:id", publisherHandler.Update)
	// publisherRouter.DELETE("/:id", publisherHandler.Delete)
}

func (publisherHandler *PublisherHandler) FindAll(ctx *gin.Context) {

	data := publisherHandler.publisherService.FindAll()

	// return data

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "success fetch data all publisher", data))

}

func (publisherHandler *PublisherHandler) FindById(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := publisherHandler.publisherService.FindById(id)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, "publisher not found", nil))
			ctx.Abort()
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, err.Error(), nil))
			ctx.Abort()
			return
		}
	}

	// return data
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "success fetch data publisher with id "+ctx.Param("id"), data))
}

func (publisherHandler *PublisherHandler) Create(ctx *gin.Context) {

	var input publisherDto.CreatePublisherRequest

	inputErr := ctx.ShouldBindJSON(&input)

	if inputErr != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, inputErr.Error(), nil))
		ctx.Abort()
		return
	}

	data, err := publisherHandler.publisherService.Create(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, err.Error(), nil))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "success create data publisher", data))

}
