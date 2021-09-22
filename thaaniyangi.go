package main

import (
	"fmt"

	"github.com/karnaprakash/thaaniyangi/noti"
)

func main() {
	sendMes := noti.Telereq("FIrst phase success")

	fmt.Println(sendMes)

}
