package main

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func init() {
	fmt.Println("SMTP 전송")
}

func main90() {
	m := gomail.NewMessage()
	m.SetHeader("From", "personar95@naver.com")
	m.SetHeader("To", "personar95@gmail.com", "personar95@naver.com")
	m.SetAddressHeader("Cc", "personar95@naver.com", "name")
	m.SetHeader("Subject", "이메일 테스트입니다2.")
	m.SetBody("text/html", "첨부파일이 있습니다.\r\n확인바랍니다.\r\n")
	m.Attach(`./sample.zip`) // 첨부파일

	d := gomail.NewDialer("smtp.naver.com", 465, "personar95", "password")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("메일전송!")
}
