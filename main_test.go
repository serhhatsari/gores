package main

import (
	"bufio"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	go main()                   // start the server
	time.Sleep(2 * time.Second) // give the server some time to start

	// run the tests
	code := m.Run()

	// stop the server
	os.Exit(code)
}

func TestConnection(t *testing.T) {

	// try to connect to the server
	conn, err := net.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		t.Fatalf("Could not connect to server: %v", err)
	}

	// close the connection
	err = conn.Close()
	if err != nil {
		t.Fatalf("Could not close the connection: %v", err)
	}
}

func TestPing(t *testing.T) {

	// try to connect to the server
	conn, err := net.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		t.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()

	// send PING command
	_, err = conn.Write([]byte("*1\r\n$4\r\nPING\r\n"))
	if err != nil {
		t.Fatalf("Could not write to connection: %v", err)
	}

	// read response
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("Could not read from connection: %v", err)
	}

	// check response
	if response != "+PONG\r\n" {
		t.Fatalf("Unexpected response: %v", response)
	}
}

func TestSet(t *testing.T) {
	// try to connect to the server
	conn, err := net.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		t.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()

	// send SET command
	_, err = conn.Write([]byte("*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n"))
	if err != nil {
		t.Fatalf("Could not write to connection: %v", err)
	}

	// read response
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("Could not read from connection: %v", err)
	}

	// check response
	if response != "+OK\r\n" {
		t.Fatalf("Unexpected response: %v", response)
	}
}

func TestGet(t *testing.T) {
	// try to connect to the server
	conn, err := net.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		t.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()

	// send GET command
	_, err = conn.Write([]byte("*2\r\n$3\r\nGET\r\n$3\r\nkey\r\n"))
	if err != nil {
		t.Fatalf("Could not write to connection: %v", err)
	}

	// read response
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("Could not read from connection: %v", err)
	}

	// if the response contains multiple lines, read the next line
	if strings.HasPrefix(response, "$") {
		nextLine, err := reader.ReadString('\n')
		if err != nil {
			t.Fatalf("Could not read from connection: %v", err)
		}
		response += nextLine
	}

	// check response
	expectedResponse := "$5\r\nvalue\r\n"
	if response != expectedResponse {
		t.Fatalf("Unexpected response: got %v want %v", response, expectedResponse)
	}
}
