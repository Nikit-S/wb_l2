package ctime

import (
	"fmt"

	_ "github.com/beevik/ntp"
)

func PrintTime() error {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	fmt.Printf("Time %s\n", time)
	return nil
}
