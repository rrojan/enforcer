package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rrojan/enforcer/enforcer"
)

type ReqStruct struct {
	Name  string `json:"name" enforce:"required between:2,10"`
	Email string `json:"email" enforce:"required match:email"`
	Phone string `json:"phone" enforce:"required match:^[0-9\\-]{7,12}$"`
}

func main() {
	router := gin.Default()

	router.POST("/test", func(c *gin.Context) {
		req := ReqStruct{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		errors := enforcer.Validate(req)
		if len(errors) > 0 {
			c.JSON(400, gin.H{"errors": errors})
			return
		}

		// Process the valid request here
		c.JSON(200, gin.H{"message": "Request validates successfully"})
	})

	router.Run(":6969")
}
