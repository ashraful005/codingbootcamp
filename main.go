package main

import "fmt"

func main(){

x := make(map[string]int)

x["abir"] = 10
x["aziz"] = 15

//delete(x, "aziz")

fmt.Println(x)
fmt.Println(x["abir"])

}