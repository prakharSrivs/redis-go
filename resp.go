package Serializer

import "strconv"

func EncodeBulkString(message string) string {
	length := len(message)
	encodedMessage := "$" + strconv.Itoa(length) + "\r\n" + message + "\r\n"
	return encodedMessage
}

func EncodeSimpleString(message string) string {
	encodedMessage := "+" + message + "\r\n"
	return encodedMessage
}
