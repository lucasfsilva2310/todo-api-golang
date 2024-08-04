package user

func GetUser(name string) *User {
	for _, user := range allUsers {
		if user.name == name {
			return &user
		}
	}
	return nil
}

func GetAllUsers() []User {
	return allUsers
}

func DeleteUserByName(name string) {
	for i, user := range allUsers {
		if user.name == name {
			allUsers = append(allUsers[:i], allUsers[i+1:]...)
			break
		}
	}
}

func UpdateUserAge(name string, age int) {
	for i, user := range allUsers {
		if user.name == name {
			allUsers[i].age = age
			break
		}
	}
}
