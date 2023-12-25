package main

import (
	"github.com/gin-gonic/gin"
	"oversold/app"
)

func main() {
	r := gin.Default()

	app.Router(r)

	panic(r.Run(":9090"))
}
