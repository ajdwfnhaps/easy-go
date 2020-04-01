package web

import (
	"github.com/ajdwfnhaps/easy-gin/conf"
	"github.com/ajdwfnhaps/easy-go/docs"
)

// SetSwaggerInfo 设置SwaggerInfo
func SetSwaggerInfo() {
	cfg := conf.Global()
	swag := cfg.Swagger

	if swag.On != 1 {
		return
	}

	// programatically set swagger info
	docs.SwaggerInfo.Title = swag.Title
	docs.SwaggerInfo.Description = swag.Description
	docs.SwaggerInfo.Version = swag.Version
	docs.SwaggerInfo.Host = swag.Host
	docs.SwaggerInfo.BasePath = swag.BasePath
	docs.SwaggerInfo.Schemes = swag.Schemes
}
