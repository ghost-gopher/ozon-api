package product

import "time"

type InfoObject struct {
	Id                    int64     `json:"id"`
	Name                  string    `json:"name"`
	OfferId               string    `json:"offer_id"`
	Barcode               string    `json:"barcode"`
	BuyboxPrice           string    `json:"buybox_price"`
	CreatedAt             time.Time `json:"created_at"`
	Images                []string  `json:"images"`
	MarketingPrice        string    `json:"marketing_price"`
	MinOzonPrice          string    `json:"min_ozon_price"`
	OldPrice              string    `json:"old_price"`
	PremiumPrice          string    `json:"premium_price"`
	Price                 string    `json:"price"`
	RecommendedPrice      string    `json:"recommended_price"`
	MinPrice              string    `json:"min_price"`
	Vat                   string    `json:"vat"`
	Visible               bool      `json:"visible"`
	PriceIndex            string    `json:"price_index"`
	Images360             []string  `json:"images360"`
	ColorImage            string    `json:"color_image"`
	PrimaryImage          string    `json:"primary_image"`
	State                 string    `json:"state"`
	ServiceType           string    `json:"service_type"`
	FboSku                int64     `json:"fbo_sku"`
	FbsSku                int64     `json:"fbs_sku"`
	CurrencyCode          string    `json:"currency_code"`
	IsKgt                 bool      `json:"is_kgt"`
	Rating                string    `json:"rating"`
	IsDiscounted          bool      `json:"is_discounted"`
	HasDiscountedItem     bool      `json:"has_discounted_item"`
	Barcodes              []string  `json:"barcodes"`
	UpdatedAt             time.Time `json:"updated_at"`
	Sku                   int64     `json:"sku"`
	DescriptionCategoryId int64     `json:"description_category_id"`
	TypeId                int64     `json:"type_id"`
	IsArchived            bool      `json:"is_archived"`
	IsAutoarchived        bool      `json:"is_autoarchived"`
}

type InfoSource struct {
	IsEnabled bool   `json:"is_enabled"`
	Sku       int64  `json:"sku"`
	Source    string `json:"source"`
}

type InfoStock struct {
	Coming   int32 `json:"coming"`
	Present  int32 `json:"present"`
	Reserved int32 `json:"reserved"`
}

type ShortInfoObject struct {
	ProductId    int64  `json:"product_id"`
	OfferId      string `json:"offer_id"`
	IsFboVisible bool   `json:"is_fbo_visible"`
	IsFbsVisible bool   `json:"is_fbs_visible"`
	Archived     bool   `json:"archived"`
	IsDiscounted bool   `json:"is_discounted"`
}
