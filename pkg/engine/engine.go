package betterwait

import (
	betterwait "betterwait/pkg/service"
	"flag"
	"log"
	"net"
	"strconv"
	"time"
)

// ConnectLoop tries to send tcp packets to the given address until it succeeds or times out.
func ConnectLoop(address *string, port *string, try *int, quiet *bool) bool {
	addr := *address
	p := *port
	t := *try
	q := *quiet

	if !q {
		log.Println("INFO: Connecting to", addr, "on port", p, "for", t, "tries...")
	}

	for i := 0; i < t; i += 1 {
		// Timeout in seconds.
		timeout := time.Duration(1 * time.Second)
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(addr, p), timeout)

		switch err {
		// Error is nil and the connection is established successfully. Then break the loop.
		case nil:
			if !q {
				log.Printf("INFO: Connected to %s after %d tries.\n", addr, i+1)
			}

			conn.Close()
			return true
		default:
			// Waiting for a while before trying again.
			if !q {
				log.Printf("INFO: Waiting %v second for %s before trying again.", t-i, addr)
			}

			time.Sleep(time.Second * 1)
		}
	}
	// If the connection is not established within the given number of tries.
	return false
}

// Entry point of the program.
func Betterwait(host *string, port *string, try *int, quiet *bool) bool {
	var res bool

	h := *host
	p := *port
	t := *try
	q := *quiet

	hostaddr := &h
	portnum := &p
	trynum := &t
	quietmode := &q

	// If the host address is not specified.
	if *host == "" {
		if !q {
			log.Println("ERROR: You must specify an host address.")
			flag.Usage()
		}
		res = false
		return res
	}

	//

	// Check if the host address has a scheme.
	if betterwait.IsHostScheme(host) {
		if !q {
			log.Println("ERROR: The host address has scheme prefix. Please use -port instead.")
			flag.Usage()
		}
		res = false
		return res
	}

	// Check if the host address has a port suffix.
	if betterwait.IsHostPort(host) {
		if !q {
			log.Println("ERROR: The host address has port suffix. Please use -port instead.")
			flag.Usage()
		}
		res = false
		return res
	}

	// Check if the port is a integer.
	if _, err := strconv.Atoi(*port); err != nil {
		if !q {
			log.Println("ERROR: The port you specified is not a number. Please enter a valid port number.")
			flag.Usage()
		}
		res = false
		return res
	}

	// Call the ConnectLoop function.
	res = ConnectLoop(hostaddr, portnum, trynum, quietmode)
	return res
}
