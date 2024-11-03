package models

type ShippingAddress struct {
	ID              int64   `db:"id"`
	UserID          int64   `db:"user_id"`
	AddressSourceID *string `db:"address_source_id"`
	AddressSource   string  `db:"address_source"`
	Name            string  `db:"name"`
	Country         string  `db:"country"`
	CountryCode     string  `db:"country_code"`
	Province        string  `db:"province"`
	City            string  `db:"city"`
	District        string  `db:"district"`
	AddressNote     *string `db:"address_note"`
}
