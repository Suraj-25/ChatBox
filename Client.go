package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "localhost: 9997")

	if err != nil {
		log.Fatalln(err)
	}

	for {
		scanA := bufio.NewScanner(os.Stdin)

		fmt.Print(" Please Enter Your Message : ")

		scanA.Scan()

		inpA := scanA.Text()

		_, err = conn.Write([]byte(inpA))
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Your Message Has Been sent Successfully :", inpA)

		buffA := make([]byte, 1400)

		dataSize, err := conn.Read(buffA)

		if err != nil {
			fmt.Println("The connection is closed, Please Try Again Later")
			return
		}

		dataA := buffA[:dataSize]

		fmt.Println("The Message Has Been Recieved Successfully : ", string(dataA))

	}
}
