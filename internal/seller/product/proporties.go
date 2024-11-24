package product

type Visibility string

const (
	VisibilityAll                    Visibility = "All"
	VisibilityVisible                Visibility = "VISIBLE"
	VisibilityInvisible              Visibility = "INVISIBLE"
	VisibilityEmptyStock             Visibility = "EMPTY_STOCK"
	VisibilityNotModerated           Visibility = "NOT_MODERATED"
	VisibilityModerated              Visibility = "MODERATED"
	VisibilityDisabled               Visibility = "DISABLED"
	VisibilityStateFailed            Visibility = "STATE_FAILED"
	VisibilityReadyToSupply          Visibility = "READY_TO_SUPPLY"
	VisibilityValidationStatePending Visibility = "VALIDATION_STATE_PENDING"
	VisibilityValidationStateFail    Visibility = "VALIDATION_STATE_FAIL"
	VisibilityValidationStateSuccess Visibility = "VALIDATION_STATE_SUCCESS"
	VisibilityToSupply               Visibility = "TO_SUPPLY"
	VisibilityInSale                 Visibility = "IN_SALE"
	VisibilityRemovedFromSale        Visibility = "REMOVED_FROM_SALE"
	VisibilityBanned                 Visibility = "BANNED"
	VisibilityOverpriced             Visibility = "OVERPRICED"
	VisibilityCriticallyOverpriced   Visibility = "CRITICALLY_OVERPRICED"
	VisibilityEmptyBarcode           Visibility = "EMPTY_BARCODE"
	VisibilityBarcodeExists          Visibility = "BARCODE_EXISTS"
	VisibilityQuarantine             Visibility = "QUARANTINE"
	VisibilityArchived               Visibility = "ARCHIVED"
	VisibilityOverpricedWithStock    Visibility = "OVERPRICED_WITH_STOCK"
	VisibilityPartialApproved        Visibility = "PARTIAL_APPROVED"
)

type InfoProperty struct {
	ProductId int64  `json:"product_id,omitempty"`
	OfferId   string `json:"offer_id,omitempty"`
	Sku       int64  `json:"sku,omitempty"`
}

type InfoProperties struct {
	ProductId []int64  `json:"product_id,omitempty"`
	OfferId   []string `json:"offer_id,omitempty"`
	Sku       []int64  `json:"sku,omitempty"`
}

type ShortInfoProperty struct {
	Visibility Visibility `json:"visibility"`
	OfferId    []string   `json:"offer_id,omitempty"`
	ProductId  []int64    `json:"product_id,omitempty"`
}
