package main

import (
	"fmt"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

var users = make([]net.Conn, 0, 10)

func main() {

	fmt.Println("Servidor rodando...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()

	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Aguardando usuários...")

	fmt.Println()
	fmt.Println()
	fmt.Println("Bate Papo:")
	fmt.Println()

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		users = append(users, connection)

		buffer := make([]byte, 1024)
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		fmt.Println("		(" + string(buffer[:mLen]) + " conectado)")
		sendMsgOtherClients(connection, "		("+string(buffer[:mLen])+" conectado)")

		go readMsg(connection, string(buffer[:mLen]))
	}

}

func sendMsgOtherClients(connection net.Conn, msgCompleta string) {
	for i := 0; i < len(users); i++ {

		if users[i] != connection {
			_, err := users[i].Write([]byte(msgCompleta))
			if err != nil {
				fmt.Println("Error write:", err.Error())
			}
		}
	}
}

func readMsg(connection net.Conn, clienteNome string) {
	buffer := make([]byte, 1024)
	defer connection.Close()
	for {
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			removeUser(connection, clienteNome)
			break
		}

		msg := string(buffer[:mLen])

		fmt.Println("	" + clienteNome + ": " + msg)
		msgCompleta := clienteNome + ": " + msg
		sendMsgOtherClients(connection, msgCompleta)
	}
}

func removeUser(connection net.Conn, clienteNome string) {
	i := 0
	for ; i < len(users); i++ {
		if users[i] == connection {
			break
		}
	}

	fmt.Println("		(" + clienteNome + " saiu)")

	users = append(users[:i], users[i+1:]...)
	msg := "		(" + clienteNome + " saiu)"
	sendMsgOtherClients(connection, msg)

}
