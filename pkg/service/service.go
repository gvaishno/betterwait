package betterwait

import (
	"errors"
	"log"
	"net"
	"net/url"
	"regexp"
	"strconv"
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
func IsValidIp(ip *string) bool {

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
func IsHostIPaddress(host *string) (bool, string) {
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
func IsHostScheme(host *string) bool {
	var trackScheme bool
	// Get the Scheme in the host address.

	// If the host is a IP address.
	isIP, ip := IsHostIPaddress(host)
	if isIP {
		// Check if ip is valid
		if IsValidIp(&ip) {
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
func IsHostPort(host *string) bool {
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
