package oauth

import (
	oauthDto "golang/src/oauth/dto"
	oauthService "golang/src/oauth/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OauthHandler struct {
	oauthService oauthService.OauthService
}

func NewOauthHandler(oauthService oauthService.OauthService) *OauthHandler {
	return &OauthHandler{oauthService}
}

func (oauthHandler *OauthHandler) Route(r *gin.RouterGroup) {
	oauthRouter := r.Group("/api/oauth")
	oauthRouter.POST("/login", oauthHandler.Login)
}

func (oauthHandler *OauthHandler) Login(ctx *gin.Context) {
	var input oauthDto.LoginRequest

	inputErr := ctx.ShouldBindJSON(&input)

	if inputErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "input not valid",
			"data":    nil,
		})

		ctx.Abort()
		return
	}

	data, err := oauthHandler.oauthService.Login(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})

		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success login",
		"data":    data,
	})
}
