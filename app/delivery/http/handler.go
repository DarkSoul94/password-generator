package http

import (
	"net/http"
	"strconv"

	"github.com/DarkSoul94/password-generator/app"
	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler struct {
	uc app.Usecase
}

// NewHandler ...
func NewHandler(uc app.Usecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

// HelloWorld ...
func (h *Handler) GeneratePassword(c *gin.Context) {
	length, err := strconv.Atoi(c.Query("length"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err})
		return
	}

	digitsCount, err := strconv.Atoi(c.Query("digitsCount"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err})
		return
	}

	withUpper, err := strconv.ParseBool(c.Query("withUpper"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err})
		return
	}

	allowRepeat, err := strconv.ParseBool(c.Query("allowRepeat"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err})
		return
	}

	password, err := h.uc.GeneratePassword(length, digitsCount, withUpper, allowRepeat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "error": err})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "success", "password": password})
}
