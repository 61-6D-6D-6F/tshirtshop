package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	"github.com/61-6D-6D-6F/tshirtshop/internal/handler"
	"github.com/61-6D-6D-6F/tshirtshop/internal/repository"
)

func main() {
	dbPath := "tshirtshop.db"
	if envPath := os.Getenv("TSHIRT_DB_PATH"); envPath != "" {
		dbPath = envPath
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("failed to open db:", err)
	}
	defer db.Close()
	if err := repository.InitDB(db); err != nil {
		log.Fatal("failed to migrate db:", err)
	}

	userRepo := repository.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepo)

	tShirtRepo := repository.NewTShirtRepository(db)
	tShirtHandler := handler.NewTShirtHandler(tShirtRepo)

	cartRepo := repository.NewCartRepository(db)
	cartHandler := handler.NewCartHandler(cartRepo)

	orderRepo := repository.NewOrderRepository(db)
	orderHandler := handler.NewOrderHandler(orderRepo)

	r := gin.Default()

	r.POST("/login", userHandler.Login)
	r.POST("/register", userHandler.Register)

	r.GET("/tshirts", tShirtHandler.ListTShirts)
	r.GET("/tshirts/:id", tShirtHandler.GetTShirt)

	admin := r.Group("/admin")
	admin.Use(handler.AdminMiddleware())
	{
		admin.POST("/tshirts", tShirtHandler.CreateTShirt)
		admin.PUT("/tshirts/:id", tShirtHandler.UpdateTShirt)
		admin.DELETE("/tshirts/:id", tShirtHandler.DeleteTShirt)

		admin.GET("/users", userHandler.ListUsers)
		admin.GET("/users/:id", userHandler.GetUser)
		admin.POST("/users", userHandler.CreateUser)
		admin.PUT("/users/:id", userHandler.UpdateUser)
		admin.DELETE("/users/:id", userHandler.DeleteUser)
	}

	cart := r.Group("/cart")
	cart.Use(handler.JWTAuthMiddleware())
	{
		cart.GET("/:userid", cartHandler.GetCart)
		cart.POST("/:userid/add", cartHandler.AddToCart)
		cart.DELETE("/:userid/remove", cartHandler.RemoveFromCart)
	}

	order := r.Group("/order")
	order.Use(handler.JWTAuthMiddleware())
	{
		order.GET("/:userid", orderHandler.ListOrders)
		order.POST("/:userid/checkout", orderHandler.Checkout)
	}

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
