package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 1. 기본
// func main88() {
// 	r := gin.Default()
// 	r.GET("/hello", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"name": "Hugo",
// 		})
// 	})
// 	r.Run(":8889") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }

// 2. 각 HTTP메서드별 처리
// func main88() {
// 	router := gin.Default()

// 	router.GET("/someGet", getting) // 이 뒤는 모두 핸들러(함수)
// 	router.POST("/somePost", posting)
// 	router.PUT("/somePut", putting)
// 	router.DELETE("/someDelete", deleting)
// 	router.PATCH("/somePatch", patching)
// 	router.HEAD("/someHead", head)
// 	router.OPTIONS("/someOptions", options)

// 	router.Run(":8889")
// }

// 3. HTML 렌더링
func main88() {
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tpml", gin.H{
			"title": "Gin Framework",
		})
	})
	router.Run(":8889")
}
