package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type: ", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind is float64: ", v.Kind() == reflect.Float64)
	fmt.Println("value: ", v.Float())
	fmt.Println("settability of v: ", v.CanSet())
	fmt.Println("-------------------")

	p := reflect.ValueOf(&x)
	fmt.Println("type of p: ", p.Type())
	fmt.Println("settability of p: ", p.CanSet())
	v = p.Elem()
	fmt.Println("settability of v: ", v.CanSet())
	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}
