package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rrojan/enforcer/enforcer"
)

type SignupReq struct {
	// Name -> Enforce `required` and `length` between 2 chars and 10 chars
	// Email -> Enforce `required` and `pattern` matches email
	// Phone -> Enforce `pattern` matches custom regex
	// Password -> Enforce `required`, `min` char value, `max` char value and `match` for password validity
	//     (We can also use `between` but this shows how we can use min / max separately)
	// UserType -> Enforce `enum` which can be "admin" or "user"
	Name     string `json:"name" enforce:"required between:2,10"`
	Email    string `json:"email" enforce:"required match:email"`
	Phone    string `json:"phone" enforce:"match:^[0-9\\-]{7,12}$"`
	Password string `json:"password" enforce:"required min:6 max:32 match:password"`
	Age      int    `json:"age" enforce:"between:2,10 exclude:9 enum:3,4"`
	// UserType string `json:"user_type" enforce:"required enum:admin,user"`
}

type ProductReq struct {
	Title       string `json:"Title" enforce:"custom:productTitleTemplate"`
	Price       int    `json:"price" enforce:"custom:isNotOverpriced,isEvenNumber"`
	IsPublished int    `json:"is_published"`
}

func main() {
	router := gin.Default()

	// Example usage of enforcer in a generic user signup controller
	router.POST("/signup", SignupController)
	// Example usage of custom enforcer function
	router.POST("/products", ProductCreateController)

	router.Run(":6969")
}

func SignupController(c *gin.Context) {
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

	c.JSON(200, gin.H{"message": "Signed up user successfully"})
}

func somePriceValidationQuery() int {
	return 1000
}

func ProductCreateController(c *gin.Context) {
	req := ProductReq{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("\n%#v\n", req)

	customEnforcements := enforcer.CustomEnforcements{
		{
			"productTitleTemplate": func(productTitle string) bool {
				isValid := true // validation logic
				return isValid
			},
			"productTitleTemplate2": func(productTitle string) bool {
				isValid := false // validation logic
				return isValid
			},
			"isNotOverpriced": func(priceStr string) bool {
				fmt.Println(priceStr, "wwww")
				price, _ := strconv.Atoi(priceStr)
				isValid := price < somePriceValidationQuery()
				return isValid
			},
			"isEvenNumber": func(priceStr string) bool {
				price, _ := strconv.Atoi(priceStr)
				fmt.Println(price)
				return price%2 == 0
			},
		},
	}

	errors := enforcer.CustomValidator(req, customEnforcements)

	if len(errors) > 0 {
		c.JSON(400, gin.H{"errors": errors})
		return
	}

	// Process the valid request here
	c.JSON(200, gin.H{"message": "Created product successfully"})
}
