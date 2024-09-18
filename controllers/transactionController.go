package controllers

import (
	"ecommerce-backend/config"
	"ecommerce-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTransactionByID(c *gin.Context) {
	var transaction models.Transaction
	id := c.Param("id")

	if err := config.DB.Preload("User").Preload("Item").First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func GetTransactionsByUserID(c *gin.Context) {
	var transactions []models.Transaction
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := config.DB.Where("user_id = ?", userID).Preload("Item").Find(&transactions).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transactions not found"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func CreateTransaction(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input struct {
		ItemIDs []uint `json:"item_ids"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	for _, itemID := range input.ItemIDs {
		var item models.Item
		if err := config.DB.First(&item, itemID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}

		transaction := models.Transaction{
			UserID: user.ID,
			ItemID: item.ID,
		}

		if err := config.DB.Create(&transaction).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction created successfully"})
}

func DeleteTransaction(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	transactionID := c.Param("id")

	var transaction models.Transaction

	if err := config.DB.First(&transaction, transactionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	if transaction.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this transaction"})
		return
	}

	if err := config.DB.Delete(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}
