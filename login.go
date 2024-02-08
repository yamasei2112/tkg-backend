package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// データベースへの接続を設定
	db, err := sql.Open("mysql", "root:yamasei2112@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	println("mysql open success")
	defer db.Close()

	// データベース接続を確認
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	// INSERT文を実行
}
