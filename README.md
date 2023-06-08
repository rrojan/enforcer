# Enforcer
## Simplified validation for Go apps


Enforcer simplifies the tedious validation process in Go applications. 

Forget messy validation code, enforcer is here to enforce your will with a few Go tags and maps!


---


### Installation:
```
go get -u github.com/rrojan/enforcer
```

  
### Basic Usage:
- Use ``enforce`` to validate enforcements

E.g. 
```
type myStruct struct {
  name string `enforce:"required between:2,64 matches:^[A-Z][a-z]+(?: [A-Z][a-z]+)*"`
}
```  

---

### Simple validations:
- `required`: mark a field as required
- `between`, `min`, `max`: string length and numerical limits ()
- `match`: match emails, passwords, phone numbers, or your own custom regex patterns
- `enum`: enforce enum options for string, int, etc
- `exclude`: check whether value is in a list of excluded values
- `wordCount`: limit the wordcount of a string input

```
type SignupReq struct {
  // Name -> Enforce `required` and `length` between 2 chars and 10 chars
  Name  string    `json:"name"     enforce:"required between:2,10"`
  
  // Email -> Enforce `required` and `pattern` matches email
  Email string    `json:"email"    enforce:"required match:email"`
  
  // Phone -> Enforce `pattern` matches custom regex
  Phone string    `json:"phone"    enforce:"match:^[0-9\\-]{7,12}$"`
  
  // Password -> Enforce `required`, `min` char value, `max` char value and `match` for password validity
  //     (We can also use `between` but this shows how we can use min / max separately)
  Password string `json:"password" enforce:"required match:password min:6 max:64"`
  
  // Age -> Enforce minimum signup age (number) to be 18
  Age int         `json:"age"      enforce:"min:18"`
  
  // UserType -> Enforce `enum` which can be "admin" or "user"
  UserType string `json:"type"     enforce:"required enum:admin,user"`
  
  // Bio -> Minimum of 3 words, maximum of 150 words, and a 256 character limit
  Bio string      `json:"bio"      enforce:"wordCount:3,150 max:256" 
}
```

#### Applying the validation

```
req := SignupReq{}

// This example uses Gin for request binding, but enforcer can be used on its own as well
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

Use `enforcer.CustomValidator` to run multiple custom validators like below

```
type ProductReq struct {
  // Enforce a `productTitleTemplate` validation for title
  Title       string `enforce:"required custom:productTitleTemplate"`
  
  // Enforce multiple custom validators for price by chaining it with a comma
  Price       int    `enforce:"required custom:isNotOverpriced,canUserSetPrice min:1000"`

  // 0 -> Draft, 1 -> Published
  IsPublished int    `enforce:"enum:0,1"`
}	
```

#### Applying the validation

```
req := ProductReq{}
if err := c.ShouldBindJSON(&req); err != nil {
  c.JSON(400, gin.H{"error": err.Error()})
  return
}
customEnforcements := []map[string]func(string) bool{
  {
    "productTitleTemplate": func(productTitle string) string {
      // Apply validation logic here
      isValid := true
      if !isValid {
        return "Product title does not match proper format"
      }
      return ""
    },
    "isNotOverpriced": func(priceStr string) bool {
      price, _ := strconv.Atoi(priceStr)
      isValid := price < somePriceValidationQuery()
      if !isValid {
        return "Product is overpriced!"
      }
      return ""
    },
    "canUserSetPrice": func(priceStr string) bool {
      price, _ := strconv.Atoi(priceStr)
      isValid := priceRoleValidate(price)
      if !isValid {
        return "User does not have authorization to set price in this range"
      }
      return ""
    },
  },
}
errors := enforcer.CustomValidator(req, customEnforcements) // Array of error messages
```
