package app

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	// 产生超卖
	r.POST("/oversold", Oversold)
	// 解决方案1：线程内 mutex
	r.POST("/oversold/mutex", OversoldMutex)
	// 解决方案2：线程内 chan 实现互斥锁
	r.POST("/oversold/chan", OversoldChan)
}
