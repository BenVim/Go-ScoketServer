package main

import (
	"fmt"
	"net"
	"log"
	"time"
)

func HandleConn(conn net.Conn)  {
	fmt.Println("on cline connection!")
	conn.Write([]byte("fuck you!!!\n"))
	conn.Write([]byte("is very good!\n"))
	time.Sleep(time.Minute/3)
	conn.Close()
}

func main()  {

	fmt.Println("socket server start, port 8080!!!")
	addr:= "0.0.0.0:8080"
	listener, err :=net.Listen("tcp", addr)

	if err != nil{
		log.Fatal(err)
	}

	defer listener.Close() //close listener

	for{
		conn, err := listener.Accept() //conn 接收链接
		if err !=nil{
			log.Fatal(err)
		}

		go HandleConn(conn)

	}
	fmt.Println("socket server stop!")
}
