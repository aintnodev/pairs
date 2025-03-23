package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type pair struct {
	Id      string  `json:"id" gorm:"primaryKey"`
	Aname   string  `json:"aname"`
	Bname   string  `json:"bname"`
	Alname  string  `json:"alname"`
	Blname  string  `json:"blname"`
	Percent float32 `json:"percent"`
}

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("database/example.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	err = db.AutoMigrate(&pair{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

func welcomeMsg(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to Pairs API"})
}

func getPair(c *gin.Context) {
	query := strings.ToLower(strings.TrimSpace(c.Query("query")))
	limStr := c.DefaultQuery("limit", "10")
	order := c.DefaultQuery("order", "desc")

	limit, err := strconv.Atoi(limStr)
	if err != nil {
		limit = 10
	}

	if order != "asc" {
		order = "desc"
	}

	var pairs []pair
	dbQuery := db

	if query != "" {
		dbQuery = dbQuery.Where("alname LIKE ? OR blname LIKE ?", "%"+query+"%", "%"+query+"%")
	}

	if err := dbQuery.Order(fmt.Sprintf("percent %s", order)).Limit(limit).Find(&pairs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get pairs"})
	}

	c.JSON(http.StatusOK, pairs)
}

func addPair(c *gin.Context) {
	var newPair pair
	if err := c.BindJSON(&newPair); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid input"})
		return
	}

	newPair.Id = uuid.NewString()
	newPair.Aname = strings.TrimSpace(newPair.Aname)
	newPair.Bname = strings.TrimSpace(newPair.Bname)
	newPair.Alname = strings.ToLower(strings.TrimSpace(newPair.Aname))
	newPair.Blname = strings.ToLower(strings.TrimSpace(newPair.Bname))
	newPair.Percent = avgPercent(newPair.Alname, newPair.Blname)

	var extPair pair
	if err := db.Where("(alname = ? AND blname = ?) OR (alname = ? AND blname = ?)", newPair.Alname, newPair.Blname, newPair.Blname, newPair.Alname).First(&extPair).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Pair already exist"})
		return
	}

	if err := db.Create(&newPair).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create couple"})
	}

	c.IndentedJSON(http.StatusCreated, newPair)
}

func deletePair(c *gin.Context) {
	id := c.Param("id")

	if err := db.Delete(&pair{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete pair"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pair deleted successfully"})
}
