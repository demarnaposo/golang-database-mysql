package belajar_golang_database

import (
	"context"
	"fmt"
	"testing"
)


func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('2', 'Demar')"

	// ExecContext = untuk mengirim perintah yang bukan query data sql ke dalam database (insert, update, delete)

	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Data berhasil disimpan")
}