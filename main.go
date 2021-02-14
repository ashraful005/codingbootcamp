package main

import  (
          "fmt"
)

func main(){

fmt.Print("Enter your age: ")

var age int

fmt.Scanf("%d", &age)

switch age{
 case 1,2:
       fmt.Print("Infant")
 case 3,4,5,6,7,8,9,10,11,12:
       fmt.Println("Children")
       fallthrough
 case 13,14,15,16,17,18,19:
       fmt.Println("Teenager")
       fallthrough
 default:
       fmt.Println("Adult")
}

fmt.Println(age)


}