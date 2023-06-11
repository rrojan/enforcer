# Enforcer
### Simplified validation for Go apps

<!-- <img src="https://github.com/rrojan/enforcer/assets/59971845/efd66255-8538-4d50-8d20-4212ba0f26e1" width="40%"> -->


Enforcer simplifies the tedious validation process in Go applications. 

Forget messy boilerplate-ridden validation code, enforcer is here to enforce your will with a few Go tags and maps!

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
  // Name is a required field with 2-64 char limit and should be "Spaced And Camel Cased"
  name string `enforce:"required between:2,64 matches:^[A-Z][a-z]+(?: [A-Z][a-z]+)*"`
}
```  

---



## Simple validations:
- `required`: mark a field as required
- `between`: string length or numerical value limit
- `min`: Minimum char length for string or minimum value for numeric type
- `max`: Maximum char length for string or maximum value for numeric type
- `match`: match emails, passwords, phone numbers, or your own custom regex patterns
- `enum`: enforce enum options for string, int, etc
- `exclude`: check whether value is in a list of excluded values
- `wordCount`: limit the wordcount of a string input
- `default`: add a default value in case not provided to the field

```
type SignupReq struct {
  // Name -> Enforce "required" and length "between" 2 chars and 10 chars
  Name  string    `json:"name"     enforce:"required between:2,10"`
  
  // Email -> Enforce "required" and pattern "match" for email
  Email string    `json:"email"    enforce:"required match:email"`
  
  // Phone -> Enforce pattern "match" for custom regex
  Phone string    `json:"phone"    enforce:"match:^[0-9\\-]{7,12}$"`
  
  // Age -> Enforce "min" and "max" signup age (number) to be in range 18-100 (we can use `between` for this as well)
  Age int         `json:"age"      enforce:"min:18 max:100"`
  
  // UserType -> Enforce "enum" which lets the value be "admin" or "user"
  UserType string `json:"type"     enforce:"required enum:admin,user"`
  
  // Bio -> Minimum of 3 "wordCount", maximum of 150, and a max" 256 character limit
  Bio string      `json:"bio"      enforce:"wordCount:3,150 max:256"
  
  // Password -> Enforce "required", "min" char limit, "max" char limit and "match" for password validity
  Password string `json:"password" enforce:"required match:password"`
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

// This is an array of all errors present
}
```



## Custom validations:

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
customEnforcements := enforcer.CustomEnforcements{
  {
    "productTitleTemplate": func(productTitle string) string {
      // Apply validation logic here
      isValid := true
      if !isValid {
        return "Product title does not match proper format"
      }
      return ""
    },
    "isNotOverpriced": func(priceStr string) string {
      price, _ := strconv.Atoi(priceStr)
      isValid := price < somePriceValidationQuery()
      if !isValid {
        return "Product is overpriced!"
      }
      return ""
    },
    "canUserSetPrice": func(priceStr string) string {
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

## Using *default*

Use ``enforce:"... default:someValue ..."`` to add a default value in case data is not provided to a struct field

```
type User struct {
  Name      string    `enforce:"required between:2,32"`
  UserType  string    `enforce:"default:user enum:admin,user"`
  IsActive  int       `enforce:"default:1 enum:0,1"
}
```
#### Applying the validation and setting defaults

This is done by the `enforcer.Validate` function, however when `default` is used in any struct field, you must provide the address of the struct, not the struct itself for it to work. Else the default will be set on a copy of the struct, not the original struct itself

```
u := User{}
c.ShouldBindJSON(&u) // Bind values to `u` from request data
errors := enforcer.Validate(&u)
```


## Variable validation

While not often used, variable validation can be performed by using the `enforcer.ValidateVar` function

```
myAge := 23
errors = enforcer.ValidateVar(myAge, "min:18 max:100")
```
