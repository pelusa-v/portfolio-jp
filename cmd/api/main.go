package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pelusa-v/portfolio-jp/internal/infrastructure/api"
)

func main() {
	engine := gin.Default()
	api.RegisterUsersRouter(engine)
	engine.Run(":8000")
}
