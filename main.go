package main

import "fmt"

//struct Definition

type Book struct{

 Title string
 Author string
 ISBN  string
 Price float32
 Pages int

}


func main(){


//Declaration method 01

var b1 Book
b1.Title = "An Introduction to Programming in Go"
b1.Author = "Caleb Doxsey"
b1.ISBN = "3869A-3894468B"
b1.Price = 10.50
b1.Pages = 165 

fmt.Println(b1)

fmt.Println(b1.Title)
fmt.Println(b1.Author)
fmt.Println(b1.ISBN)
fmt.Println(b1.Price)
fmt.Println(b1.Pages)


//Using custom datatypes

var w1 Weight
w1 = 30.30
fmt.Println(w1, name)

}


