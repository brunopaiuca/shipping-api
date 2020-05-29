// main_test.go

package main

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM shipping")
	a.DB.Exec("ALTER SEQUENCE shipping_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS shipping
(
	id SERIAL,
	ship_id  VARCHAR(2) NOT NULL,
    initial_zipcode VARCHAR(8) NOT NULL,
    final_zipcode VARCHAR(8) NOT NULL,
    CONSTRAINT shipping_pkey PRIMARY KEY (id)
)`
