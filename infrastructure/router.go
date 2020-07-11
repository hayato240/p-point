package infrastructure

import (
	gin "github.com/gin-gonic/gin"
	"github.com/hayato240/p-point/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	userController := controllers.NewUserController(NewSqlHandler())
	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })
	router.POST("/users/:id/points", func(c *gin.Context) { userController.Points(c) })
	Router = router
}
