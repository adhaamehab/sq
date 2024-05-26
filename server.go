package sq

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Server struct {
	Queue *Queue
	Port  string

	listener net.Listener
}

// NewServer will create a new server with a default port of 6565
func NewServer(port string) *Server {
	if port == "" {
		port = "6565"
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	return &Server{
		Queue:    &Queue{},
		Port:     "6565",
		listener: listener,
	}
}

// Start will start the server and listen for incoming connections.
// Connection commands will be passed to queue handler.
func (s *Server) Start() {
	fmt.Println("Server started on port", s.Port)
	defer s.listener.Close()
	for {
		// Accept a connection
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}

		go s.handleConnection(conn)
	}
}

// handleConnection will read the message from the connection and pass it to the queue
func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		input = strings.TrimSpace(input)

		command, args := parseCommand(input)

		fmt.Println("Received command:", command)
		// Handle the command
		switch command {
		case PINGCMD:
			conn.Write([]byte("PONG\n"))
		case APPEND:
			if len(args) == 0 {
				conn.Write([]byte("Missing argument\n"))
				continue
			}

			item := Item(strings.Join(args, " "))
			s.Queue.Append(item)
			conn.Write([]byte("OK\n"))

		default:
			fmt.Println("Unknown command:", command)
			conn.Write([]byte("Unknown command\n"))
		}
	}
}
