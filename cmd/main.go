package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	fmt.Println("Server is running on :8080")
	if err := r.Run(":8080"); err != nil {
		fmt.Println(err.Error())
	}
}
