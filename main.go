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
}else if age>2 && age<13{
 fmt.Println("Children")
}else if age>12 && age<=19{
 fmt.Println("Teenager")
}else{
 fmt.Println("Adult")
}

fmt.Println(age)


}