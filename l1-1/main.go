package main

import "fmt"

type Human struct {
	FirstName string
	LastName  string
	Age       int
	Phone     string
}

type Action struct {
	Human
}

func (h Human) GetFullName() string { return h.FirstName + " " + h.LastName }

func (h Human) GetYearOfBirth() int { return 2025 - h.Age }

func (a Action) Run() {
	fmt.Println(a.GetFullName(), "running")
}

func (a Action) Call() string {
	return fmt.Sprintf("%s is calling: %s", a.GetFullName(), a.Phone)
}
func main() {
	human := Human{
		FirstName: "James",
		LastName:  "Bond",
		Age:       32,
		Phone:     "123-456-789",
	}

	action := Action{
		Human: human,
	}

	action.Run()

	fmt.Println(action.Call())
	fmt.Println(action.GetFullName())
	fmt.Println(action.GetYearOfBirth())
}
