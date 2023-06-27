package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	port  int
	url   string
	https bool
)

func printStartLine() {
	startStr := strings.Repeat(">", 100)
	fmt.Println(startStr)
}

func printEndLine() {
	endStr := strings.Repeat("<", 100)
	fmt.Println(endStr)
}

func helloHandle(c *gin.Context) {
	// separator
	printStartLine()

	// headers
	fmt.Println("[Request Headers]:")
	for k, v := range c.Request.Header {
		fmt.Printf("%s: %s\n", k, v)
	}

	// body
	body, _ := c.GetRawData()
	fmt.Println("[Request Body]:\n", string(body))

	// response
	c.JSON(200, gin.H{"message": "Received Request!"})

	// separator
	printEndLine()
}

func registerHandler(c *gin.Context) {
	url := c.Request.URL.Path
	fmt.Println("Registering function for:", url)
	//r.Any(url, helloHandle)
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// pprof路由注册需要在最前面
	r.Any("/debug/pprof/*pprof", gin.WrapH(http.DefaultServeMux))

	r.Any(url, helloHandle)
	return r
}

func main() {
	// 解析参数
	flag.IntVar(&port, "port", 8080, "监听端口")
	flag.StringVar(&url, "url", "/", "监听path")
	flag.BoolVar(&https, "https", false, "https接口")
	flag.Parse()

	// 设置默认路由
	r := SetupRouter()

	// 拼接端口
	portStr := fmt.Sprintf(":%d", port)

	// run
	if https {
		r.RunTLS(portStr, "ca.crt", "ca.key")
	} else {
		r.Run(portStr)
	}
}
