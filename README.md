# Enforcer
## Simplified validation for Go applications

---

<WIP>

Enforcer simplifies the tedious validation process in Go applications. Forget messy validation code, enforcer is here to enforce your will with a few Go tags and maps.


### See `main.go` for example gin application using Enforcer

### Usage:
- Use ``enforce`` to validate enforcements

E.g. 
```
name string `enforce:"required between:2,64 matches:^[A-Z][a-z]+(?: [A-Z][a-z]+)*"`
```  

---

### Simple validations:
- `required`
- string length (`between`, `min`, `max`)
- limits for int / float and derivatives (`between`, `min`, `max`)
- `match` (match emails, passwords, phone numbers, or your own custom regex patterns)
- `enum` (enforce enum options for string, int, etc)

```
type SignupReq struct {
	// Name -> Enforce `required` and `length` between 2 chars and 10 chars
	// Email -> Enforce `required` and `pattern` matches email
	// Phone -> Enforce `pattern` matches custom regex
	// Password -> Enforce `required`, `min` char value, `max` char value and `match` for password validity
	//     (We can also use `between` but this shows how we can use min / max separately)
	// UserType -> Enforce `enum` which can be "admin" or "user"
	Name  string `json:"name" enforce:"required between:2,10"`
	Email string `json:"email" enforce:"required match:email"`
	Phone string `json:"phone" enforce:"match:^[0-9\\-]{7,12}$"`
	Password string `json:"password" enforce:"required min:6 max:64 match:password"`
	Age int `json:"age" enforce:"min:18"`
	UserType string `json:"user_type" enforce:"required enum:admin,user"`
}
```

```
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
```

### Custom validations:
- Use custom validations like below

```
type ProductReq struct {
  Title       string `enforce:"required custom:productTitleTemplate"`
  Price       int    `enforce:"required custom:isNotOverpriced min:1000"`
  IsPublished int
}	
```

```
req := ProductReq{}
if err := c.ShouldBindJSON(&req); err != nil {
  c.JSON(400, gin.H{"error": err.Error()})
  return
}
customEnforcements := []map[string]func(string) bool{
  {
    "productTitleTemplate": func(productTitle string) bool {
      isValid := true // validation logic
      return isValid
    },
    "isNotOverpriced": func(priceStr string) bool {
      price, _ := strconv.Atoi(priceStr)
      isValid := price < somePriceValidationQuery()
      return isValid
    },
  },
}
errors := enforcer.CustomValidator(req, customEnforcements)
```
