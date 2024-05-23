package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Connect to the TCP server
	conn, err := net.Dial("tcp", "localhost:6565")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Create a reader for user input
	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	for {
		// Read user input
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		// Trim newline characters
		input = strings.TrimSpace(input)

		// Send the command to the server
		_, err = conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Println("Error sending command to server:", err)
			return
		}

		// Read the response from the server
		response, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading response from server:", err)
			return
		}

		// Print the server's response
		fmt.Println(response)
	}
}
