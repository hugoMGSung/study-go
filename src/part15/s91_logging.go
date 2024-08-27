package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	fmt.Println("로그 처리")
}

var myLogger *log.Logger

func main91() {
	log.SetFlags(3) // 로그처리 방법 설정(간단출력) 0~3
	log.Println("로그출력")

	// 로그파일 출력
	fpLog, err := os.OpenFile("./log/logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	// myLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	myLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	//...
	run()

	myLogger.Println("프로그램 종료")
}

func run() {
	myLogger.Print("테스트!")
}
