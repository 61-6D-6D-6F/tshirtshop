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
	userHandle := handler.NewUserHandler(userRepo)

	tShirtRepo := repository.NewTShirtRepository(db)
	tShirtHandle := handler.NewTShirtHandler(tShirtRepo)

	cartRepo := repository.NewCartRepository(db)
	cartHandle := handler.NewCartHandler(cartRepo)

	r := gin.Default()

	r.POST("/login", userHandle.Login)
	r.POST("/register", userHandle.Register)

	r.GET("/tshirts", tShirtHandle.ListTShirts)
	r.GET("/tshirts/:id", tShirtHandle.GetTShirt)

	admin := r.Group("/admin")
	admin.Use(handler.AdminMiddleware())
	{
		admin.POST("/tshirts", tShirtHandle.CreateTShirt)
		admin.PUT("/tshirts/:id", tShirtHandle.UpdateTShirt)
		admin.DELETE("/tshirts/:id", tShirtHandle.DeleteTShirt)

		admin.GET("/users", userHandle.ListUsers)
		admin.GET("/users/:id", userHandle.GetUser)
		admin.POST("/users", userHandle.CreateUser)
		admin.PUT("/users/:id", userHandle.UpdateUser)
		admin.DELETE("/users/:id", userHandle.DeleteUser)
	}

	cart := r.Group("/cart")
	cart.Use(handler.JWTAuthMiddleware())
	{
		cart.GET("/:userid", cartHandle.GetCart)
		cart.POST("/:userid/add", cartHandle.AddToCart)
		cart.DELETE("/:userid/remove", cartHandle.RemoveFromCart)
	}

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
