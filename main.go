package main

import "fmt"

func main(){

var Students [3]string

Students[0] = "Ashraful"
Students[1] = "Mainul"
Students[2] = "Anonnya"


fmt.Println(Students)
fmt.Println(len(Students))
fmt.Println(Students[0])
fmt.Println(Students[2])

Clients := [3]string{"Aziz", "Akib", "Ashish"}
fmt.Println(Clients)

Fruits := [...]string{"Mango", "Banana"}
fmt.Println(Fruits)

}