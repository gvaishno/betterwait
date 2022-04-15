package main

import "testing"

func TestGoodCases(t *testing.T) {
	var goodtests = []struct {
		host  string
		port  string
		try   int
		quiet bool
		want  bool
	}{
		{"google.com", "80", 1, true, true},
		{"google.com", "80", 2, true, true},
		{"192.168.0.1", "80", 1, true, true},
	}

	for _, test := range goodtests {
		host := &test.host
		port := &test.port
		try := &test.try
		quiet := &test.quiet

		if !betterwait(host, port, try, quiet) {
			t.Errorf("betterwait(%q, %q, %d, %t) = %t, want %t", test.host, test.port, test.try, test.quiet, false, test.want)
		}
	}
}

func TestBadCases(t *testing.T) {
	var badtests = []struct {
		host  string
		port  string
		try   int
		quiet bool
		want  bool
	}{
		{"http://google.com", "80", 1, true, false},
		{"http://google.com:80", "80", 1, true, false},
		{"http://google.com:80", "", 1, true, false},
		{"http://google.com", "", 1, true, false},
		{"google.com:80", "80", 1, true, false},
		{"google.com", "", 1, true, false},
		{"192.168.0.1:80", "", 1, true, false},
		{"192.168.0.1", "", 1, true, false},
	}

	for _, test := range badtests {
		host := &test.host
		port := &test.port
		try := &test.try
		quiet := &test.quiet

		if betterwait(host, port, try, quiet) {
			t.Errorf("betterwait(%q, %q, %d, %t) = %t, want %t", test.host, test.port, test.try, test.quiet, true, test.want)
		}
	}
}
