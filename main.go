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
	"github.com/61-6D-6D-6F/tshirtshop/internal/service"
)

func main() {
	dbPath := "tshirts.db"
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

	repo := repository.NewTShirtRepository(db)
	srv := service.NewTShirtService(repo)
	h := handler.NewTShirtHandler(srv)

	r := gin.Default()

	r.POST("/login", handler.Login)

	r.GET("/tshirts", h.ListTShirts)
	r.GET("/tshirts/:id", h.GetTShirt)

	admin := r.Group("/admin")
	admin.Use(handler.JWTAuthMiddleware())
	{
		admin.POST("/tshirts", h.CreateTShirt)
		admin.PUT("/tshirts/:id", h.UpdateTShirt)
		admin.DELETE("/tshirts/:id", h.DeleteTShirt)
	}

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
