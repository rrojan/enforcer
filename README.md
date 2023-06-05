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
name string `enforce:"required min:2 max:20 matches:^[A-Z][a-z]+(?: [A-Z][a-z]+)*"`
```  

---

### Simple validations:
- `required`
- string length (`between`, `min`, `max`)
- limits for int / float and derivatives (`between`, `min`, `max`)
- `match` (match emails, passwords, phone numbers, or your own custom regex patterns)
- `enum` (enforce enum options for string, int, etc)
  

<img width="883" alt="image" src="https://github.com/rrojan/enforcer/assets/59971845/d8df7c8d-6ead-46d7-8a35-279f015eb814">
<img width="747" alt="image" src="https://github.com/rrojan/enforcer/assets/59971845/335e505a-4205-4a3b-8a42-8d6815c78aeb">


### Custom validations:
- Use custom validations like below

```
type ProductReq struct {
  Title       string `json:"Title" enforce:"custom:productTitleTemplate"`
  Price       int    `json:"price" enforce:"custom:isNotOverpriced"`
  IsPublished int    `json:"is_published"`
}	
```

```
req := ProductReq{}
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
