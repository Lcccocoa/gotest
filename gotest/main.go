package main

import (
	"fmt"
	"gotest/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.User{}, &models.Todo{})

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	todo := r.Group("/todo")

	todo.GET("/", func(c *gin.Context) {
		todos := []models.Todo{}
		db.Find(&todos)
		c.JSON(200, gin.H{
			"list": todos,
		})
	})
	todo.POST("/add", func(c *gin.Context) {
		todo := &models.Todo{}
		c.BindJSON(todo)
		db.Create(todo)

		c.JSON(200, gin.H{
			"data": todo,
		})
	})
	todo.POST("/update", func(c *gin.Context) {
		todo := &models.Todo{}
		c.BindJSON(todo)
		fmt.Println(todo)

		db.Model(todo).Update("done", todo.Done)
		c.JSON(200, gin.H{
			"data": todo,
		})
	})
	todo.POST("/delete", func(c *gin.Context) {
		todo := &models.Todo{}
		c.BindJSON(todo)
		db.Delete(todo)

		c.JSON(200, gin.H{
			"data": todo,
		})
	})
	r.Run()
}

// 跨域请求中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
