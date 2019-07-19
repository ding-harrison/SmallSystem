package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

const (
	PORT            = ":5000"
	NUM_CONNECTIONS = 100
)

var wg sync.WaitGroup

func createConnection(i int) {
	conn, err := net.Dial("tcp4", PORT)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect with server")
	}
	fmt.Fprintf(conn, "Created connection as %d\n", i)
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Fprintf(os.Stdout, "%s", status)
	wg.Done()
}

func main() {
	wg.Add(NUM_CONNECTIONS)
	for i := 0; i < NUM_CONNECTIONS; i++ {
		go createConnection(i)
	}
	wg.Wait()
}
