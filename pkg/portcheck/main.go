package portcheck

import (
	"net"
	"time"
	"fmt"
	"os"
)

func Main() {
	config := NewFromCmd()

	timeout := time.Second * time.Duration(config.Timeout)

	conn, err := net.DialTimeout(config.Protocol, net.JoinHostPort(config.Host, config.Port), timeout)
	if err != nil {
		fmt.Println("close")
		println(err.Error())
		os.Exit(31)
	}
	defer conn.Close()

	fmt.Println("open")

	os.Exit(0)
}

