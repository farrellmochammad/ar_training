package models


import (
	"time"

)


type Supplier struct {
	SupplierId        int          `json:"supplier_id"`
	Name      string       `json:"name"`
	Created *time.Time   `json:"created_at,omitempty"`
}


type SupplierStore struct {
	StoreId int `json:"store_id"`
	StoreName string `json:"store_name"`
    StoreAddress string `json:"store_address"`
	Created *time.Time `json:"created_at,omitempty"`
}
