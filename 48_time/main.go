package main

import (
	"fmt"
	"time"
)

func main() {
	// re-define print function
	p := fmt.Println

	now := time.Now()

	// Print now
	p("1: ", now)

	// Obtain then
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p("2: ", then)

	p("2: ", then.Year())
	p("3: ", then.Month())
	p("5: ", then.Day())
	p("6: ", then.Hour())
	p("7: ", then.Minute())
	p("8: ", then.Second())
	p("9: ", then.Nanosecond())
	p("a: ", then.Location())
	p("b: ", then.Weekday())
	p("c: ", then.Before(now))
	p("d: ", then.After(now))
	p("e: ", then.Equal())

	diff := now.Sub(then)
	p("f: ", diff)
	p("g: ", diff.Hours())
	p("h: ", diff.Minutes())
	p("i: ", diff.Seconds())
	p("j: ", diff.Nanoseconds())
	p("k: ", then.Add(diff))
	p("l: ", then.Add(-diff))
}