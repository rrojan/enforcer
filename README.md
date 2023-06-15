# Enforcer 



### Simplified validation for Go apps

<!-- <img src="https://github.com/rrojan/enforcer/assets/59971845/f3be7493-344d-41ab-a2eb-f69b980b030f" width="20%"> -->
<img src="https://github.com/rrojan/enforcer/assets/59971845/c07d6a27-931f-4d3c-a6aa-214fa9a39b99" alt="enforcer-gopher-logo" width="20%">


Enforcer simplifies the tedious validation process in Go applications. 

Forget messy boilerplate-ridden code, enforcer is here to enforce your will with a handful of simple Go tags!

---


### Installation:
```
go get -u github.com/rrojan/enforcer
```

  
### Basic Usage:
- Use ``enforce`` to validate enforcements

E.g.: `name` is a *required* field *between* 2-64 chars, and should be *"Spaced and Cased"*
```
type myStruct struct {
  name string `enforce:"required between:2,64 matches:^[A-Z][a-z]+(?: [A-Z][a-z]+)*"`
}
```

---
### Contents
1. [Simple Validations](#simple-validations)
    - [Validations list](#validations-list)
    - [Binding simple validations with `enforce`](#binding-simple-validations-with-enforce)
    - [Applying the validation](#applying-simple-validations)
2. [Setting Defaults & Prohibits](#setting-defaults-and-prohibits)
    - [Setting default values for common data types](#setting-default-values)
    - [Setting default time (custom / time now, after or before)](#setting-default-time)
    - [Prohibited fields](#prohibited-fields)
3. [Custom Validations](#custom-validations)
    - [Using `custom` to bind custom validations to a field](#using-custom-to-bind-validations)
    - [Applying custom validation](#applying-the-custom-validations)
4. [Single Variable Validation](#variable-validation)
5. [Example projects](#example-projects)

---


## Simple validations

### Validations list

- `required`: mark a field as required
- `between`: string length or numerical value limit
- `min`: Minimum char length for string or minimum value for numeric type
- `max`: Maximum char length for string or maximum value for numeric type
- `match`: match emails, passwords, phone numbers, or your own custom regex patterns
- `enum`: enforce enum options for string, int, etc
- `exclude`: check whether value is in a list of excluded values
- `wordCount`: limit the wordcount of a string input
- `default`: add a default value in case not provided to the field
- `prohibit`: make sure a field is empty (user input cannot populate a struct field)

### Binding simple validations with enforce

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

  AuthUID string `enforce:"prohibit"`
}
```

#### Applying simple validations

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


## Setting Defaults and Prohibits

Enforcer allows you to set default values for struct fields. Default values take over in case values aren't provided while validating the struct.

### Setting Default Values
```
type User struct {
    Email     string `enforce:"required"`
    Username  string `enforce:"default:Anonymous"`
    UserType  int    `enforce:"enum:0,1,2 default:0"
    Score     float  `enforce:"default:5.0 between:0,10"`
}
```

### Setting Default Time
Time can be set to a custom value by default in the format "YYYY-MM-DD HH;MM;SS +TZHH:TZMM"

You can also set default to the current time using timeNow. Time before and after current date can be done using a semantic addition like `timeNow-1_day` or `timeNow+10_days`

```
type Coupon struct {
    ValidFrom    time.Time  `enforce:"default:2023-06-15 00;00;00 +5;45"`
    ActivatedAt  time.Time  `enforce:"default:timeNow"`
    NotifyAt     time.Time  `enforce:"default:timeNow+1_minute"`
    NextCoupon   time.Time `enforce:"default:timeNow+30_minutes"`
    ExpiresAt    time.Time  `enforce:"default:timeNow+5_days"`
}
```

Note that you must use semicolons `;` instead of `:` while referring to time and timezone offsets because of the way tag parsing in Go works.

### Prohibited Fields

There are certain cases where a field input must never be binded from user input, or where user input should not bypass the default value. In these cases, `prohibit` will reset the field to its corresponding zero or default value.

```
type User struct {
    Username  string  `enforce:"required"`
    Password  string  `enforce:"required match:password"`

    AuthUID   string  `enforce:"prohibit"` // Even if user provides the `AuthUID` field, it will be reset to null value
    UserType  string  `enforce:"prohibit default:user enum:user,admin"` // Using prohibit means that User won't be able to override default
}
```


## Custom validations:

### Using custom to bind validations

Use `custom:` and `enforcer.CustomValidator` to run multiple custom validators like below

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

#### Applying the custom validations

Use `enforcer.CustomValidator` to validate and run the custom bindings through a `enforcer.CustomEnforcements` map

Note that the argument of the CustomEnforcement function is always a string, regardless of what the actual field type might be
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

### Example Projects
- [Enforcer Examples](https://github.com/rrojan/enforcer-examples)
