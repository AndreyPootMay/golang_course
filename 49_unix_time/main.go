package main

import (
	"fmt"
	"time"
)

func main() {
	// Obtains now
	now := time.Now()

	secs := now.Unix()

	nanos := now.UnixNano()

	fmt.Println("now:", now);

	// Divide over one millon
	millis := nanos / 1000000
	fmt.Println("Secs: ", secs)
	fmt.Println("Millis: ", millis)
	fmt.Println("Nanos: ", nanos)

	fmt.Println("time.Unix secs: ", time.Unix(secs, 0))
	fmt.Println("time.Unix nanos: ", time.Unix(0, nanos))
}