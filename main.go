package main

import  (
          "fmt"
)

func main(){

fmt.Print("Enter your age: ")

var age int

fmt.Scanf("%d", &age)

if age < 3{
 fmt.Println("Infant")
}

fmt.Println(age)


}