package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jimmiepr/Gin-TDD/internal/service"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/getdata", service.GetData)
	}
	r.Run(":3000")
}
