package user

import (
	"golang/pkg/utils"
	userDto "golang/src/user/dto"
	userService "golang/src/user/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

)

type UserHandler struct {
	userService userService.UserService
}

func NewUserHandler(userService userService.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (userHandler *UserHandler) Route(r *gin.RouterGroup) {

	userRouter := r.Group("/api/user")

	userRouter.GET("", userHandler.FindAll)
	userRouter.GET("/:id", userHandler.FindById)
	userRouter.POST("", userHandler.Create)
	userRouter.PATCH("/:id", userHandler.Update)
	userRouter.DELETE("/:id", userHandler.Delete)
}

func (userHandler *UserHandler) FindAll(ctx *gin.Context) {
	data := userHandler.userService.FindAll()

	// return data
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "success fetch data all user", data))
}

func (userHandler *UserHandler) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := userHandler.userService.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, err.Error(), nil))
		return
	}

	// return data
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "success fetch data user with id "+ctx.Param("id"), data))
}

func (userHandler *UserHandler) Create(ctx *gin.Context) {
	var input userDto.CreateUserRequest

	requestErr := ctx.ShouldBindJSON(&input)

	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "Bad Request", requestErr.Error()))
		ctx.Abort()
		return
	}

	createdUser, err := userHandler.userService.Create(input)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, err.Error(), nil))
		return
	}

	// return data
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "success create user", createdUser))
}

func (userHandler *UserHandler) Update(ctx *gin.Context) {
	var input userDto.UpdateUserRequestBody

	requestErr := ctx.ShouldBindJSON(&input)

	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, requestErr.Error(), nil))
		ctx.Abort()
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	updatedUser, err := userHandler.userService.Update(id, input)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, err.Error(), nil))
		return
	}

	// return data
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "success update user", updatedUser))
}

func (userHandler *UserHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := userHandler.userService.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, err.Error(), nil))
		return
	}

	// return data
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "success delete user with id "+ctx.Param("id"), nil))
}