/**
 * Created with IntelliJ IDEA.
 * User: bogdan
 * Date: 1/26/13
 * Time: 8:32 PM
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
	"net/http"
	"bufio"
	"os"
	"net"
	"time"
)

func main() {
	service := ":8081"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)

	}


}

func handleConnection(conn net.Conn) {
	request, err := http.ReadRequest(bufio.NewReader(conn))

	checkError(err)
	conn.Write([]byte("HTTP/1.1 200 OK\r\nDate: " + time.Now().String() + "\r\nContent-Type: text/html\r\nServer: goserv\r\nConnection: close\r\n\r\n<html><p>erwertwertwertwr</p></html>"))
	conn.Close()

	fmt.Fprintf(os.Stderr, "Method %s", request.Method)


}

func checkError(err error) {
	if err != nil && err.Error() != "EOF" {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}  else {
		if err != nil && err.Error() == "EOF" {
			fmt.Fprintf(os.Stderr, "Another error: %s", err.Error())
		}
	}

}


