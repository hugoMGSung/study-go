package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	fmt.Println("오랜만에 초기화, MySQL 삭제!")
}

func main() {
	db, err := sql.Open("mysql", "madang:madang@tcp(127.0.0.1:3306)/madangdb")
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Panic(err)
	}

	defer tx.Rollback()

	// Prepared Statement 생성
	stmt, err := db.Prepare("DELETE FROM book WHERE bookid = ?")
	checkError(err)
	defer stmt.Close()

	// Prepared Statement 실행
	_, err = stmt.Exec(11) //Placeholder 파라미터 순서대로 전달
	checkError(err)

	err = tx.Commit()
	if err != nil {
		log.Panic(err)
	}

	getData()
}
