package user

type User struct {
	name string
	age  int
}

var allUsers = []User{
	{name: "Lucas", age: 29},
	{name: "Roberto", age: 34},
	{name: "Priscila", age: 18},
	{name: "Claudia", age: 20},
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetAge() int {
	return u.age
}
