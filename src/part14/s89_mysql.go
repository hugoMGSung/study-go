package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main89() {
	db, err := sql.Open("mysql", "madang:madang@tcp(127.0.0.1:3306)/madangdb")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB연결 성공!")

	// 하나의 Row를 갖는 SQL 쿼리
	var name string
	err = db.QueryRow("SELECT bookname FROM madangdb.book where bookid = 2").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	// 여러 행
	var bookid int
	var bookname string
	var publisher string
	var price int32
	var query string = `SELECT 	bookid,
	bookname,
	publisher,
	price
	FROM book`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Printf("%3s\t%20s\t%10s\t%s\n", "순번", "책제목", "출판사", "가격")

	for rows.Next() {
		err := rows.Scan(&bookid, &bookname, &publisher, &price)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%3d\t%20s\t%10s\t%d\n", bookid, bookname, publisher, price)
	}
}
