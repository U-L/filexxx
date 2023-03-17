package main

import "github.com/gin-gonic/gin"

func main()  {
	router := gin.Default()
	router.GET("/file", func(c *gin.Context) {
		c.File("/file.go")
	})
	err := router.Run(":8081")
	if err != nil {
		return
	}
}
