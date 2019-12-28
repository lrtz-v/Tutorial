package library

import (
	"fmt"
	"log"
	"testing"
)

func TestQuery(T *testing.T) {
	conn, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := conn.Query("SELECT data_type, data_status FROM offline_data_sync_status WHERE ID = 1417")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var data_type int
	var date_status int
	for rows.Next() {
		if err := rows.Scan(&data_type, &date_status); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("data_type:%d ,date_status:is %d\n", data_type, date_status)
	}
}

func TestQueryV2(T *testing.T) {
	conn, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := conn.Query("SELECT data_type, data_status FROM offline_data_sync_status WHERE ID = 1417")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var data_type int
	var date_status int
	for rows.Next() {
		if err := rows.Scan(&data_type, &date_status); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("data_type:%d ,date_status:is %d\n", data_type, date_status)
	}
}
