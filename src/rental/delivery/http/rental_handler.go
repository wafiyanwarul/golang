package rental

import (
	"golang/src/middleware"
	rentalDto "golang/src/rental/dto"
	rentalService "golang/src/rental/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RentalHandler struct {
	rentalService rentalService.RentalService
}

func NewRentalHandler(rentalService rentalService.RentalService) *RentalHandler {
	return &RentalHandler{rentalService}
}

func (rentalHandler RentalHandler) Route(r *gin.RouterGroup) {
	rentalRouter := r.Group("/api/rental")
	rentalRouter.Use(middleware.AuthJwt)
	{
		rentalRouter.POST("/", rentalHandler.Create)
		rentalRouter.POST("/:id/return", rentalHandler.Return)
		rentalRouter.PUT("/:id", rentalHandler.Update)
	}

	rentalRouter.GET("/", rentalHandler.FindAll)
	rentalRouter.GET("/:id", rentalHandler.FindById)
}

func (rentalHandler RentalHandler) Create(ctx *gin.Context) {
	var input rentalDto.CreateRentalRequest

	inputErr := ctx.ShouldBindJSON(&input)

	if inputErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": inputErr.Error(),
			"data":    nil,
		})

		ctx.Abort()
		return
	}

	rental, err := rentalHandler.rentalService.Create(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})

		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Success create rental",
		"data":    rental,
	})
}

func (rentalHandler RentalHandler) FindAll(ctx *gin.Context) {
	rentals := rentalHandler.rentalService.FindAll()

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get all rentals",
		"data":    rentals,
	})
}

func (rentalHandler RentalHandler) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	rental, err := rentalHandler.rentalService.FindById(id)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "rental not found",
				"data":    nil,
			})

			ctx.Abort()
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "there is something wrong",
				"data":    err.Error(),
			})

			ctx.Abort()
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get rental by id " + ctx.Param("id"),
		"data":    rental,
	})
}

func (rentalHandler RentalHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input rentalDto.UpdateRentalRequest

	inputErr := ctx.ShouldBindJSON(&input)

	if inputErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": inputErr.Error(),
			"data":    nil,
		})

		ctx.Abort()
		return
	}

	rental, err := rentalHandler.rentalService.Update(id, input)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "rental not found",
				"data":    nil,
			})

			ctx.Abort()
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "there is something wrong",
				"data":    nil,
			})

			ctx.Abort()
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success update rental by id" + ctx.Param("id"),
		"data":    rental,
	})
}

func (rentalHandler RentalHandler) Return(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input rentalDto.ReturnRentalRequest

	inputErr := ctx.ShouldBind(&input)

	if inputErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": inputErr.Error(),
			"data":    nil,
		})
		ctx.Abort()
		return
	}

	rental, err := rentalHandler.rentalService.Return(id, input)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "rental not found",
				"data":    nil,
			})

			ctx.Abort()
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "there is something wrong",
				"data":    nil,
			})

			ctx.Abort()
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success returned rental by id" + ctx.Param("id"),
		"data":    rental,
	})
}
