package routes

import (
	"google-translate-go/pkg/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.POST("/translate", controller.Translate)
}
