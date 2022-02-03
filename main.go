package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type datos struct {
	Key   string      `json:"key" binding:"required"`
	Value interface{} `json:"value" binding:"required"`
}

var personas map[string]interface{} = make(map[string]interface{})

func main() {
	router := config()
	router.Run(":8080")
}

func config() *gin.Engine {
	router := gin.Default()

	router.POST("/guardados", func(c *gin.Context) {
		var keyDatos datos

		if err := c.ShouldBindJSON(&keyDatos); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		personas[keyDatos.Key] = keyDatos.Value
		c.JSON(http.StatusOK, gin.H{"mensaje": fmt.Sprintf("se guardo %v", keyDatos.Value)})
	})

	router.GET("/guardados/:key", func(c *gin.Context) {
		key := c.Param("key")
		val, ok := personas[key]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{
				"mensaje": "No hay datos",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": val,
		})
	})

	router.DELETE("/guardados/:key", func(c *gin.Context) {
		key := c.Param("key")

		delete(personas, key)

		c.JSON(200, gin.H{"se borro": fmt.Sprint(key)})

	})

	router.PUT("/guardados/:key", func(c *gin.Context) {

		var nuevo datos

		key := c.Param("key")
		//val = personas[key]
		if err := c.ShouldBindJSON(&nuevo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		personas[key] = nuevo.Value
		c.JSON(http.StatusOK, gin.H{"mensaje": fmt.Sprintf("se guardo %v", nuevo.Value)})
	})

	return router
}
