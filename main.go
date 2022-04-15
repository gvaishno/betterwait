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
	betterwait "betterwait/pkg/engine"
	"flag"
	"os"
)

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
	betterwait.Betterwait(host, port, try, quiet)
}
