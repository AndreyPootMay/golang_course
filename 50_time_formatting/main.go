package main

import (
	"fmt"
	"time"
)

func main() {
	// Redifinir print
	p := fmt.Println()

	// Obtain the time
	t := time.Now()

	// Imprimir
	p("1: ", t.Format(time.RFC3339))

	// Parsing time
	t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:44+00:00")
	
	p("2: ", t1)
	p("3: ", t.Format("3:04PM"))
	p("4: ", t.Format("Mon Jan _2 15:04:05 2006"))
	p("5: ", t.Format("2006-01-02T15:04:05.999999-07:00"))

	form := "3 04 PM"

	t2, e := time.Parse(form, "8 41 PM")
	p("6: ", t2)

	fmt.Printf("7: %d-%02d-%02dT%02d:%02d:%02d-00:00\n",
	t.Year(), t.Month(), t.Day(),
	t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p("8: ", e)
}