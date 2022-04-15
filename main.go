/*
Copyright © 2022 The betterwait authors. All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.


Binary package betterwait is a easy and highly effective solution to test and
wait on the availability of a TCP host and port for developers and
DevOps engineers.

Usage:
	betterwait [options]

Options:
	-help		Show this help message.
    -host	    Specify a host address.
	-port		Specify a port number.
	-try		Specify max tries before giving up.
	-quiet		Quiet mode.
*/

package main

import (
	"errors"
	"flag"
	"log"
	"net"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"
)

// Parse the Ip address
func ParseIP(s string) (string, error) {
	ip, _, err := net.SplitHostPort(s)
	if err == nil {
		return ip, nil
	}

	ip2 := net.ParseIP(s)
	if ip2 == nil {
		return "", errors.New("invalid IP")
	}

	return ip2.String(), nil
}

// Is IP address valid or not
func isValidIp(ip *string) bool {

	var validIP bool

	ipaddr := net.ParseIP(*ip)

	if ipaddr == nil {
		validIP = false
	} else {
		validIP = true
	}

	return validIP
}

// Check if the host is a IP address
func isHostIPaddress(host *string) (bool, string) {
	var isIP bool
	// Get the IP address in the host address.
	ip, _ := ParseIP(*host)

	// Check if the IP address is valid.
	if ip != "" {
		isIP = true
	} else {
		isIP = false
	}

	return isIP, ip
}

// Check host scheme.
func isHostScheme(host *string) bool {
	var trackScheme bool
	// Get the Scheme in the host address.

	// If the host is a IP address.
	isIP, ip := isHostIPaddress(host)
	if isIP {
		// Check if ip is valid
		if isValidIp(&ip) {
			trackScheme = false
		} else {
			trackScheme = true
		}
	} else {
		u, err := url.Parse(*host)
		if err != nil {
			log.Fatal(err)
		}

		// Check is the Scheme has . using regex.
		re := regexp.MustCompile(`\.`)
		if u.Scheme != "" && re.MatchString(u.Scheme) {
			// Its not a valid Scheme.
			trackScheme = false
		} else if u.Scheme == "" {
			// Its not a valid Scheme.
			trackScheme = false
		} else {
			// Its a valid Scheme.
			trackScheme = true
		}
	}
	return trackScheme
}

// Check is the host address has : suffix.
func isHostPort(host *string) bool {
	var isNum bool
	// Get the port in the host address.
	_, port, _ := net.SplitHostPort(*host)

	// Check if the port is a integer.
	_, err := strconv.Atoi(port)
	if err != nil {
		isNum = false
	} else {
		isNum = true
	}

	return isNum
}

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
func betterwait(host *string, port *string, try *int, quiet *bool) bool {
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
	if isHostScheme(host) {
		if !q {
			log.Println("ERROR: The host address has scheme prefix. Please use -port instead.")
			flag.Usage()
		}
		res = false
		return res
	}

	// Check if the host address has a port suffix.
	if isHostPort(host) {
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

func main() {
	// Define the flags
	var help = flag.Bool("help", false, "Show help")

	host := flag.String("host", "", "Specify a host address.")
	port := flag.String("port", "80", "Specify a port number.")
	try := flag.Int("try", 10e8-1, "Specify max tries before giving up.")
	quiet := flag.Bool("q", false, "Quiet mode.")

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	// Enable command-line parsing
	flag.Parse()

	// Call the betterwait function to test and wait for the host address.
	betterwait(host, port, try, quiet)
}
