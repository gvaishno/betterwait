/*
Author: Vaishno Chaitanya
License: https://github.com/gvaishno/betterwait/LICENSE
Website: https://github.com/gvaishno/betterwait
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

// ConnectLoop tries to send tcp packets to the given address until it succeeds or times out.
func ConnectLoop(address *string, port *string, try *int, quiet *bool) {
	addr := *address
	p := *port
	t := *try
	q := *quiet

	if !q {
		log.Println("INFO: Connecting to", addr, "on port", p, "for", t, "tries...")
	}

	for i := 0; i < t; i += 1 {
		conn, err := net.Dial("tcp", net.JoinHostPort(addr, p))

		switch err {
		// Error is nil and the connection is established successfully. Then break the loop.
		case nil:
			if !q {
				log.Printf("INFO: Connected to %s:%v after %d tries.\n", addr, p, i+1)
			}

			conn.Close()
			return
		default:
			// Waiting for a while before trying again.
			if !q {
				log.Printf("INFO: Waiting %v second for %s:%v before trying again.", t-i, addr, p)
			}

			time.Sleep(time.Second * 1)
		}
	}
	// If the connection is not established within the given number of tries.
	fmt.Printf("ERROR: Failed to connect to %s:%v.", addr, p)
}

func main() {
	// Define the flags
	var help = flag.Bool("help", false, "Show help")

	host := flag.String("h", "", "Specify a host address.")
	port := flag.String("p", "80", "Specify a port number.")
	try := flag.Int("t", 10e8-1, "Specify max tries before giving up.")
	quiet := flag.Bool("q", false, "Quiet mode.")

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	// Enable command-line parsing
	flag.Parse()

	h := *host
	p := *port
	t := *try
	q := *quiet

	// If the host address is not specified.
	if *host == "" {
		log.Println("ERROR: You must specify an host address.")
		fmt.Println("")
		flag.Usage()
		return
	}

	hostaddr := &h
	portnum := &p
	trynum := &t
	quietmode := &q

	// Call the ConnectLoop function.
	ConnectLoop(hostaddr, portnum, trynum, quietmode)
}
