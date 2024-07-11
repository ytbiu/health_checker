package router

import (
	"github.com/gin-gonic/gin"
	"health_checker/controller"
)

func Init(r *gin.Engine) {
	r.GET("api/v1/report", controller.Report)

}
