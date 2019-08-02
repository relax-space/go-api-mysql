package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/mysql?charset=utf8",
		os.Getenv("USER_NAME"), os.Getenv("PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("PORT")))
	fmt.Println(os.Getenv("USER_NAME"), os.Getenv("PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("PORT"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	for index := 0; index < 100; index++ {
		if err = db.Ping(); err != nil {
			fmt.Printf("ping is err:%v \n", err)
			time.Sleep(1 * time.Second)
			continue
		}
	}

	fmt.Println("ping is ok")
	time.Sleep(1 * time.Hour)
}
