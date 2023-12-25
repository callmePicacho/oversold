package app

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	// 产生超卖
	r.POST("/oversold", Oversold)
}
