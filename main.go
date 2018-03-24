package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func HandleConn(conn net.Conn) {
	fmt.Println("on cline connection!")
	conn.Write([]byte("fuck you!!!\n"))
	conn.Write([]byte("is very good!\n"))

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			conn.Close()
		}
		fmt.Print(string(buf[:n]))
	}
	//conn.Write(buf)
	//fmt.Println(string(buf[:n])) //将接受的内容都读取出来。

}

func main() {

	fmt.Println("socket server start, port 8080!!!")
	addr := "0.0.0.0:8080"
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close() //close listener

	for {
		conn, err := listener.Accept() //conn 接收链接
		if err != nil {
			log.Fatal(err)
		}

		go HandleConn(conn)

	}
	fmt.Println("socket server stop!")
}
