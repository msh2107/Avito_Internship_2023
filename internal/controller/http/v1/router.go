package v1

import (
	_ "Avito/docs"
	"Avito/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			DynamicUserSegmentationService
//	@version		1.0
//	@description	This is a service for dynamic user segmentation.
//	@host			localhost:8080
//	@BasePath		/v1

func NewRouter(handler *gin.Engine, services *service.Services) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// Routers
	h := handler.Group("/v1")
	{
		newSegmentRoutes(h.Group("/segment"), services.Segment)
		newUserSegmentRoutes(h, services.UserSegment)
	}
}
