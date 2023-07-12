package register

import (
	"golang/pkg/utils"
	registerDto "golang/src/register/dto"
	registerService "golang/src/register/service"
	userDto "golang/src/user/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	registerService registerService.RegisterService
}

func NewRegisterHandler(registerService registerService.RegisterService) *RegisterHandler {
	return &RegisterHandler{registerService}
}

func (registerHandler *RegisterHandler) Route(r *gin.RouterGroup) {
	r.POST("/api/register", registerHandler.Register)
	r.POST("/api/verify", registerHandler.Verify)
}

func (registerHandler *RegisterHandler) Register(ctx *gin.Context) {
	var registerRequest userDto.CreateUserRequest

	inputErr := ctx.ShouldBindJSON(&registerRequest)

	if inputErr != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "Invalid request body", nil))
		ctx.Abort()
		return
	}

	requestErr := registerHandler.registerService.Register(registerRequest)

	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusInternalServerError, requestErr.Error(), nil))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Register success", "Please check your email"))
}

func (registerHandler *RegisterHandler) Verify(ctx *gin.Context) {
	var verifyRequestBody registerDto.VerifyEmail

	inputErr := ctx.ShouldBindJSON(&verifyRequestBody)

	if inputErr != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "Invalid request body", nil))
		ctx.Abort()
		return
	}

	requestErr := registerHandler.registerService.VerifyEmail(verifyRequestBody)

	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusInternalServerError, requestErr.Error(), nil))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "Verify success", nil))
}
