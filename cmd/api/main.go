package main

import (
	"github.com/gin-gonic/gin"
	api "github.com/pelusa-v/portfolio-jp/internal/infrastructure/api/routes"
)

func main() {
	engine := gin.Default()
	api.RegisterUsersRouter(engine)
	api.RegisterTagsRouter(engine)
	engine.Run(":8000")
}
