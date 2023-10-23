package routes

import (
	"github.com/OvictorVieira/transact.ease/internal/config"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	var mode = gin.ReleaseMode
	if config.AppConfig.Debug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	router := gin.New()

	return router
}
