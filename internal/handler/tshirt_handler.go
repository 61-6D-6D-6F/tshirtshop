package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/61-6D-6D-6F/tshirtshop/internal/model"
	"github.com/61-6D-6D-6F/tshirtshop/internal/service"
)

type TShirtHandler struct {
	tShirtService service.TShirtService
}

func NewTShirtHandler(s service.TShirtService) *TShirtHandler {
	return &TShirtHandler{tShirtService: s}
}

func (h *TShirtHandler) ListTShirts(c *gin.Context) {
	tshirts, err := h.tShirtService.ListTShirts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tshirts)
}

func (h *TShirtHandler) CreateTShirt(c *gin.Context) {
	var tshirt model.TShirt
	if err := c.ShouldBindJSON(&tshirt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.tShirtService.CreateTShirt(tshirt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *TShirtHandler) UpdateTShirt(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var tShirt model.TShirt
	if err := c.ShouldBindJSON(&tShirt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.tShirtService.UpdateTShirt(id, tShirt); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
}

func (h *TShirtHandler) DeleteTShirt(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.tShirtService.DeleteTShirt(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
