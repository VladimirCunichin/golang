package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntpTime, err := ntp.Time("time.apple.com")
	if err != nil {
		println(err)
		os.Exit(1)
	}

	ntpTimeFormatted := ntpTime.Format(time.UnixDate)
	fmt.Printf("Network time: %v\n", ntpTime)
	fmt.Printf("Unix Date Network time: %v\n", ntpTimeFormatted)
}
