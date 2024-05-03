package main

import (
	"bufio"
	"net"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	go main()                   // start the server
	time.Sleep(2 * time.Second) // give the server some time to start

	// run the tests
	code := m.Run()

	// here you would add code to stop the server

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

func TestHandleClient(t *testing.T) {

	// try to connect to the server
	conn, err := net.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		t.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()

	// send a request to the server
	_, err = conn.Write([]byte("*1\r\n$4\r\nPING\r\n"))
	if err != nil {
		t.Fatalf("Could not write to connection: %v", err)
	}

	// read the response from the server
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("Could not read from connection: %v", err)
	}

	// check the response
	expectedResponse := "+PONG\r\n"
	if response != expectedResponse {
		t.Fatalf("Unexpected response: got %v want %v", response, expectedResponse)
	}
}
