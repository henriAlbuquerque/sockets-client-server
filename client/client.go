package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

var clienteNome string

func main() {

	fmt.Println("Qual seu nome: ")
	clienteNome = inputTextCliente()

	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error connection:", err.Error())
	}
	defer connection.Close()

	_, err = connection.Write([]byte(clienteNome))
	if err != nil {
		fmt.Println("Error write:", err.Error())
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("Bate Papo:")
	fmt.Println()

	for {
		go readMsg(connection)
		sendMsg(connection)
	}

}

func sendMsg(connection net.Conn) {
	msg := inputTextMsg()

	if msg != "" {
		_, err := connection.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error write:", err.Error())
		}
	}

	return

}

func readMsg(connection net.Conn) {
	buffer := make([]byte, 1024)
	defer connection.Close()
	for {
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}

		msg := string(buffer[:mLen])

		fmt.Println(msg)
	}
}

func inputTextCliente() string {
	text := ""
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	text = scanner.Text()

	if text == "" {
		return "Cliente"
	}

	return text

}

func inputTextMsg() string {
	text := ""
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	text = scanner.Text()

	return text

}
