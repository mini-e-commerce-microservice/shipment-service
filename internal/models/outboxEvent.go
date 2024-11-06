package models

type OutboxEvent struct {
	ID            int64   `db:"id"`
	AggregateType string  `db:"aggregatetype"`
	AggregateID   string  `db:"aggregateid"`
	Type          string  `db:"type"`
	Payload       any     `db:"payload"`
	TraceParent   *string `db:"trace_parent"`
}

type OutboxEventCourierRate struct {
	OrderID              int64                               `json:"order_id"`
	UserID               int64                               `json:"user_id"`
	DestinationAddressID int64                               `json:"destination_address_id"`
	OriginAddressID      int64                               `json:"origin_address_id"`
	CourierCode          string                              `json:"courier_code"`
	CourierServiceCode   string                              `json:"courier_service_code"`
	TotalAmount          float64                             `json:"total_amount"`
	PaymentMethodCode    string                              `json:"payment_method_code"`
	Items                []OutboxEventCourierRateProductItem `json:"items"`
}

type OutboxEventCourierRateProductItem struct {
	ID                int64   `db:"id" json:"id"`
	Name              string  `db:"name" json:"name"`
	Description       string  `db:"description" json:"description"`
	OrderID           int64   `db:"order_id" json:"order_id"`
	ProductItemID     int64   `db:"product_item_id" json:"product_item_id"`
	Qty               int32   `db:"qty" json:"qty"`
	UnitPrice         float64 `db:"unit_price" json:"unit_price"`
	TotalPrice        float64 `db:"total_price" json:"total_price"`
	Discount          float64 `db:"discount" json:"discount"`
	Weight            int32   `db:"weight" json:"weight"`
	PackageLength     float64 `db:"package_length" json:"package_length"`
	PackageWidth      float64 `db:"package_width" json:"package_width"`
	PackageHeight     float64 `db:"package_height" json:"package_height"`
	DimensionalWeight float64 `db:"dimensional_weight" json:"dimensional_weight"`
}
