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

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	var (
		UserID   string
		Userpass string
	)
	fmt.Println("UserIDの入力")
	fmt.Scan(&UserID)
	fmt.Println("passwordの入力")
	fmt.Scan(&Userpass)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	println("mysql open success")
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

}
