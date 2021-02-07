package main

import "fmt"
import "reflect"

func main(){

var students [3]string

students[0] = "Asgor"
students[1] = "Mainul"
students[2] = "Anonnya"

x := students[0:2]

fmt.Println(x)

var fruits []string

fruits = append(fruits, "Apple", "Banana", "Mango")

fmt.Println(fruits, len(fruits))

fmt.Printf("%T\n", fruits)

fmt.Printf("%T\n", students)

a := reflect.TypeOf(students).Kind().String()
b := reflect.TypeOf(fruits).Kind().String()

fmt.Println(a)
fmt.Println(b)

}