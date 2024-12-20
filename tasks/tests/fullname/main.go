package fullname

import "fmt"

func main() {
	u := User{
		FirstName: "Misha",
		LastName:  "Popov",
	}

	fmt.Println(u.FullName())
}
