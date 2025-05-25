package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/61-6D-6D-6F/tshirtshop/internal/model"
	"github.com/61-6D-6D-6F/tshirtshop/internal/repository"
)

type OrderHandler struct {
	orderRepository repository.OrderRepository
}

func NewOrderHandler(r repository.OrderRepository) *OrderHandler {
	return &OrderHandler{orderRepository: r}
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order, err := h.orderRepository.List(id)
	if err != nil {
		c.JSON(http.StatusOK, []model.Order{})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) Checkout(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// simulated payment always succeed
	order, err := h.orderRepository.Checkout(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, order)
}
