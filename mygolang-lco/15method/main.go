package main

import "fmt"

func main() {

	fmt.Println("Structs in Go")

	Yash := User{"Yash ", "Kumar", 21, true, "yashkatyan@mail.com"}
	fmt.Println(Yash)
	fmt.Printf("Yash details are %+v\n", Yash)
	fmt.Printf("First Name is %v and Email is %v\n", Yash.FirstName, Yash.Email)
	Yash.GetStatus()
	Yash.GetAge()

}

type User struct {
	FirstName, LastName string
	Age                 int
	Status              bool
	Email               string
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) GetAge() {
	fmt.Println("User age is: ", u.Age)
}
