// model.go

package main

import (
	"database/sql"
)

type Shipping struct {
	Id                   string `json:"id,omitempty"`
	Ship_id              string `json:"ship_id,omitempty"`
	Shipping_type        string `json:"ship_type,omitempty"`
	Destination_zipcode  string `json:"destination_zipcode,omitempty"`
	Weight_in_kilogramas string `json:"weight_in_kilogramas,omitempty"`
	Ship_price_per_kg    string `json:"ship_price_per_kg,omitempty"`
	Ship_totalprice      string `json:"ship_totalprice,omitempty"`
}

const getShippingDataForQuotationQuery = `SELECT
id,
ship_id,
shipping_type,
$1::NUMERIC AS weight_in_kilogramas,
CASE
    WHEN 10 >= $1 THEN price_per_kg_until_10kg
    WHEN 15 >= $1 THEN price_per_kg_until_15kg
    WHEN 20 >= $1 THEN price_per_kg_until_20kg
    WHEN 25 >= $1 THEN price_per_kg_until_25kg
    WHEN 30 >= $1 THEN price_per_kg_until_30kg
    WHEN 35 >= $1 THEN price_per_kg_until_35kg
    WHEN 40 >= $1 THEN price_per_kg_until_40kg
    WHEN 50 >= $1 THEN price_per_kg_until_50kg
    WHEN 60 >= $1 THEN price_per_kg_until_60kg
    WHEN 60 <  $1 THEN price_per_kg_above_60kg
    END AS ship_price_per_kg,
CASE
    WHEN 10 >= $1 THEN GREATEST(minimum_price, CAST(price_per_kg_until_10kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    WHEN 15 >= $1 THEN GREATEST(minimum_price, CAST(price_per_kg_until_15kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    WHEN 20 >= $1 THEN GREATEST(minimum_price, CAST(price_per_kg_until_20kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    WHEN 25 >= $1 THEN GREATEST(minimum_price, CAST(price_per_kg_until_25kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    WHEN 30 >= $1 THEN GREATEST(minimum_price, CAST(price_per_kg_until_30kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    WHEN 35 >= $1 THEN GREATEST(minimum_price, CAST(price_per_kg_until_35kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    WHEN 40 >= $1 THEN GREATEST(minimum_price, CAST(price_per_kg_until_40kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    WHEN 50 >= $1 THEN GREATEST(minimum_price, CAST(price_per_kg_until_50kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    WHEN 60 >= $1 THEN GREATEST(minimum_price, CAST(price_per_kg_until_60kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    WHEN 60 <  $1 THEN GREATEST(minimum_price, CAST(price_per_kg_above_60kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    END AS ship_totalprice
FROM 
shipping 
WHERE 
ship_id=$2 AND
(initial_zipcode <= $3 AND final_zipcode >= $3)`

func (s *Shipping) getShippingDataForQuotation(db *sql.DB) error {
	return db.QueryRow(getShippingDataForQuotationQuery, s.Weight_in_kilogramas, s.Ship_id, s.Destination_zipcode).Scan(&s.Id, &s.Ship_id, &s.Shipping_type, &s.Weight_in_kilogramas, &s.Ship_price_per_kg, &s.Ship_totalprice)
}

const getShippingDataForFullQuotationQuery = `SELECT
id,
ship_id,
shipping_type,
$2 AS destination_zipcode,
$1::NUMERIC AS weight_in_kilogramas,
CASE
    WHEN 10 >= $1 THEN price_per_kg_until_10kg
    WHEN 15 >= $1 THEN price_per_kg_until_15kg
    WHEN 20 >= $1 THEN price_per_kg_until_20kg
    WHEN 25 >= $1 THEN price_per_kg_until_25kg
    WHEN 30 >= $1 THEN price_per_kg_until_30kg
    WHEN 35 >= $1 THEN price_per_kg_until_35kg
    WHEN 40 >= $1 THEN price_per_kg_until_40kg
    WHEN 50 >= $1 THEN price_per_kg_until_50kg
    WHEN 60 >= $1 THEN price_per_kg_until_60kg
    WHEN 60 <  $1 THEN price_per_kg_above_60kg
    END AS ship_price_per_kg,
CASE
    WHEN 10 >= $1 THEN GREATEST(minimum_price, CAST(price_per_kg_until_10kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    WHEN 15 >= $1 THEN GREATEST(minimum_price, CAST(price_per_kg_until_15kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    WHEN 20 >= $1 THEN GREATEST(minimum_price, CAST(price_per_kg_until_20kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    WHEN 25 >= $1 THEN GREATEST(minimum_price, CAST(price_per_kg_until_25kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    WHEN 30 >= $1 THEN GREATEST(minimum_price, CAST(price_per_kg_until_30kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    WHEN 35 >= $1 THEN GREATEST(minimum_price, CAST(price_per_kg_until_35kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    WHEN 40 >= $1 THEN GREATEST(minimum_price, CAST(price_per_kg_until_40kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    WHEN 50 >= $1 THEN GREATEST(minimum_price, CAST(price_per_kg_until_50kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    WHEN 60 >= $1 THEN GREATEST(minimum_price, CAST(price_per_kg_until_60kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    WHEN 60 <  $1 THEN GREATEST(minimum_price, CAST(price_per_kg_above_60kg * CAST($1 AS DECIMAL(10,3)) AS DECIMAL(10,2)))
    END AS ship_totalprice

FROM 
shipping 
WHERE 
  (initial_zipcode <= $2 AND final_zipcode >= $2)`

func getShippingDataForFullQuotation(db *sql.DB, weight string, destination_zipcode string) ([]Shipping, error) {

	rows, err := db.Query(getShippingDataForFullQuotationQuery, weight, destination_zipcode)

	if err != nil {
		return nil, err
	}

	shippings := []Shipping{}

	defer rows.Close()

	for rows.Next() {
		var s Shipping
		if err := rows.Scan(&s.Id, &s.Ship_id, &s.Shipping_type, &s.Destination_zipcode, &s.Weight_in_kilogramas, &s.Ship_price_per_kg, &s.Ship_totalprice); err != nil {
			return nil, err
		}
		shippings = append(shippings, s)
	}

	return shippings, nil
}
