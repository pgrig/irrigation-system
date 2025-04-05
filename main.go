package main

import "github.com/gin-gonic/gin"

func main() {
  router := gin.Default()
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  router.GET("/version", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "version": "v0.1.0",
    })
  })
  router.Run() 
}