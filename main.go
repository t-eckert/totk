package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("America/New_York")
	launch := time.Date(2023, time.Month(5), 12, 0, 0, 0, 0, loc)
	countdown := int(launch.Sub(time.Now()).Hours() / 24)

	fmt.Printf("Legend of Zelda: Tears of the Kingdom comes out in %d days!\n", countdown)
}
