package controllers

import (
	"testing"
)

func TestIPCheck(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{in: "127.0.0.1", want: "ipv4"},
		{in: "::1", want: "ipv6"},
		{in: "[288:8001:cd0a::11:1]", want: "ipv6"},
		{in: "[288:8001:cd0a::11:1]:80", want: "ipv6"},
		{in: "[::127.0.0.1]:80", want: "ipv6"},
		{in: "8.8.8.8", want: "ipv4"},
		{in: "1:2:3:4:5:6:7:8::1", want: "ipv6"},
		{in: "[1:2:3:4:5:6:7:8::1]:123", want: "ipv6"},
	}

	for _, c := range cases {
		actual := getIPVersion(c.in)
		if actual != c.want {
			t.Errorf("get result of %s fail, want %s but got %s\n", c.in, c.want, actual)
		}
	}

}
