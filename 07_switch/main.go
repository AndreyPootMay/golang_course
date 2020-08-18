package main

import (
	"fmt"
	"time"
)

func main()  {
	i := 2

	fmt.Print("Escribe ", i, " como ")

	switch i {
    case 1:
      fmt.Println("Uno")
    case 2:
      fmt.Println("Dos")
    case 3:
      fmt.Println("tres")
	}

	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("Es fin de semana")
		default:
			fmt.Println("Día entre semana")
	}

	t := time.Now()

	switch {
		case t.Hour() < 12:
			fmt.Println("La hora es por la mañana")
		default:
			fmt.Println("La hora es por la tarde")
	}

  whatAmI := func(i interface {}) {
      switch t := i.(type) {
      case bool:
        fmt.Println("I'm a bool")
      case int:
        fmt.Println("I'm an int")
      default:
        fmt.Println("Don't know type %T\n", t)
    }
  }

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}