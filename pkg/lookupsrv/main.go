package lookupsrv

import (
	"net"
	"fmt"
	"os"
)

func Main() {
	config := NewFromCmd()

	_, records, err := net.LookupSRV("", "", config.Host)
	if err != nil {
		println(err.Error())
		os.Exit(21)
	}

	var result string
	for _, record := range records {
		addr := fmt.Sprintf(config.Format, record.Target[:len(record.Target)-1])
		if result == "" {
			result = addr
		} else {
			result += config.Delimiter + addr
		}
	}

	fmt.Println(result)
	os.Exit(0)
}

