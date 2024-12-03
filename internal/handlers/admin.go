package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"smart-card/internal/database"
)

func HandleAddUser(c *gin.Context) {
	var user database.User

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("Ошибка парсинга JSON:", err) // Лог ошибки парсинга
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		fmt.Println("Ошибка создания пользователя:", err) // Лог ошибки базы данных
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать пользователя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Пользователь добавлен", "user": user})
}

func HandleGetUsers(c *gin.Context) {
	var users []database.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить список пользователей"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func HandleGetEvents(c *gin.Context) {
	var events []database.Event

	successFilter := c.Query("success")
	query := database.DB

	if successFilter != "" {
		if successFilter == "true" {
			query = query.Where("success = ?", true)
		} else if successFilter == "false" {
			query = query.Where("success = ?", false)
		}
	}

	if err := query.Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить список событий"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"events": events})
}

func HandleDeleteUser(c *gin.Context) {
	var request struct {
		CardUID string `json:"card_uid"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	var user database.User
	if err := database.DB.Where("card_uid = ?", request.CardUID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить пользователя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Пользователь удалён", "user": user.Name})
}
