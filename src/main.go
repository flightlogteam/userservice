package main

import "fmt"
import "userservice/user"

func main() {
	repo, err := user.NewRepository("root", "pass", "localhost", "3306", "flightloguser")

	if err != nil {
		fmt.Errorf("Cannot create connection %v", err)
		return
	}

	/*res, err := repo.Create(&user.User{
		Givenname:  "Martin",
		Familyname: "Klingenberg",
		Email:      "martin@klingenberg.as",
		Active:     true,
		Privacy:    0,
	})*/

	res, err := repo.UserByEmail("martin@klingenberg.as")

	if err != nil {
		fmt.Errorf("Unable to create user: %v", err)
		return
	}
	fmt.Println("Works", res)
}
