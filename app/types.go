package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Req 语义：谁购买了多少个商品
type Req struct {
	UserId int `json:"user_id"`
	SkuId  int `json:"sku_id"`
	Num    int `json:"num"`
}

type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func Marshal(c *gin.Context) *Req {
	var req Req
	err := c.ShouldBind(&req)
	if err != nil {
		panic("参数错误")
	}

	return &req
}

func Fail(c *gin.Context, msg string) {
	code := http.StatusInternalServerError
	c.JSON(code, &Resp{Code: code, Msg: msg})
}

func OK(c *gin.Context) {
	code := http.StatusOK
	c.JSON(code, &Resp{Code: code, Msg: "OK"})
}
