CREATE TABLE IF NOT EXISTS shipping
(
    id SERIAL,
    ship_id         VARCHAR(4) NOT NULL,
    shipping_type   VARCHAR(35) NOT NULL,
    area            VARCHAR(50) NOT NULL, 
    radius          VARCHAR(4) NOT NULL, 
    multiple_factor NUMERIC(2,1) NOT NULL, 
    initial_zipcode VARCHAR(8) NOT NULL, 
    final_zipcode   VARCHAR(8) NOT NULL, 
    minimum_price   NUMERIC(10,2) NOT NULL, 
    price_per_kg_until_10kg NUMERIC(10,2) NOT NULL, 
    price_per_kg_until_15kg NUMERIC(10,2) NOT NULL, 
    price_per_kg_until_20kg NUMERIC(10,2) NOT NULL, 
    price_per_kg_until_25kg NUMERIC(10,2) NOT NULL,
    price_per_kg_until_30kg NUMERIC(10,2) NOT NULL,
    price_per_kg_until_35kg NUMERIC(10,2) NOT NULL,
    price_per_kg_until_40kg NUMERIC(10,2) NOT NULL,
    price_per_kg_until_50kg NUMERIC(10,2) NOT NULL,
    price_per_kg_until_60kg NUMERIC(10,2) NOT NULL,
    price_per_kg_above_60kg NUMERIC(10,2) NOT NULL,
    CONSTRAINT shipping_pkey PRIMARY KEY (id)
);

CREATE UNIQUE INDEX get_shipping_data_per_zipcode_and_shipid ON shipping USING btree (ship_id, initial_zipcode, final_zipcode);
