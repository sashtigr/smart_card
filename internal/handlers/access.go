package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smart-card/internal/database"
)

func HandleAccess(c *gin.Context) {
	var request struct {
		CardUID string `json:"card_uid"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	var user database.User
	success := true
	if err := database.DB.Where("card_uid = ?", request.CardUID).First(&user).Error; err != nil {
		success = false
		database.DB.Create(&database.Event{CardUID: request.CardUID, Success: success})
		c.JSON(http.StatusForbidden, gin.H{"status": "Доступ запрещен"})
		return
	}

	if !user.Access {
		success = false
		database.DB.Create(&database.Event{CardUID: request.CardUID, Success: success})
		c.JSON(http.StatusForbidden, gin.H{"status": "Доступ запрещен"})
		return
	}

	database.DB.Create(&database.Event{CardUID: request.CardUID, Success: success})
	c.JSON(http.StatusOK, gin.H{"status": "Доступ разрешен", "user": user.Name})
}
