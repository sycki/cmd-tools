package portcheck

import (
	"net"
	"time"
	"fmt"
	"os"
)

func Main() {
	config := NewFromCmd()

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(config.Host, config.Port), time.Second*3)
	if err != nil {
		fmt.Println("close")
		println(err.Error())
		os.Exit(31)
	}

	fmt.Println("open")

	conn.Close()
	os.Exit(0)
}

