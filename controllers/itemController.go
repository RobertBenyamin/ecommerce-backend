package controllers

import (
	"ecommerce-backend/config"
	"ecommerce-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllItems(c *gin.Context) {
	var items []models.Item
	if err := config.DB.Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
		return
	}
	c.JSON(http.StatusOK, items)
}

func GetItemByID(c *gin.Context) {
	var item models.Item
	id := c.Param("id")

	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func GetItemsByUserID(c *gin.Context) {
	var items []models.Item
	
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := config.DB.Where("user_id = ?", userID).Find(&items).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Items not found"})
		return
	}

	c.JSON(http.StatusOK, items)
}

func CreateItem(c *gin.Context) {
    var item models.Item

    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    if err := c.ShouldBindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    if item.Name == "" || item.Description == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "All fields (name, price, description) are required"})
        return
    }

    if item.Price <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be a positive number"})
        return
    }

    item.UserID = userID.(uint)

    config.DB.Create(&item)
    c.JSON(http.StatusOK, item)
}


func UpdateItem(c *gin.Context) {
    var item models.Item
    id := c.Param("id")

    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    if err := config.DB.First(&item, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
        return
    }

    if item.UserID != userID.(uint) {
        c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this item"})
        return
    }

    var input models.Item
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    if input.Name == "" || input.Description == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "All fields (name, price, description) are required"})
        return
    }

    if input.Price <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be a positive number"})
        return
    }

    item.Name = input.Name
    item.Price = input.Price
    item.Description = input.Description

    config.DB.Save(&item)

    c.JSON(http.StatusOK, item)
}

func DeleteItem(c *gin.Context) {
	var item models.Item
	id := c.Param("id")

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if item.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this item"})
		return
	}

	config.DB.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
}
