package command_handlers

import (
	"fmt"
	"net"
	"strconv"
	"time"

	Serializer "github.com/codecrafters-io/redis-starter-go"
	"github.com/codecrafters-io/redis-starter-go/database"
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

var COMMANDS = map[string]func([]string, net.Conn, *database.Store){
	PING: ping,
	ECHO: echo,
	SET:  set,
	GET:  get,
}

func ping(args []string, conn net.Conn, store *database.Store) {
	conn.Write([]byte("+PONG\r\n"))
}

func echo(args []string, conn net.Conn, store *database.Store) {
	if len(args) == 1 {
		fmt.Println("Invalid use of ECHO")
		return
	}
	echoMessage := string(args[1])
	conn.Write([]byte(Serializer.EncodeBulkString(echoMessage)))
}

func get(args []string, conn net.Conn, store *database.Store) {
	val, err := store.Get(args[1])
	if err != nil {
		conn.Write([]byte("$-1\r\n"))
		return
	}
	conn.Write([]byte(Serializer.EncodeBulkString(val)))
}

func set(args []string, conn net.Conn, store *database.Store) {
	store.Set(args[1], args[2])
	if len(args) > 3 {
		timeMs, err := strconv.Atoi(args[4])
		if err != nil {
			timeMs = 100
		}
		go func(key string, expiry int) {
			time.Sleep(time.Duration(expiry) * time.Millisecond)
			store.Remove(key)
		}(args[1], timeMs)
	}
	conn.Write([]byte(Serializer.EncodeSimpleString("OK")))
}
