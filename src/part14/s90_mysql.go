package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main90() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("DB 오류!", r)
		}
	}()

	db, err := sql.Open("mysql", "madang:madang@tcp(127.0.0.1:3306)/madangdb")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB연결 성공!")

	// INESRT 문
	result, err := db.Exec("INSERT INTO book VALUES (?, ?, ?, ?)", 12, "MySQL 학습", "한빛아카데미", 25000)
	if err != nil {
		//log.Fatal(err)
		panic(err)
	}

	// sql.Result.RowsAffected() 체크
	n, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%d row inserted", n)
	}

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

	// Prepared Statement 생성
	stmt, err := db.Prepare("UPDATE book SET price=? WHERE bookid=?")
	checkError(err)
	defer stmt.Close()

	// Prepared Statement 실행
	_, err = stmt.Exec(1, 7050) //Placeholder 파라미터 순서대로 전달
	checkError(err)
	_, err = stmt.Exec(9, 7550)
	checkError(err)

	rows1, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%3s\t%20s\t%10s\t%s\n", "순번", "책제목", "출판사", "가격")

	for rows1.Next() {
		err := rows.Scan(&bookid, &bookname, &publisher, &price)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%3d\t%20s\t%10s\t%d\n", bookid, bookname, publisher, price)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
