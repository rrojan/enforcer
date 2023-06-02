package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rrojan/enforcer/enforcer"
)

type SignupReq struct {
	Name  string `json:"name" enforce:"required between:2,10"`
	Email string `json:"email" enforce:"required match:email"`
	Phone string `json:"phone" enforce:"required match:^[0-9\\-]{7,12}$"`
	Password string `json:"password" enforce:"min:6 max:64 match:password"`
}

func main() {
	router := gin.Default()

	// Example usage of enforcer in a generic user signup controller
	router.POST("/signup", func(c *gin.Context) {
		req := SignupReq{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// enforcer.Validate reads the `enforce:"..."` tags and applies enforcements
		errors := enforcer.Validate(req)

		if len(errors) > 0 {
			c.JSON(400, gin.H{"errors": errors})
			return
		}

		// Process the valid request here
		c.JSON(200, gin.H{"message": "Signed up user successfully"})
	})

	router.Run(":6969")
}
