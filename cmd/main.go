package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"smart-card/internal/database"
	"smart-card/internal/handlers"
)

func main() {
	database.InitDB()
	r := gin.Default()

	r.GET("/admin/users", handlers.HandleGetUsers)
	r.POST("/access", handlers.HandleAccess)
	r.POST("/admin/add-user", handlers.HandleAddUser)
	r.GET("/admin/events", handlers.HandleGetEvents)
	r.DELETE("/admin/delete-user", handlers.HandleDeleteUser)

	log.Println("Сервер запущен на :8080")
	r.Run(":8080")
}
