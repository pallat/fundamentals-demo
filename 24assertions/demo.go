package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(greeting(123))
	whatDoToday(time.Now().Local().Add(-48 * time.Hour))
}

func assertion() {
	var i interface{}
	i = "text"
	// var s string = i --> doesn't work

	var n int
	if v, ok := i.(int); ok {
		n = v
	}
	fmt.Println(n)

}

func greeting(v interface{}) string {
	if s, ok := v.(string); ok {
		return "hello " + s
	}
	return "hello friend"
}

func whichType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("this is a string %s\n", v)
	case int:
		fmt.Printf("this is an integer %d\n", v)
	case float64:
		fmt.Printf("this is a floating point %f\n", v)
	default:
		fmt.Println("don't know")
	}
}

func whatdoToday(t time.Time) {
	switch {
	case t.Day() == 1:
		fallthrough
	case t.Day() == 16:
		fmt.Println("Did you buy a lotto?")
	case t.Weekday() == time.Sunday:
		fmt.Println("Netflix day!!")
	}
}

func whatDoToday(t time.Time) {
	switch t.Weekday() {
	case time.Saturday:
		fmt.Println("go shopping")
	case time.Sunday:
		fmt.Println("take some rest")
	case time.Monday:
		fallthrough
	case time.Tuesday:
		fmt.Println("go planing")
	default:
		fmt.Println("go working")
	}
}
