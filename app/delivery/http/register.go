package http

import (
	"github.com/DarkSoul94/password-generator/app"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.RouterGroup, uc app.Usecase) {
	h := NewHandler(uc)

		router.GET("/pass", h.GeneratePassword)
}
