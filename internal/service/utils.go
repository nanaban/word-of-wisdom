package service

import (
	"encoding/binary"
	"io"
	"net"
)

// ReadMessage reads a message from the connection.
func ReadMessage(conn net.Conn) (msg []byte, err error) {
	// Read the message length.
	var length uint64
	if err = binary.Read(conn, binary.BigEndian, &length); err != nil {
		return
	}

	// Read the message.
	msg = make([]byte, length)
	_, err = io.ReadFull(conn, msg)
	return
}

// WriteMessage writes a message to the connection.
func WriteMessage(conn net.Conn, msg []byte) (err error) {
	// Write the message length.
	if err = binary.Write(conn, binary.BigEndian, uint64(len(msg))); err != nil {
		return
	}

	// Write the message.
	_, err = conn.Write(msg)
	return
}
