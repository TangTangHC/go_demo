package main

import "fmt"

type MyStruct struct {
	name string
	age  int
}

func main() {

	m1 := MyStruct{name: "m1", age: 19}
	m2 := MyStruct{name: "m2", age: 20}

	hash := make(map[string]MyStruct)

	hash["h1"] = m1
	hash["h2"] = m2

	fmt.Println(hash)

	//hash["h1"].name = "h1"
	//
	//fmt.Println(hash["h1"])

	//arr := []*MyStruct{&m1, &m2}
	arr := []MyStruct{m1, m2}
	fmt.Println(arr[0])
	changeAge(200, arr)
	fmt.Println(arr[0])

	// chapter3/sources/slice_append.go
	var s []int // s被赋予零值nil
	fmt.Println(len(s), cap(s))
	s = append(s, 11)
	fmt.Println(len(s), cap(s)) //1 1
	s = append(s, 12)
	fmt.Println(len(s), cap(s)) //2 2
	s = append(s, 13)
	fmt.Println(len(s), cap(s)) //3 4
	s = append(s, 14)
	fmt.Println(len(s), cap(s)) //4 4
	s = append(s, 15)
	fmt.Println(len(s), cap(s)) //5 8
}

func changeAge(age int, arr []MyStruct) {
	myStruct := arr[0]
	myStruct.age = age
}
