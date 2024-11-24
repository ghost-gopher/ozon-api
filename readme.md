# Ozon Common API

Functional includes: 
- **Seller API** - for manage products, reports, sell strategies, warehouses more details [here](https://docs.ozon.ru/api/seller).

## How to start
### API
Get ```Client-Id``` and ```Api-Key``` in your seller profile [here](https://seller.ozon.ru/app/settings/api-keys?locale=en)

for begin add dependency to your project and ready done.
```bash
go get github.com/ghost-gopher/ozon-api
```

A simple example on how to use this library:

```Golang
package main

import (
	"context"
	"github.com/ghost-gopher/ozon-api/internal/seller"
	"github.com/ghost-gopher/ozon-api/internal/seller/product"
	"log"
)

const (
	ApiKey   = "-"
	ClientId = "-"
)

func main() {
	ctx := context.Background()
	slr, _ := seller.New(ApiKey, ClientId)

	pdt, err := slr.Product().Info(ctx, product.InfoProperty{
		Sku: 123456789,
	})
	if err != nil {
		log.Fatalf("error %s", err.Error())
	}

	if pdt.Id > 0 {
		log.Printf("product_id: %d, product_name: %s", pdt.Id, pdt.Name)
	} else {
		log.Printf("products list empty")
	}
}
```


## Contribution
If you need some endpoints ASAP, create an issue and list all the endpoints. I will add them to library soon.

Or you can implement them and contribute to the project. Contribution to the project is welcome.