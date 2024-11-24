package seller_tests

import (
	"context"
	"github.com/ghost-gopher/ozon-api/internal/seller"
	"github.com/ghost-gopher/ozon-api/internal/seller/product"
	"github.com/ghost-gopher/ozon-api/internal/seller/utils"
	"testing"
)

const (
	ApiKey   = "-"
	ClientId = "-"
)

func TestInfo(t *testing.T) {
	ctx := context.Background()
	slr, _ := seller.New(ApiKey, ClientId)

	pdt, err := slr.Product().Info(ctx, product.InfoProperty{
		OfferId: "123456789",
	})
	if err != nil {
		t.Fatal(err)
	}

	if pdt.Id > 0 {
		t.Logf("product_id: %d, product_name: %s", pdt.Id, pdt.Name)
	} else {
		t.Log("products list empty")
	}

}

func TestList(t *testing.T) {
	ctx := context.Background()
	slr, _ := seller.New(ApiKey, ClientId)

	pdts, err := slr.Product().List(ctx, utils.Sort{Filter: product.ShortInfoProperty{
		Visibility: product.VisibilityAll,
	},
		Limit: 3,
	})
	if err != nil {
		t.Fatal(err)
	}

	if pdts.Total > 0 {
		items := *pdts.Items.(*[]product.ShortInfoObject)
		if len(items) > 0 {
			ids := make([]int64, 0, len(items))
			for _, item := range items {
				ids = append(ids, item.ProductId)
			}

			t.Logf("product_ids: %d", ids)
		} else {
			t.Log("products list empty")
		}

		t.Logf("total: %d", pdts.Total)
		t.Logf("last_id: %s", pdts.LastId)
	} else {
		t.Log("products list empty")
	}

}

func TestInfoList(t *testing.T) {
	ctx := context.Background()
	slr, _ := seller.New(ApiKey, ClientId)

	pdts, err := slr.Product().InfoList(ctx, product.InfoProperties{
		ProductId: []int64{123456788, 123456789},
	})
	if err != nil {
		t.Fatal(err)
	}

	items := *pdts.Items.(*[]product.InfoObject)
	if len(items) > 0 {
		ids := make([]int64, 0, len(items))
		for _, item := range items {
			ids = append(ids, item.Id)
		}

		t.Logf("product_ids: %d", ids)
	} else {
		t.Log("products list empty")
	}
}
