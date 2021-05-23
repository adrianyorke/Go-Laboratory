package main

import (
	"fmt"
	"time"
	"math"
	"strings"
    "reflect"
)

func main() {
	// https://golang.org/pkg/math/rand/
	var now time.Time = time.Now()
	var year int = now.Year()

    fmt.Printf("now = %T\n", now)
    fmt.Printf("year = %T\n", year)
    
    fmt.Println("now (TypeOf) =", reflect.TypeOf(now))
    fmt.Println("now (ValueOf) =", reflect.ValueOf(now).Kind())
    
	fmt.Println(year)
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("head first go"))
}
