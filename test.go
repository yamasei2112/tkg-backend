package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	login()
}

func login() {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	var (
		UserID   string
		Userpass string
	)
	fmt.Println("UserIDの入力:")
	fmt.Scan(&UserID)
	fmt.Println("passwordの入力:")
	fmt.Scan(&Userpass)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("データベース接続成功")

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	var name string
	query := "SELECT * FROM users WHERE id = 'john1111'"
	err = db.QueryRow(query, UserID).Scan(&name)
	if err != nil {
		log.Fatal("ユーザー情報の取得に失敗しました: ", err)
	}

	fmt.Printf("ユーザー名: %s\n", name)
}
