package routes

import (
	"github.com/OvictorVieira/transact.ease/internal/config"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// set the runtime mode
	var mode = gin.ReleaseMode
	if config.AppConfig.Debug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	// create a new router instance
	router := gin.New()

	return router
}
