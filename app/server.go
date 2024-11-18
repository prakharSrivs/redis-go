package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/database"
	command_handlers "github.com/codecrafters-io/redis-starter-go/internals"
)

const (
	PING = "PING"
	ECHO = "ECHO"
	SET  = "SET"
	GET  = "GET"
)

const (
	BULK_STRING   string = "$"
	SIMPLE_STRING string = "+"
	ARRAY         string = "*"
)

var store = database.CreateNewStore()

func main() {
	l, er := net.Listen("tcp", "127.0.0.1:6379")
	if er != nil {
		logger("Error in listening for TCP Connections "+er.Error(), "red")
		os.Exit(1)
	}

	logger("Listening for TCP Connections ", "green")

	for {
		conn, err := l.Accept()
		if err != nil {
			logger("Error in waiting for TCP Connection"+err.Error(), "red")
			os.Exit(1)
		}
		logger("TCP Client Detected", "green")
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	for {
		b := make([]byte, 128)
		n, err := conn.Read(b)
		if err != nil {
			fmt.Println("Error while accepting Input for the connection", err.Error())
			return
		}
		message := string(b[:n])
		messageArray := strings.Split(message, "\r\n")[1:]
		commandsArray := make([]string, 0)

		for i := 1; i < len(messageArray); i = i + 2 {
			commandsArray = append(commandsArray, messageArray[i])
		}
		fmt.Println(commandsArray)
		command_handlers.COMMANDS[commandsArray[0]](commandsArray, conn, store)
	}
}

func logger(message string, color string) {
	red := "\033[31m"
	green := "\033[32m"
	reset := "\033[0m"

	switch color {
	case "red":
		fmt.Println(red, message, reset)
	case "green":
		fmt.Println(green, message, reset)
	}

}
