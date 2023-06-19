package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	name string
	age  int
	id   int
}

// Global Variable
var CountStudent = 10

// Package Level Variable CamelCase
var cntDevices = 9

func sum(a int, b int) int {

	return a + b
}

func summation(vals ...int) int {
	var s = 0
	var l = len(vals)

	for _, x := range vals {
		s += x
	}
	return s + l

}

func calc(a int, b int) (int, int) {
	return (a + b), (a - b)
}

func main() {

	// // Int to string
	fmt.Println(calc(5, 3))
	fmt.Println(10, 20, "Ok")
	fmt.Println(summation(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(sum(10, 12))
	s1 := Student{
		"Paresh",
		21,
		10011,
	}
	s1.name = "Karan"
	var s11 Student
	//  s11.name="cbsj";
	//  s11.age=19
	//  s11.id=52626
	fmt.Println(s1, s11)
	var conint = 10
	var converted string = strconv.Itoa(conint)
	fmt.Println(converted)

	// fmt.Println("Nainesh")
	// // local variable
	// var b=1
	// var e=50
	// var a int=5
	// var c=100
	// d:=19
	// var name="Nainesh";
	// fmt.Println(a+b+c+d+e+cntDevices);
	// fmt.Printf("%v , %T",a, a)
	// fmt.Println(name)

	// /// arr
	// var arr=[...]int{1,2,3,4}
	// arr[1]=20
	// fmt.Println(arr[1:4])
	// arr2:=arr
	// fmt.Println(arr2);
	// fmt.Println(len(arr))
	// // 2D matrix

	// var matrix=[...][3]int{{1,2,3},{1,2,3},{1,2,3}}
	// fmt.Println(matrix)

	// // map
	// empSal:=make(map[string]int)
	// empSal=map[string]int{
	// 	"Raj":2000,
	// 	"Rohot":1000,
	// }
	// empSal["Rbi"]=200
	// delete(empSal,"Rbi")

	// var Nainesh,ok=empSal["Nainesh"]

	// fmt.Println(empSal)
	// fmt.Println(ok,Nainesh)

	for i := 1; i <= 6; i++ {

		if i&1 == 1 {
			fmt.Printf("%v is odd\n", i)
		} else {

			fmt.Printf("%v is even\n", i)
		}
	}
}
