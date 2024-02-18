package receipt_manager

import item "receipt_manager/item"

type Receipt struct {
	Retailer      string       `json:"retailer"`
	PurchaseDate  string       `json:"purchaseDate"`
	PurchaseTime  string       `json:"purchaseTime"`
	Items         []item.Item  `json:"items"`
	Total         string       `json:"total"`
}