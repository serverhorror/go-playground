package test

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestPostgreSQLConnection(t *testing.T) {
	db, err := sql.Open("postgres", "postgresql://reader:reader@vie-bio-postgres.eu.boehringer.com/postgres?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	rows, err := db.Query("SELECT 1")
	if err != nil {
		t.Fatal(err)
	}

	for rows.Next() {
		var n int
		err := rows.Scan(&n)
		if err != nil {
			t.Fatal(err)
		}
	}
}
