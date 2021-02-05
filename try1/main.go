package main

import (
	"fmt"
	"reflect"
)

var (sentence string = "Hello World !")

func main() {
	fmt.Println(sentence)
	fmt.Println(reflect.ValueOf(sentence).Kind())
}

func monAddition(param1 int, param2 int) int{
	return param1+param2
}


func maSoustraction(param1 int, param2 int) int {
	return param1-param2
}

func maMultiplication(param1 int, param2 int) int {
	return param1*param2
}

func maDivision(param1 int, param2 int) int {
	return param1/param2
}


