package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// データベースへの接続を設定
	db, err := sql.Open("mysql", "ユーザー名:パスワード@tcp(ホスト名:ポート)/データベース名")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// データベース接続を確認
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// INSERT文を実行
	query := `INSERT INTO employees (name, position, salary) VALUES (?, ?, ?)`
	_, err = db.Exec(query, "Alice", "Developer", 60000.00)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Employee added successfully")
}
