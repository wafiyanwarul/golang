package middleware

import (
	utils "golang/pkg/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Header struct {
	Authorization string `header:"authorization" binding:"required"`
}

func AuthJwt(ctx *gin.Context) {
	var input Header

	err := ctx.ShouldBindHeader(&input)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.Response(http.StatusUnauthorized, "Unauthorized", nil))
		ctx.Abort()
		return
	}

	reqToken := input.Authorization

	splitToken := strings.Split(reqToken, "Bearer ")

	if len(splitToken) != 2 {
		ctx.JSON(http.StatusUnauthorized, utils.Response(http.StatusUnauthorized, "Unauthorized", nil))
		ctx.Abort()
		return
	}

	reqToken = splitToken[1]

	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.Response(http.StatusUnauthorized, "Unauthorized", err.Error()))
		ctx.Abort()
		return
	}

	if !token.Valid {
		ctx.JSON(http.StatusUnauthorized, utils.Response(http.StatusUnauthorized, "Unauthorized", err.Error()))
		ctx.Abort()
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	ctx.Set("user", claims["email"])

	ctx.Next()

}
