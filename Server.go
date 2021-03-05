package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	fmt.Println("The Server is currently listening to 9997")

	listener, err := net.Listen("tcp", "localhost:9997")

	if err != nil {
		log.Fatalln(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("The New Connection Has Been Estabilished ", conn.LocalAddr())

		go listenConn(conn)
	}
}

func listenConn(conn net.Conn) {
	for {
		buffA := make([]byte, 1400)

		dataSize, err := conn.Read(buffA)

		if err != nil {
			fmt.Println("The Estabilished Connection Has Been Closed")
			return
		}

		dataA := buffA[:dataSize]

		fmt.Println("The Message Has Been Recieved : ", string(dataA))

		scanA := bufio.NewScanner(os.Stdin)

		fmt.Print("Please Enter Your Message Here :")

		scanA.Scan()

		inpA := scanA.Text()

		_, err = conn.Write([]byte(inpA))
		if err != nil {
			log.Fatalln(err)
		}
	}
}
