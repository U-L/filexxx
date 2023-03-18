package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main()  {
	v := viper.New()
	v.SetConfigFile("store-api/user-web/nacos-dev.yaml")
	if err := v.ReadInConfig();err !=nil{
		panic(any(err))
	}
	router := gin.Default()
	router.GET("/file", func(c *gin.Context) {
		c.File("/file.go")
	})
	err := router.Run(":8081")
	if err != nil {
		return
	}
}
