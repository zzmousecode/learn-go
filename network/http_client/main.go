package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
)

const (
	processNum        = 1
	connNumPerProcess = 1
	reqNumPerConn     = 2
	host              = "127.0.0.1:12000"
	uri               = "/api"
)

var wg sync.WaitGroup

func main() {
	wg.Add(processNum)
	num := processNum
	for num > 0 {
		go do_task()
		num--
	}
	wg.Wait()
	fmt.Println("Over")
}
func do_task() {
	defer wg.Done()
	var conns []net.Conn
	sum := connNumPerProcess
	for sum > 0 {
		// 建立连接
		conn, err := net.Dial("tcp", host)
		if err != nil {
			fmt.Println("dial error:", err)
			conn.Close()
			return
		}
		conns = append(conns, conn)
		reqNum := reqNumPerConn
		for reqNum > 0 {
			// 发送请求, http 1.0 协议
			fmt.Fprintf(conn, "GET "+uri+" HTTP/1.1\r\n")
			fmt.Fprintf(conn, "Host: "+host+"\r\n")
			fmt.Fprintf(conn, "Accept: */*\r\n")
			if reqNum == 1 {
				fmt.Fprintf(conn, "Cookie: =\r\n")
			} else {
				fmt.Fprintf(conn, "Cookie: =\r\n")
			}
			fmt.Fprintf(conn, "Cookie: =\r\n")
			fmt.Fprintf(conn, "Connection: keep-alive\r\n\r\n")
			// 读取response
			// 读取response
			var sb strings.Builder
			buf := make([]byte, 2218)
			n, err := io.ReadFull(conn, buf)
			if err != nil {
				if err != io.EOF && err != io.ErrUnexpectedEOF {
					fmt.Println("read error:", err)
				}
			}
			sb.Write(buf[:n])
			// 显示结果
			fmt.Println("response:", sb.String())
			fmt.Println("total response size:", sb.Len())
			reqNum--
		}
		sum--
	}
}
