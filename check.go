package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destination string, port string) string {
	address := destination + ":" + port // taking the input from values passed by the user
	timeout := time.Duration(
		5 * time.Second,
	) // duration after which it will be declared unreachable
	_, err := net.DialTimeout(
		"tcp",
		address,
		timeout,
	) // the output determines if unreachable or reachable
	var status string
	if err != nil {
		status = fmt.Sprintf(
			"[DOWN] %v is unreachable.\n",
			destination,
		) // Status if the domain is not reachable
	} else {
		status = fmt.Sprintf("[UP] %v is reachable.\n", destination) // Status if the doamin is reachable
	}
	return status // returning the status of the website
}
