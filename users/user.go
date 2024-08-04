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

// String Array
// var users = map[string]string{
// 	"Lucas":   "Lucas",
// 	"Roberto": "Roberto",
// 	"Priscila": "Priscila",
// 	"Claudia": "Claudia",
// }
