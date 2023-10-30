package routes

import (
	"github.com/agmguerra/go-crud-example/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/", controller.FindUsers)
	r.GET("/user/mail/:userMail", controller.FindUserByEmail)
	r.GET("/user/:userId", controller.FindUserById)
	r.POST("/user/", controller.CreateUser)
	r.PUT("/user/:userId", controller.UpdateUser)
	r.DELETE("/user/:userId", controller.DeleteUser)
}
