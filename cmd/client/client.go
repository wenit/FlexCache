package main

import (
	"aedis/internal/protocol"
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
		// text := "set name a"
		data, _ := protocol.MarshalStr(text)
		_, err := conn.Write(data)
		if err != nil {
			fmt.Println("Error Write", err.Error())
			return
		}

		respReader := bufio.NewReader(conn)

		respData := protocol.UnMarshal(respReader)
		if len(respData) == 0 {
			fmt.Println("(nil)")
		} else {
			for _, v := range respData {
				fmt.Println(string(v))
			}
		}
	}
}
