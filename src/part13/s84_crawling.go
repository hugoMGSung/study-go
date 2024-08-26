package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const (
	urlRoot = "https://www.danawa.com/"
)

// 첫번째 방문 대상으로 원하는 url 파싱후 반환
func parseMainNodes(n *html.Node) bool {
	// must check for nil values
	if n.DataAtom == atom.A && n.Parent.DataAtom == atom.Li && n.Parent.Parent != nil {
		return scrape.Attr(n.Parent.Parent, "class") == "prod-list"
	}
	return false
}

// 에러체크 공통함수
func customErrCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// 동기화를 위한 작업 그룹 선언
var wg sync.WaitGroup

func main84() {
	// 메인 페이지 get방식 요청
	resp, err := http.Get(urlRoot) // request, response
	customErrCheck(err)

	//time.Sleep(5 * time.Second)

	// 요청 body 닫기
	defer resp.Body.Close()

	// 응답 데이터(html)
	root, err := html.Parse(resp.Body)
	customErrCheck(err)

	// ParseMainNodes 메서드 크롤링 대상 URL 추출
	urlList := scrape.FindAll(root, parseMainNodes)

	for _, link := range urlList {
		// 대상 URL 출력
		//fmt.Println("Check link", idx, link) // 1
		fmt.Println("TargetURL", scrape.Attr(link, "href")) // 2
		fileName := strings.Replace(scrape.Attr(link, "href"), "http://dasale.danawa.com/?controller=Goods&methods=lodingGoodsBlog&goodsSeq=", "", 1)
		//fmt.Println(fileName)

		// 작업 대기열 추가
		wg.Add(1) // Done 개수와 일치

		// 고루틴 시작
		go scrapContents(scrape.Attr(link, "href"), fileName)
	}

	wg.Wait() // 모든 작업이 종료될 때까지 대기
}

// URL 대상이 되는 페이지 파싱후 반환
func scrapContents(url string, fn string) {
	defer wg.Done() // 종료 알림

	resp, err := http.Get(url)
	customErrCheck(err)

	//time.Sleep(2 * time.Second)
	// 요청 body 닫기
	defer resp.Body.Close()

	// 응답데이터 html
	root, err := html.Parse(resp.Body)
	customErrCheck(err)

	// Respose 데이터 에 원하는 부분 파싱
	matchNode := func(n *html.Node) bool {
		return n.DataAtom == atom.Title
	}

	// 파일 스트림 생성
	file, err := os.OpenFile("./result/"+fn+".txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	customErrCheck(err)

	// 메서드 종료 파일 닫기
	defer file.Close()

	// 쓰기 버퍼
	w := bufio.NewWriter(file)

	// matchNode 메서드 사용, 원하는 노드 순회 출력
	for _, g := range scrape.FindAll(root, matchNode) {
		// 데이터 출력
		fmt.Println("result", scrape.Text(g))
		// 파싱 데이터 -> 버퍼에 기록
		w.WriteString(scrape.Text(g) + "\r\n")
	}

	w.Flush()
}
