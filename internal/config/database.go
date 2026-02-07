package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import driver (init only)
)

func ConnectDB() *sql.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/expense_db?parseTime=true"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Gagal Connect DB: ", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("DB Tidak Respon: ", err)
	}

	fmt.Println("Database Connected!")

	return db
}
