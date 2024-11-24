package product

import (
	"context"
	"github.com/ghost-gopher/ozon-api/internal"
	"github.com/ghost-gopher/ozon-api/internal/seller/utils"
	"net/http"
)

const (
	EntityPath             = "product"
	APIProductInfoPath     = "info"
	APIProductListPath     = "list"
	APIProductInfoListPath = "info/list"
)

type Product struct {
	client internal.Client
}

func New(client internal.Client) *Product {
	return &Product{
		client: client,
	}
}

//	Product.List - load product list -> https://docs.ozon.ru/api/seller/#operation/ProductAPI_GetProductList
/*
	Version: v2

	Params:
		- ctx - context of request for example: [context.Background]
		- src - incoming parameter [utils.Sort]: [utils.Sort.Filter], [utils.Sort.LastId], [utils.Sort.Limit]

	Returns: slice [product.ShortInfoObject] or error
*/
func (p *Product) List(ctx context.Context, src utils.Sort) (*utils.List, error) {
	tgt := &utils.List{
		Items: &[]ShortInfoObject{},
	}
	rps, err := p.client.Request(ctx, http.MethodPost, utils.PathEndpointVersionTwo(EntityPath, APIProductListPath), src, tgt)
	if err != nil {
		return nil, err
	}

	return rps.Result.(*utils.List), nil
}

//	Product.Info - load one product info -> https://docs.ozon.ru/api/seller/#operation/ProductAPI_GetProductInfoV2
/*
	Version: v2

	Params:
		- ctx - context of request for example: [context.Background]
		- src - incoming parameter [product.InfoProperty]: [product.InfoProperty.ProductId], [product.InfoProperty.OfferId], [product.InfoProperty.Sku]

	Returns: [product.InfoObject] or error
*/
func (p *Product) Info(ctx context.Context, src InfoProperty) (*InfoObject, error) {
	rps, err := p.client.Request(ctx, http.MethodPost, utils.PathEndpointVersionTwo(EntityPath, APIProductInfoPath), src, &InfoObject{})
	if err != nil {
		return nil, err
	}

	return rps.Result.(*InfoObject), nil
}

//	Product.InfoList - load product list info -> https://docs.ozon.ru/api/seller/#operation/ProductAPI_GetProductInfoListV2
/*
	Version: v2

	Params:
		- ctx - context of request for example: [context.Background]
		- src - incoming parameter [product.InfoProperties]: [product.InfoProperties.ProductId], [product.InfoProperties.OfferId], [product.InfoProperties.Sku]

	Returns: [product.InfoObject] or error
*/
func (p *Product) InfoList(ctx context.Context, src InfoProperties) (*utils.List, error) {
	tgt := &utils.List{
		Items: &[]InfoObject{},
	}

	rps, err := p.client.Request(ctx, http.MethodPost, utils.PathEndpointVersionTwo(EntityPath, APIProductInfoListPath), src, tgt)
	if err != nil {
		return nil, err
	}

	return rps.Result.(*utils.List), nil
}
