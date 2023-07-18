package main

import (
	mysql "golang/pkg/db/mysql"
	author "golang/src/author/injector"
	book "golang/src/book/injector"
	oauth "golang/src/oauth/injector"
	publisher "golang/src/publisher/injector"
	register "golang/src/register/injector"
	rental "golang/src/rental/injector"
	user "golang/src/user/injector"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	pathdir, _ := os.Getwd()
	environment := godotenv.Load(filepath.Join(pathdir, ".env"))

	if environment != nil {
		panic(environment)
	}
}

func main() {
	r := gin.Default()

	db := mysql.DB() // <- yg ini

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Successfully connected",
		})
	})

	user.InitializeUserHandler(db).Route(&r.RouterGroup)
	register.InitializeService(db).Route(&r.RouterGroup)
	author.InitializeService(db).Route(&r.RouterGroup)
	publisher.InitializeService(db).Route(&r.RouterGroup)
	book.InitializeService(db).Route(&r.RouterGroup)
	oauth.InitializeService(db).Route(&r.RouterGroup)
	rental.InitializeService(db).Route(&r.RouterGroup)

	r.Run("127.0.0.1:8000") // listen and serve on 0.0.0.0:8000 (for windows "localhost:8000")
}
