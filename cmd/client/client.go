package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const defaultAddress = "localhost:6379"

func main() {
	conn, err := net.Dial("tcp", defaultAddress)

	if err != nil {
		fmt.Println("Error dial", err.Error())
		return
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(defaultAddress + " >")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Error Write", err.Error())
			return
		}

		buff := make([]byte, 512)
		len, err := conn.Read(buff)
		if err != nil {
			fmt.Println("Error Read", err.Error())
			return
		}
		rsp := string(buff[0:len])
		fmt.Println(rsp)
	}
}
