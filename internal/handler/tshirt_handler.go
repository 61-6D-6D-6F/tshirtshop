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
	tShirts, err := h.tShirtService.ListTShirts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tShirts)
}

func (h *TShirtHandler) GetTShirt(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tShirt, err := h.tShirtService.GetTShirt(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tShirt)
}

func (h *TShirtHandler) CreateTShirt(c *gin.Context) {
	var tShirt model.TShirt
	if err := c.ShouldBindJSON(&tShirt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tShirt.Name == "" || tShirt.Size == "" || tShirt.Color == "" ||
		tShirt.Price == 0.0 || tShirt.Stock == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.tShirtService.CreateTShirt(&tShirt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, tShirt)
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
	if tShirt.Name == "" || tShirt.Size == "" || tShirt.Color == "" ||
		tShirt.Price == 0.0 || tShirt.Stock == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.tShirtService.UpdateTShirt(id, &tShirt); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tShirt)
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
