package main

import (
	"fmt"
	"testing"
)

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"10:11:12", true},
		{"1:14:52", true},
		{"0:-11:12", false},
		{"0:59:59", true},
		{"", false},
		{"11:22", false},
		{"aa:bb:cc", false},
		{"5:23:", false},
	}
	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			fmt.Printf("%v: %v, error should be nil", data.time, err)
		}
	}
}
